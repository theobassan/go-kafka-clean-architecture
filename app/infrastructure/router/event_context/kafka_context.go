package kafka

import (
	"context"
	"encoding/json"
	"errors"
	"go-kafka-clean-architecture/app/input/router"

	"github.com/segmentio/kafka-go"
)

type KafkaContext struct {
	ctx          context.Context
	kafkaReader  *kafka.Reader
	kafkaMessage kafka.Message
}

func NewKafkaContext(ctx context.Context, kafkaReader *kafka.Reader, kafkaMessage kafka.Message) router.EventContext {
	return &KafkaContext{ctx, kafkaReader, kafkaMessage}
}

func (context *KafkaContext) Bind(v any) error {
	err := json.Unmarshal(context.kafkaMessage.Value, v)
	if !errors.Is(err, nil) {
		return err
	}
	return nil
}

func (context *KafkaContext) Acknowledge() error {
	err := context.kafkaReader.CommitMessages(context.ctx, context.kafkaMessage)
	if !errors.Is(err, nil) {
		return err
	}
	return nil
}