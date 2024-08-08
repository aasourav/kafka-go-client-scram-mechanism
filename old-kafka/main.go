package main

import (
	"fmt"
	"time"

	"github.com/aasourav/go-kafka/kafkaconf"
)

func main() {
	go kafkaconf.StartKafka()
	fmt.Println("Kafka started")
	time.Sleep(10 * time.Second)
}
