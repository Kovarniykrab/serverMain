#!/bin/bash
FROM golang:1.23-alpine AS builder

RUN apk add --no-cache make git build-base --repository=http://dl-cdn.alpinelinux.org/alpine/v3.19/main

WORKDIR /serverMain
COPY . .
RUN go mod download
ENV GO111MODULE=on
RUN go mod tidy
RUN export `cat .env`
RUN make build

FROM alpine:latest AS runner
EXPOSE $SERVER_WEB_PORT
EXPOSE 7000
COPY --from=builder /serverMain/bin/http /serverMain/bin/http
COPY --from=builder /serverMain /serverMain/
WORKDIR /serverMain/
ENTRYPOINT ["bin/http"]