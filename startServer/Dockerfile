FROM golang:latest  


ADD . /go/src/github.com/parth-gr/rh_project/startServer


RUN go install /go/src/github.com/parth-gr/rh_project/startServer

ENTRYPOINT /go/bin/startServer

EXPOSE 4000


