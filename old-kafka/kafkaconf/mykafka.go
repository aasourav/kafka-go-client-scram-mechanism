package kafkaconf

import (
	"context"
	"fmt"

	kafka "github.com/segmentio/kafka-go"
)

func StartKafka() {

	conf := kafka.ReaderConfig{
		Brokers:  []string{"localhost:9090"},
		Topic:    "myTopic",
		MaxBytes: 50,
	}
	_, err := kafka.Dial("tcp", "localhost:9092")

	fmt.Println(err)

	reader := kafka.NewReader(conf)

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Println(err.Error())
			break
		}

		fmt.Println(string(msg.Value))
	}

}
