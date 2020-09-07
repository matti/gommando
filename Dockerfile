FROM golang:1.15.0-alpine3.12

WORKDIR /build
COPY go.* ./
RUN go mod download

COPY . .
RUN go build gommando.go
