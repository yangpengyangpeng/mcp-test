FROM golang:latest

ADD . /go/src/mcp-test

WORKDIR /go/src/mcp-test
RUN go build
CMD ["./mcp-test"]
EXPOSE 8082
