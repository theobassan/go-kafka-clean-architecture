package event_api

import (
	"context"
	"go-kafka-clean-architecture/app/interfaces/api"
	"strconv"

	"github.com/segmentio/kafka-go"
)

type kafkaApi struct {
	kafkaWriter *kafka.Writer
}

func NewKafkaApi(kafkaURL string) api.EventApi {
	//l := log.New(os.Stdout, "kafka writer: ", 0)
	kafkaWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{kafkaURL},
		Balancer: &kafka.LeastBytes{},
		//Logger:   l,
	})

	return &kafkaApi{kafkaWriter}
}

func (kafkaApi *kafkaApi) WriteMessage(i interface{}) error {
	msg := i.(kafka.Message)
	return kafkaApi.kafkaWriter.WriteMessages(context.Background(), msg)
}

func (kafkaApi *kafkaApi) Bind(topic string, value []byte) interface{} {
	return kafka.Message{
		Key:   []byte(strconv.Itoa(0)),
		Topic: topic,
		Value: value,
	}
}
