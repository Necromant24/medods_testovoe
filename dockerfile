# Указываем базовый образ с установленным Go
FROM golang:1.23 as builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файл go.mod и go.sum, чтобы установить зависимости
COPY go.mod go.sum ./

# Устанавливаем зависимости
RUN go mod download

# Копируем исходный код проекта
COPY . .

# Сборка проекта
RUN go build -o main .

# Используем минимальный образ для финального контейнера
FROM debian:bullseye-slim

# Создаем директорию для приложения
WORKDIR /app

# Копируем бинарный файл из предыдущего этапа
COPY --from=builder /app/main .

# Указываем порт, на котором работает приложение (опционально)
EXPOSE 8080
EXPOSE 5432

# Команда для запуска приложения
CMD ["./main"]
