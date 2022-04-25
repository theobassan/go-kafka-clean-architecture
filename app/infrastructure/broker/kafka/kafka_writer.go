package kafka

import (
	"context"
	"go-kafka-clean-architecture/app/interfaces/broker"
	"log"
	"os"
	"strconv"

	"github.com/segmentio/kafka-go"
)

type eventWriter struct {
	kafkaWriter *kafka.Writer
}

func NewKafkaWriter(kafkaURL string, topic string) broker.EventWriter {
	l := log.New(os.Stdout, "kafka writer: ", 0)
	kafkaWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{kafkaURL},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
		// assign the logger to the writer
		Logger: l,
	})

	return &eventWriter{kafkaWriter}
}

func (writer *eventWriter) WriteMessage(i interface{}) error {
	return writer.kafkaWriter.WriteMessages(context.Background(), i.(kafka.Message))
}

func (writer *eventWriter) Bind(value []byte) interface{} {
	return kafka.Message{
		Key:   []byte(strconv.Itoa(0)),
		Value: value,
	}
}
