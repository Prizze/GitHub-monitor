APP_NAME := GitHub-monitor
SRC := ./cmd/main.go

# Компиляция бинарного файла
build:
	go build -o $(APP_NAME) $(SRC)

# Очитска бинарного файла
clean:
	rm -f $(APP_NAME)

# Запуск проекта
run: build
	./$(APP_NAME)

# Прогон тестов
test:
	go test ./...

# Генерация swagger документации
swagger:
	swag init -g ${SRC}