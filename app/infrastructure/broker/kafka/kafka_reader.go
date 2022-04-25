package kafka

import (
	"context"
	"go-kafka-clean-architecture/app/interfaces/broker"

	"github.com/segmentio/kafka-go"
)

type eventReader struct {
	kafkaReader *kafka.Reader
}

func NewKafkaReader(kafkaURL string, topic string, groupID string) broker.EventReader {
	//l := log.New(os.Stdout, "kafka reader: ", 0)
	kafkaReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{kafkaURL},
		Topic:   topic,
		GroupID: groupID,
		// assign the logger to the writer
		//Logger: l,
	})

	return &eventReader{kafkaReader}
}

func (reader *eventReader) ReadMessage() (interface{}, error) {
	return reader.kafkaReader.ReadMessage(context.Background())
}

func (reader *eventReader) Bind(i interface{}) []byte {
	return i.(kafka.Message).Value
}
