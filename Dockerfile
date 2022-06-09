# syntax=docker/dockerfile:1
FROM golang:1.18-alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./build .

EXPOSE 3000

CMD ["./build"]
