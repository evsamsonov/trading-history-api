FROM golang:1.12-alpine

RUN apk update && apk add git && apk add bash

WORKDIR /go/src/github.com/evsamsonov/trading-history-server

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

RUN chmod +x ./docker/golang/wait-for-it.sh

EXPOSE 8080

CMD ["./docker/golang/wait-for-it.sh", "db:3306", "--", "trading-history-server"]