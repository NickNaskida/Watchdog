package main

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/NickNaskida/Watchdog/backend/configs"
	"github.com/NickNaskida/Watchdog/backend/pkg/models"
	"github.com/NickNaskida/Watchdog/backend/services"
	"log"
	"math/rand"
	"time"
)

func SetupProducer() (sarama.SyncProducer, error) {
	producer, err := sarama.NewSyncProducer(configs.KafkaBrokers, nil)
	if err != nil {
		log.Fatal("Error creating Kafka producer: ", err)
	}

	return producer, err
}

func SendKafkaMessage(producer sarama.SyncProducer, alertMessage models.Alert) {
	alertMessageJSON, err := json.Marshal(alertMessage)
	if err != nil {
		log.Fatal("Error marshalling alert message to JSON: ", err)
	}

	// Create a new Kafka message
	kafkaMessage := &sarama.ProducerMessage{
		Topic: configs.KafkaTopic,
		Value: sarama.StringEncoder(alertMessageJSON),
	}

	// Send the message
	_, _, err = producer.SendMessage(kafkaMessage)
	if err != nil {
		log.Fatal("Error producing Kafka message: ", err)
	}
}

func main() {
	// Create a new Kafka producer
	producer, err := SetupProducer()
	if err != nil {
		return
	}
	defer producer.Close()

	fmt.Printf("Kafka producer started publishing to topic: %s\n", configs.KafkaTopic)

	var alertCounter int
	for {
		randomAlert := services.NewAlert()
		fmt.Println("[*] Alert Message: ", randomAlert.Message)

		// Send the alert message to the Kafka topic
		SendKafkaMessage(producer, randomAlert)

		// Display total alert events produced
		alertCounter++
		fmt.Println("[*] Total alert events produced: ", alertCounter)

		// Sleep for a random number of seconds
		sleepSeconds := rand.Intn(5)
		time.Sleep(time.Duration(sleepSeconds) * time.Second)
	}
}
