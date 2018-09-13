FROM golang:1.11-alpine AS build-env

RUN apk update && apk add git

RUN mkdir -p /go/src/github.com/cscatolini/pitaya-maestro-demo
ADD . /go/src/github.com/cscatolini/pitaya-maestro-demo

WORKDIR /go/src/github.com/cscatolini/pitaya-maestro-demo
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure
RUN go build -o main .

FROM alpine:3.8

RUN apk update && apk add ca-certificates

WORKDIR /app
COPY --from=build-env /go/src/github.com/cscatolini/pitaya-maestro-demo/main /app

CMD ["./main"]
