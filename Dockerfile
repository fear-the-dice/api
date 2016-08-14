FROM golang

ADD . /go/src/github.com/fear-the-dice/api

ENV FTDAUTHKEY=supersecret
ENV DB=ftd-test
ENV MONGOLAB_URI=mongodb://localhost
ENV REDIS_URL=redis://localhost:6379
ENV PORT=:3000

RUN go install github.com/fear-the-dice/api
ENTRYPOINT /go/bin/api

EXPOSE 3000
