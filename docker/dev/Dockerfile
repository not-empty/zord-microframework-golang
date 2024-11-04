
FROM golang:1.23-alpine

ENV GO111MODULE=on

WORKDIR /app
RUN apk add git

COPY . .

RUN apk add curl
ENV ATLAS_VERSION=v0.20.1-7793896-canary
RUN curl -sSf https://atlasgo.sh | sh

RUN go mod download
RUN go install github.com/air-verse/air@v1.61.1
RUN go install github.com/swaggo/swag/cmd/swag@v1.16.4
RUN apk add build-base

CMD ["air", "http"]
