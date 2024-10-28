package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	// Создание продюсера
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		log.Fatalf("Failed to create producer: %s", err)
	}
	defer producer.Close()

	// Создание канала для обработки сигналов
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			value := "Hello, Kafka!"
			// Отправка сообщения в топик "test"
			err := producer.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &[]string{"test"}[0], Partition: kafka.PartitionAny},
				Value:          []byte(value),
			}, nil)

			if err != nil {
				log.Printf("Failed to produce message: %s", err)
			} else {
				log.Printf("Produced message: %s", value)
			}

			// Ждем 1 секунду перед отправкой следующего сообщения
			time.Sleep(10000 * time.Microsecond)
		}
	}()

	// Обработка завершения
	go func() {
		<-sigs
		log.Println("Shutting down producer...")
		producer.Flush(15 * 1000) // Ждем завершения отправки сообщений
		os.Exit(0)
	}()

	// Создание потребителя
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "my-group",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Fatalf("Failed to create consumer: %s", err)
	}
	defer consumer.Close()

	// Подписка на топик "test"
	err = consumer.Subscribe("test", nil)
	if err != nil {
		log.Fatalf("Failed to subscribe to topic: %s", err)
	}

	// Получение сообщений
	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			log.Printf("Consumed message: %s", msg.Value)
		} else {
			log.Printf("Error while consuming message: %s", err)
		}
	}
}
