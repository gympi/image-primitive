FROM golang:1.8

WORKDIR /go/src/app

COPY . .

ENTRYPOINT go run *.go

EXPOSE 9001
