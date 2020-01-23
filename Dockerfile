FROM golang:1.13.6-alpine3.11 as build-env

RUN mkdir /app

ADD . /app

WORKDIR /app/cmd/monitor

RUN go build

CMD ["./monitor"]