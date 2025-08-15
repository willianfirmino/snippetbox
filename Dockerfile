FROM golang:1.25-alpine AS builder
WORKDIR /app

COPY go.mod .
COPY go.sum .


RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/web ./cmd/web

FROM alpine:latest

ENV TZ America/Sao_Paulo
RUN apk add --no-cache tzdata

WORKDIR /app

COPY --from=builder /app/web .

COPY tls ./tls

EXPOSE 4000

CMD ["./web"]