FROM golang:1.16-buster

WORKDIR /go/src/app

COPY go.* ./
RUN go mod download

RUN go get -u github.com/cosmtrek/air

COPY . ./
