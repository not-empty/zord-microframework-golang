FROM golang:1.18-alpine AS build

ENV GO111MODULE="on"
ENV GOOS="linux"
ENV CGO_ENABLED=0

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /go-skeleton

## Deploy
FROM alpine

WORKDIR /

COPY --from=build /go-skeleton /

RUN touch .env
RUN touch .config.env

EXPOSE 8080

CMD ["./go-skeleton", "http"]
