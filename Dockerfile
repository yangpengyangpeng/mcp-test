FROM golang:latest

ADD . /go/src/app

WORKDIR /go/src/app
RUN go build
CMD ["./mcp-test"]
EXPOSE 8085
