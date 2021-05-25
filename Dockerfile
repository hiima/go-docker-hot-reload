FROM golang:1.16-buster

WORKDIR /go/src/app

COPY ./go.mod .
COPY ./go.sum .

RUN go mod download
# ホットリロードのため air を入れる
RUN go get -u github.com/cosmtrek/air

COPY . .
