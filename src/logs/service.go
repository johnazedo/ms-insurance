package logs

import (
	"encoding/json"
	"time"
	"github.com/streadway/amqp"
)

type Config struct {
	Microservice string
	Thread string
	Context string
	IP string
	Key string
}

type LogService interface {
	SendLog(input LogInput) 
}

type LogServiceImpl struct {
	LogService
	Config *Config
	RabbitMQ *amqp.Connection
}

func (s *LogServiceImpl) SendLog(input LogInput) {

	channel , err := s.RabbitMQ.Channel()
	if err != nil {
		panic(err)
	}
	defer channel.Close()

	body := Log {
		TimeStamp: time.Now().String(),
		Level: input.Level,
		Microservice: s.Config.Microservice,
		Thread: s.Config.Thread,
		Class: input.Class,
		Method: input.Method,
		Message: input.Message,
		Context: s.Config.Context,
		Ip: s.Config.IP,
	}

	buffer, err := json.Marshal(&body)
	if err != nil {
		panic(err)
	}

	err = channel.Publish(
		"",        // exchange
		s.Config.Key, // key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing {
			ContentType: "application/json",
			Body: buffer,
		},
	)
	if err != nil {
		panic(err)
	}
}
