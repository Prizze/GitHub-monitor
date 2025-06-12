APP_NAME := GitHub-monitor
SRC := ./cmd/main.go

build:
	go build -o $(APP_NAME) $(SRC)

run: build
	./$(APP_NAME)

clean:
	rm -f $(APP_NAME)

test:
	go test ./...

