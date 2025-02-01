FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o sigma_evolution ./cmd/main.go

CMD ["./sigma_evolution"]