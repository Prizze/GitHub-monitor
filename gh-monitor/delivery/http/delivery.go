package http

import "net/http"

type GHMonitorDelivery struct {
	usecase ghMonitorUsecase
}

func NewGHMonitorDelivery(usecase ghMonitorUsecase) *GHMonitorDelivery {
	return &GHMonitorDelivery{
		usecase: usecase,
	}
}

func (h *GHMonitorDelivery) GetTop(w http.ResponseWriter, r *http.Request) {

}
