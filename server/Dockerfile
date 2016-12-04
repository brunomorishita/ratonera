FROM golang

WORKDIR /go/src/server

COPY . .

RUN go get -d -v

RUN go build server.go

CMD ["./server", "--mongo", "mongo-server"]
