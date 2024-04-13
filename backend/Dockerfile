FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o main cmd/main.go

FROM ubuntu:latest

WORKDIR /app

COPY . .

COPY --from=builder /app/main .

ENV CONFIG_PATH=config/local.yaml

EXPOSE 8080

CMD ["./main"]
