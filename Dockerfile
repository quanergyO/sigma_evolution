FROM golang:1.23-alpine AS builder

RUN apk add --no-cache postgresql-client bash

WORKDIR /app

COPY . .

COPY build/backend-builder/wait-for-db.sh /build/backend-builder/wait-for-db.sh
RUN chmod +x /build/backend-builder/wait-for-db.sh

RUN go mod download
RUN go build -o sigma_evolution ./cmd/main.go

CMD ["./sigma_evolution"]