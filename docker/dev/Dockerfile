
FROM golang:1.22-alpine AS build

ENV GO111MODULE=on

WORKDIR /app
COPY . .

RUN apk add curl
ENV ATLAS_VERSION=v0.20.1-7793896-canary
RUN curl -sSf https://atlasgo.sh | sh

RUN go mod download
RUN go install github.com/air-verse/air@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN apk add build-base

CMD ["air", "http"]
