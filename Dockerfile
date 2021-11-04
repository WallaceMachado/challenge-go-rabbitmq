FROM golang:1.16.7

WORKDIR /go/src/app

COPY . .

RUN go mod download
RUN go build

CMD ["./challenge-go-rabbitmq"]