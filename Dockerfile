FROM golang:1.15-alpine as build-env

WORKDIR /echo-server/


COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o /go/bin/echo-server echo-server

CMD [ "/go/bin/echo-server"]

EXPOSE 8081