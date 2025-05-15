FROM golang:latest

ADD . /go/src/app

WORKDIR /go/src/app
COPY go.mod .
RUN go build
CMD ["./mcp-test"]
EXPOSE 8081
