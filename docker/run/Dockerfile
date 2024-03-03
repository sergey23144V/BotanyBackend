# Используем базовый образ с Go для сборки
FROM golang:alpine3.18 as builder

COPY . /botany-service

WORKDIR /botany-service

# Установка Git
RUN apk --no-cache add git

# Копируем файлы проекта и выполняем сборку
COPY ./go.mod ./go.sum ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o botany github.com/sergey23144V/BotanyBackend/cmd/



# Вторая ступень - образ Alpine для выполнения
FROM alpine:3.14

WORKDIR /botany-service

 # Установка временных зон в Alpine Linux
RUN apk add tzdata

# Применение новой временной зоны
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime

RUN echo $TZ > /etc/timezone

# Копируем бинарный файл из первой ступени
COPY --from=builder /botany-service/. .

RUN chmod +x /bin/sh && apk --no-cache add nmap bash curl postgresql-client

# Команда для запуска приложения
CMD ["/botany-service/botany"]