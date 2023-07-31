package main

import (
	"encoding/json"
	"fmt"
	"producer/entity"
	"producer/infra/kafka"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	sale := entity.Sale{}
	jsonData, err := json.Marshal(sale.MakeSale())
	if err != nil {
		fmt.Println(err)
	}

	// Create the Kafka producer
	producer := kafka.MakeProducer(&ckafka.ConfigMap{"bootstrap.servers": "localhost:9092"})

	// Send a sample message
	topic := "sales-ingestion"
	err = kafka.SendMessage(producer, topic, jsonData)
	if err != nil {
		fmt.Println("Error sending message:", err)
	}

	// Close the producer when it's no longer needed
	producer.Close()
}
