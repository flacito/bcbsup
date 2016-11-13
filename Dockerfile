FROM golang:1.6.3-alpine

MAINTAINER btwebb@gmail.com

COPY . /go/src/github.com/flacito/bcbsup
RUN go install github.com/flacito/bcbsup/create_machine
RUN go install github.com/flacito/bcbsup/bcbsup
# RUN ls -R /go
