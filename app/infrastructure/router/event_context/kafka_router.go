package kafka

import (
	"context"
	"go-kafka-clean-architecture/app/command/controller/event_context"
	"log"

	"github.com/go-errors/errors"

	"github.com/segmentio/kafka-go"
)

type kafkaHandler struct {
	appController *event_context.AppController
}

func StartKafkaRouter(appController *event_context.AppController, kafkaURL string) {
	kafkaHandler := &kafkaHandler{appController}

	//l := log.New(os.Stdout, "kafka reader: ", 0)
	kafkaReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{kafkaURL},
		GroupTopics: []string{"product"},
		GroupID:     "clean-architecture",
		//Logger:      l,
	})

	kafkaHandler.start(kafkaReader)
}

func (handler *kafkaHandler) start(kafkaReader *kafka.Reader) {
	ctx := context.Background()
	for {
		msg, err := kafkaReader.FetchMessage(ctx)
		if !errors.Is(err, nil) {
			log.Fatalln(err)
		}

		if msg.Topic == "product" {
			kafkaContext := NewKafkaContext(ctx, kafkaReader, msg)
			err = handler.appController.ProductTranslatedController.Create(kafkaContext)
			if !errors.Is(err, nil) {
				log.Fatalln(err)
			}
		}
	}
}
