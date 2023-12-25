FROM golang:1.22rc1-alpine3.18 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/ cmd/
COPY src/ src/

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main cmd/api/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
