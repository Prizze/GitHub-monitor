package middleware

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

const (
	rateLimitRPS    = 1               // Количество запросов в секунду
	rateLimitBurst  = 1	              // Максимум запросов за раз
	cleanUpInterval = 5 * time.Minute // Таймер очистки лимитеров
)

type clientLimiter struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

type rateLimiter struct {
	clients map[string]*clientLimiter
	mu      sync.Mutex
}

// Очистка лимитеров
func (rl *rateLimiter) cleanUpOldClients() {
	ticker := time.NewTicker(cleanUpInterval)
	defer ticker.Stop()

	for {
		<-ticker.C

		rl.mu.Lock()
		for ip, cl := range rl.clients {
			if time.Since(cl.lastSeen) > cleanUpInterval {
				delete(rl.clients, ip)
			}
		}
		rl.mu.Unlock()
	}
}

func (rl *rateLimiter) getLimiter(ip string) *rate.Limiter {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	if cl, exists := rl.clients[ip]; exists {
		cl.lastSeen = time.Now()
		return cl.limiter
	}

	limiter := rate.NewLimiter(rateLimitRPS, rateLimitBurst)
	rl.clients[ip] = &clientLimiter{limiter: limiter, lastSeen: time.Now()}
	return limiter
}

func NewRateLimiter() *rateLimiter {
	rl := &rateLimiter{
		clients: make(map[string]*clientLimiter),
	}
	go rl.cleanUpOldClients()
	return rl
}

func (rl *rateLimiter) RateLimitingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// @TODO Заменить на 'X-Forwarded-For'
		ip := req.RemoteAddr
		limiter := rl.getLimiter(ip)

		if !limiter.Allow() {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, req)
	})
}
