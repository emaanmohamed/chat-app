FROM golang:1.23.0-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go get -d -v ./...

RUN go build -o chat-app ./cmd/main.go

FROM alpine:latest

WORKDIR /app


RUN mkdir -p /app/media/uploads


COPY --from=builder /app/chat-app .
COPY --from=builder /app/web ./web

EXPOSE 8083

CMD ["./chat-app"]