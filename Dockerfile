FROM golang:1.25 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o wishforge ./cmd/app

FROM alpine:3.20

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/wishforge .
COPY --from=builder /app/migrations ./migrations
COPY .env .

EXPOSE 8080

CMD ["./wishforge"]