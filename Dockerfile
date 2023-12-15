FROM golang:1.21-alpine3.17 AS builder

WORKDIR /app

COPY . .

RUN go build -o main .

EXPOSE 8080
CMD "[/app/main]"