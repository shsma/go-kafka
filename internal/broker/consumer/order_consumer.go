package consumer

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"log"
)

type OrderConsumer struct {
	Consumer *kafka.Consumer
	Topic    string
}

func NewOrderConsumer(topic string) *OrderConsumer {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "goapp-consumer",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Fatalf("failed to create consumer: %v", err)
	}

	return &OrderConsumer{
		Consumer: consumer,
		Topic:    topic,
	}
}

func (oc *OrderConsumer) Consume() {
	err := oc.Consumer.SubscribeTopics([]string{oc.Topic}, nil)
	if err != nil {
		log.Fatalf("failed to subscribe topic: %v", err)
	}

	for {
		ev := oc.Consumer.Poll(100)
		switch e := ev.(type) {
		case *kafka.Message:
			log.Printf("%% Message on %s:\n%s\n", e.TopicPartition, string(e.Value))
		case kafka.Error:
			log.Printf("%% Error: %v\n", e)
		}
	}
}
