FROM golang:alpine

RUN apk update && apk add --no-cache git
RUN mkdir -p /go/src/app
WORKDIR /go/src/app

ADD . /go/src/app

RUN go get -v
