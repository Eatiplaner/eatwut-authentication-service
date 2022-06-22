FROM golang:1.17.2-alpine AS base

# Add required packages
RUN apk add --update git curl bash && \
		go get -u github.com/revel/revel && \
		go get -u github.com/revel/cmd/revel

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
ENV CGO_ENABLED 0

ADD . .

# Dev target
FROM base as development
ENTRYPOINT revel run

# Test target
FROM base as test
RUN revel test .

FROM base as builder
ENV GOOS=linux \
    GOARCH=amd64
