FROM golang:1.22-alpine AS build

ENV GO111MODULE="on"
ENV GOOS="linux"
ENV CGO_ENABLED=0
ENV GOGC=70

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /go-skeleton cmd/http/main.go

## Deploy
FROM alpine:3.20.2

WORKDIR /

COPY --from=build /go-skeleton /

RUN touch .env
RUN touch .config.env

EXPOSE 8080

CMD ["./go-skeleton", "http"]
