FROM golang:1.4.2
MAINTAINER Abdulkadir Yaman <abdulkadiryaman@gmail.com>

RUN mkdir /tmp/gopath
ENV GOPATH /tmp/gopath

RUN go get github.com/yaman/timeout

ENTRYPOINT ${GOPATH}/bin/timeout -proto=$PROTO -port=$PORT

