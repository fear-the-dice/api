FROM golang

ADD . /go/src/github.com/fear-the-dice/api

RUN go install github.com/fear-the-dice/api
ENTRYPOINT /go/bin/api

EXPOSE 3000
