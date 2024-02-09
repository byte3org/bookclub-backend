FROM golang:1.19.2-alpine as builder 

RUN apk add bash

RUN apk add --no-cache openssh-client ansible git

WORKDIR /backend

COPY go.mod go.sum ./ 

RUN go mod download

COPY . .

RUN go build -o app cmd/backend/main.go

RUN chmod a+x app 

EXPOSE 8001

CMD ["./app"]
