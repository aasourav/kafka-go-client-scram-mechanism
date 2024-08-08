package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
)

func main() {
	// Kafka broker address
	kafkaAdd := "my-cluster-kafka-1.my-cluster-kafka-brokers.default.svc:9092"

	mechanism, err := scram.Mechanism(scram.SHA512, "sourav-custom", "aes")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	dialer := &kafka.Dialer{
		Timeout:       10 * time.Second,
		DualStack:     true,
		SASLMechanism: mechanism,
	}

	conn, err := dialer.DialContext(context.Background(), "tcp", kafkaAdd)

	if err != nil {
		fmt.Println("Dial Error: ", err.Error())
		return
	}

	fmt.Println("Success login")
	log.Println(conn.Broker())

	controller, err := conn.Controller()
	if err != nil {
		fmt.Println("Error getting controller: ", err.Error())
		return
	}

	controllerAddr := fmt.Sprintf("%s:%d", controller.Host, controller.Port)
	controllerConn, err := dialer.DialContext(context.Background(), "tcp", controllerAddr)
	if err != nil {
		fmt.Println("Error connecting to controller: ", err.Error())
		return
	}
	defer controllerConn.Close()

	topicConfigs := []kafka.TopicConfig{
		{
			Topic:             "myTopic",
			NumPartitions:     1,
			ReplicationFactor: 1,
		},
	}

	err = controllerConn.CreateTopics(topicConfigs...)
	if err != nil {
		fmt.Println("Error during creating topic: ", err.Error())
		return
	}

	sharedTransport := &kafka.Transport{
		SASL: mechanism,
	}

	w := kafka.Writer{
		Addr:      kafka.TCP(kafkaAdd),
		Topic:     "myTopic",
		Balancer:  &kafka.Hash{},
		Transport: sharedTransport,
	}
	err = w.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte("greetings"),
		Value: []byte("Hello Kafka!"),
	})

	if err != nil {
		fmt.Println("Error to write data: ", err.Error())
		return
	}
	fmt.Println("Success to write message")
	w.Close()
}
