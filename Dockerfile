#################
# Build Stage
#################
FROM golang:1.13-alpine as builder

RUN apk add curl git make g++ gcc librdkafka-dev --repository=http://dl-cdn.alpinelinux.org/alpine/edge/main --repository=http://dl-cdn.alpinelinux.org/alpine/edge/community

## Create a directory inside the container to store all our application and then make it the working directory.
RUN mkdir -p /go/src/github.com/shipt/campaign-service
WORKDIR /go/src/github.com/shipt/campaign-service

## Copy the app directory (where the Dockerfile lives) into the container.
COPY . .

ARG GITHASH
# Disable Go cross compiling, Compile Linux only, Run Makefile

# API
RUN GOOS=linux go build -mod vendor -a -installsuffix cgo -o cliqueup-service cmd/api/main.go
###################
# Docs Stage
###################
FROM node:12 as docs

RUN mkdir -p /docs

###################
# Package Stage
###################
FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

COPY --from=builder /go/src/github.com/samuelmahr/cliqueup-service/campaign-service-api /usr/local/bin/cliqueup-service-api

CMD ["/usr/local/bin/cliqueup-service-api"]