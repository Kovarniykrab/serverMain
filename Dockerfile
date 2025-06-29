#!/bin/bash
FROM golang:1.23-alpine AS builder

WORKDIR /servermain
COPY . /servermain
RUN apk add make
ENV GO111MODULE=on
RUN go mod tidy
RUN make build
FROM scratch

COPY --from=builder /servermain/bin /servermain/bin
EXPOSE  $SERVER_WEB_PORT
CMD ["/servermain/bin/srv"]
