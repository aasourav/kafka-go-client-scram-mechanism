package main

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
)

func main() {
	// Kafka broker address
	brokerAddress := "my-cluster-kafka-1.my-cluster-kafka-brokers.default.svc:9092"
	// Kafka group ID and topic
	topic := "myTopic"

	mechanism, _ := scram.Mechanism(scram.SHA512, "sourav-custom", "aes")

	daler := &kafka.Dialer{
		Timeout:       10 * time.Second,
		DualStack:     true,
		SASLMechanism: mechanism,
	}

	// Create a new reader (consumer)
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{brokerAddress},
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
		Dialer:   daler,
	})

	// Read messages from the topic
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("Error to read data.")
			continue
		}
		fmt.Printf("Received message: %s\n", string(m.Value))
	}
}
