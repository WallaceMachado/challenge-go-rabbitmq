package queue

import (
	"github.com/streadway/amqp"
	"github.com/wallacemachado/challenge-go-rabbitmq/src/config"
)

func Connect() *amqp.Channel {
	dsn := "amqp://" + config.RabbitmqUser + ":" + config.RabbitmqPass + "@" + config.RabbitmqHost + ":" + config.RabbitmqPort + config.RabbitmqVhost

	conn, err := amqp.Dial(dsn)
	if err != nil {
		panic(err.Error())
	}

	channel, err := conn.Channel()
	if err != nil {
		panic(err.Error())
	}
	return channel
}

func Notify(payload []byte, exchange string, routingKey string, ch *amqp.Channel) {

	err := ch.Publish(
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(payload),
		})

	if err != nil {
		panic(err.Error())
	}

}
