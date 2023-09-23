ARG GO_VERSION=1.21

FROM golang:${GO_VERSION}-alpine AS builder

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

WORKDIR /prbot

COPY . .

RUN go mod download

RUN go build -o /prbot/app main.go

FROM alpine:latest

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

COPY --from=builder /prbot/app /app

EXPOSE 8080

ENTRYPOINT ["/app"]