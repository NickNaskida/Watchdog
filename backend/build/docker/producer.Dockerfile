FROM --platform=linux/amd64 golang:1.21.6-alpine as builder

MAINTAINER "Nikoloz Naskidashvili"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o producer cmd/producer.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/producer .

CMD ["/app/producer"]
