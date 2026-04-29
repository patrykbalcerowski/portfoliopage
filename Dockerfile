# ETAP 1: Budowanie (Builder)
FROM golang:1.22-alpine AS builder

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY go-application/go.mod go.sum ./
RUN go mod download

COPY go-application/. .
RUN CGO_ENABLED=0 GOOS=linux go build -o nextcloud-exporter .

FROM alpine:latest

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
WORKDIR /root/
COPY --from=builder /app/nextcloud-exporter .

EXPOSE 8082

CMD ["./nextcloud-exporter"]