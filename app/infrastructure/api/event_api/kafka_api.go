package event_api

import (
	"context"
	"go-kafka-clean-architecture/app/interfaces/api"
	"strconv"

	"github.com/segmentio/kafka-go"
)

type kafkaAPI struct {
	kafkaWriter *kafka.Writer
}

func NewKafkaAPI(kafkaURL string) api.EventAPI {
	//l := log.New(os.Stdout, "kafka writer: ", 0)
	kafkaWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{kafkaURL},
		Balancer: &kafka.LeastBytes{},
		//Logger:   l,
	})

	return &kafkaAPI{kafkaWriter}
}

func (kafkaAPI *kafkaAPI) WriteMessage(i interface{}) error {
	msg := i.(kafka.Message)
	return kafkaAPI.kafkaWriter.WriteMessages(context.Background(), msg)
}

func (kafkaAPI *kafkaAPI) Bind(topic string, value []byte) interface{} {
	return kafka.Message{
		Key:   []byte(strconv.Itoa(0)),
		Topic: topic,
		Value: value,
	}
}
