FROM golang:1.17.2-alpine AS base

# Add required packages
RUN apk add --update git curl bash

# Install revel framework
RUN go get -u github.com/revel/revel
RUN go get -u github.com/revel/cmd/revel

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
ENV CGO_ENABLED 0

ADD . .

# Dev target
FROM base as dev
ENTRYPOINT revel run

# Unit Test target
FROM base as unit-test
RUN revel test .

# Prod target
FROM base as prod
ENV GOOS=linux \
    GOARCH=amd64

RUN revel package .

# Run stage
FROM alpine:3.13
RUN apk update && \
    apk add mailcap tzdata && \
    rm /var/cache/apk/*
WORKDIR /app
COPY --from=builder /app/app.tar.gz .
RUN tar -xzvf app.tar.gz && rm app.tar.gz
ENTRYPOINT /app/run.sh
