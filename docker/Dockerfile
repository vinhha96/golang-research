FROM golang:alpine AS builder

RUN apk update && apk add git && apk add ca-certificates &&  apk add bash

WORKDIR /go/src/github.com/vinhha96/golang-research

ADD  . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./main/web_service ./main/main.go

RUN chmod +x ./docker/wait-for-it.sh

EXPOSE 3000

CMD ["./main/web_service"]

