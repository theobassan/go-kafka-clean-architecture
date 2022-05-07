package event_context

import (
	"context"
	"go-kafka-clean-architecture/app_func/infrastructure/command/event_context"
	"go-kafka-clean-architecture/app_func/logger"

	"github.com/go-errors/errors"

	"github.com/segmentio/kafka-go"
)

func StartKafkaRouter(kafkaURL string) func(logError logger.LoggerError) func(create event_context.ProductTranslatedControllerCreate) {
	return func(logError logger.LoggerError) func(create event_context.ProductTranslatedControllerCreate) {
		return func(create event_context.ProductTranslatedControllerCreate) {
			//l := log.New(os.Stdout, "kafka reader: ", 0)
			reader := kafka.NewReader(kafka.ReaderConfig{
				Brokers:     []string{kafkaURL},
				GroupTopics: []string{"product"},
				GroupID:     "clean-architecture",
				//Logger:      l,
			})

			start(reader)(logError)(create)
		}
	}
}

func start(kafkaReader *kafka.Reader) func(logError logger.LoggerError) func(create event_context.ProductTranslatedControllerCreate) {
	return func(logError logger.LoggerError) func(create event_context.ProductTranslatedControllerCreate) {
		return func(create event_context.ProductTranslatedControllerCreate) {

			ctx := context.Background()
			for {
				msg, err := kafkaReader.FetchMessage(ctx)
				if !errors.Is(err, nil) {
					logError(err)
				}

				if msg.Topic == "product" {
					kafkaContext := NewKafkaContext(ctx, kafkaReader, msg)

					err = create(kafkaContext.Bind, kafkaContext.Acknowledge)
					if !errors.Is(err, nil) {
						logError(err)
					}
				}
			}
		}
	}
}
