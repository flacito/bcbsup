FROM golang:1.6.3-alpine

MAINTAINER btwebb@gmail.com

COPY . /go/src/app
RUN go install app/create_machine
