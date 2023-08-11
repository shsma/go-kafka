package main

import (
	"fmt"
	"github.com/shsma/go-kafka/internal/broker/consumer"
	"github.com/shsma/go-kafka/internal/broker/producer"
	"log"
	"time"
)

func main() {
	topic := "test-topic"

	go func() {
		oc := consumer.NewOrderConsumer(topic)
		oc.Consume()
	}()

	op := producer.NewOrderProducer(topic)

	var err error
	for i := 0; i < 10; i++ {
		err = op.Produce([]byte(fmt.Sprintf("Hello Go: %d", i)))
		if err != nil {
			log.Fatalf("failed to produce: %v", err)
		}
		time.Sleep(3 * time.Second)
	}

}
