FROM golang:1.16-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /challenge-go-rabbitmq


FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /challenge-go-rabbitmq /challenge-go-rabbitmq

ENTRYPOINT ["/challenge-go-rabbitmq"]