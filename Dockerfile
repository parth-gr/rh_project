FROM golang:latest

RUN mkdir /app
COPY . /app
WORKDIR /app

#WORKDIR /go/src/github.com/parth-gr/rh_project/startServer/main

#COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD  ["go","run","../app/startServer/main.go"]
