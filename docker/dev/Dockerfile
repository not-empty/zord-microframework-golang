
FROM golang:1.22-alpine AS build

ENV GO111MODULE=on

WORKDIR /app
COPY . .

RUN apk add curl
RUN curl -sSf https://atlasgo.sh | sh

RUN go mod download
RUN go install github.com/cosmtrek/air@latest
