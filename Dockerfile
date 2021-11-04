FROM golang:1.16.7

WORKDIR /go/src/app

ADD go.mod .
ADD go.sum .
RUN go mod download

COPY . .

RUN go build

CMD ["./challenge-go-rabbitmq"]