package producer

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"log"
)

type OrderProducer struct {
	Producer     *kafka.Producer
	Topic        string
	deliveryChan chan kafka.Event
}

func NewOrderProducer(topic string) *OrderProducer {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "0.0.0.0:9092",
		"client.id":         "goapp-consumer",
		"acks":              "all",
	})
	if err != nil {
		log.Fatalf("failed to create consumer: %v", err)
	}

	return &OrderProducer{
		Producer:     p,
		Topic:        topic,
		deliveryChan: make(chan kafka.Event, 10000),
	}
}

func (op *OrderProducer) Produce(msg []byte) error {
	err := op.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &op.Topic,
			Partition: kafka.PartitionAny,
		},
		Value: msg,
	}, op.deliveryChan)
	if err != nil {
		return err
	}
	<-op.deliveryChan
	return nil
}
