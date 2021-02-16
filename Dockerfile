FROM golang:latest

RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go get -d -v ./...
RUN go install -v ./...
CMD  ["go","run","../app/startServer/main.go"]
