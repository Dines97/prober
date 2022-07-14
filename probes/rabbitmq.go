package probes

import "github.com/rabbitmq/amqp091-go"

type RabbitMQ struct {
	ConnectionString string `json:"connectionString"`
}

func (r RabbitMQ) Run() (RunResult, error) {
	_, err := amqp091.Dial(r.ConnectionString)
	if err != nil {
		return Fail, err
	}

	return Success, nil
}
