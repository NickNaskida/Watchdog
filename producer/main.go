package main

import (
	"fmt"
	"github.com/NickNaskida/Watchdog/producer/cmd"
	"github.com/NickNaskida/Watchdog/producer/services"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Starting the alert producer ...")

	var alertCounter int

	// Create a new Kafka producer
	producer, err := cmd.SetupProducer()
	if err != nil {
		return
	}
	defer producer.Close()

	for {
		randomAlert := services.NewAlert()
		fmt.Println("Alert Message: ", randomAlert.Message)

		// Send the alert message to the Kafka topic
		cmd.SendKafkaMessage(producer, randomAlert)

		// Display total alert events produced
		alertCounter++
		fmt.Println("Total alert events produced: ", alertCounter)

		// Sleep for a random number of seconds
		sleepSeconds := rand.Intn(5)
		time.Sleep(time.Duration(sleepSeconds) * time.Second)
	}
}
