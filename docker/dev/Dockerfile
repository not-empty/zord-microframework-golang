FROM golang:1.18-alpine AS build

ENV GO111MODULE=on

WORKDIR /app

COPY . .

RUN go mod download
RUN go install github.com/cosmtrek/air@latest
