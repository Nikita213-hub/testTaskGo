FROM golang:1.23.5-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o quotes_server cmd/main.go

FROM alpine:3.21.3

WORKDIR /app

ENV ADDRESS="0.0.0.0:8080"

COPY --from=builder /app/quotes_server .

EXPOSE 8080

ENTRYPOINT ["./quotes_server"]