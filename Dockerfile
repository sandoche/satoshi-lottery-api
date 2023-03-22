FROM golang:1.20.2-alpine3.17

ENV GIN_MODE=release
ENV PORT=8080

WORKDIR /go/src/app

COPY . /go/src/app

RUN apk update && apk add --no-cache git
RUN go mod vendor

RUN go build server.go

EXPOSE $PORT

ENTRYPOINT ["./server"]
