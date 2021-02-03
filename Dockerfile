FROM golang:latest

RUN mkdir /build
WORKDIR /build


RUN  export GO111MODULE=on

RUN go get github.com/parth-gr/rh_project/startServer


RUN cd /build && git clone https://github.com/parth-gr/rh_project.git

RUN cd /build/rh_project/startServer && go build
EXPOSE 4000

ENTRYPOINT ["bash","/build/rh_project/startServer/main"]

