package configs

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/streadway/amqp"
)

func ConnectRabbitMQ() (*amqp.Channel, error) {
	conn, errConn := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if errConn != nil {
		return nil, errConn
	}
	defer conn.Close()

	ch, errConn := conn.Channel()
	if errConn != nil {
		return nil, errConn
	}
	defer ch.Close()

	return ch, nil
}
