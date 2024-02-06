package cmd

import (
	"github.com/IBM/sarama"
	"github.com/NickNaskida/Watchdog/producer/services"
	"log"
)

const kafkaTopic string = "alerts"

var kafkaBrokers []string = []string{"localhost:9092"}

func SetupProducer() (sarama.SyncProducer, error) {
	producer, err := sarama.NewSyncProducer(kafkaBrokers, nil)
	if err != nil {
		log.Fatal("Error creating Kafka producer: ", err)
	}

	return producer, err
}

func SendKafkaMessage(producer sarama.SyncProducer, alertMessage services.Alert) {
	alertMessageJSON, err := alertMessage.ToJSON()
	if err != nil {
		log.Fatal("Error marshalling alert message to JSON: ", err)
	}

	// Create a new Kafka message
	kafkaMessage := &sarama.ProducerMessage{
		Topic: kafkaTopic,
		Value: sarama.StringEncoder(alertMessageJSON),
	}

	// Send the message
	_, _, err = producer.SendMessage(kafkaMessage)
	if err != nil {
		log.Fatal("Error producing Kafka message: ", err)
	}
}
