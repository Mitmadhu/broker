FROM golang:1.21.1-alpine AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o broker-service

FROM alpine:latest

WORKDIR /app

COPY --from=builder  /app/broker-service .

COPY config.json .

EXPOSE 8080

CMD [ "./broker-service" ]

