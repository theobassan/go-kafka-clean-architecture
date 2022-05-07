package event_api

import (
	"context"
	"go-kafka-clean-architecture/app_func/interfaces/api"
	"strconv"

	"github.com/segmentio/kafka-go"
)

func NewKafkaApi(kafkaURL string) (api.EventApiBind, api.EventApiWriteMessage) {
	//l := log.New(os.Stdout, "kafka writer: ", 0)
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{kafkaURL},
		Balancer: &kafka.LeastBytes{},
		//Logger:   l,
	})

	return bind(), writeMessage(writer)
}

func writeMessage(kafkaWriter *kafka.Writer) api.EventApiWriteMessage {
	return func(i interface{}) error {
		msg := i.(kafka.Message)
		return kafkaWriter.WriteMessages(context.Background(), msg)
	}
}

func bind() api.EventApiBind {
	return func(topic string, value []byte) interface{} {
		return kafka.Message{
			Key:   []byte(strconv.Itoa(0)),
			Topic: topic,
			Value: value,
		}
	}
}
