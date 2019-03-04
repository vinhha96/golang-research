FROM golang:alpine AS builder

RUN apk update && apk add git && apk add ca-certificates

WORKDIR /go/src/github.com/vinhha96/golang-research

ADD  . .

RUN ls -lah

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./main/web_service ./main/main.go

EXPOSE 3000

ENTRYPOINT ["./main/web_service"]

