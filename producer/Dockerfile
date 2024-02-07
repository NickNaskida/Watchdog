FROM --platform=linux/amd64 golang:1.21.6-alpine as builder

MAINTAINER "Nikoloz Naskidashvili"

WORKDIR /app

COPY . .

RUN go get

RUN go build -o producer

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/producer .

CMD ["/app/producer"]
