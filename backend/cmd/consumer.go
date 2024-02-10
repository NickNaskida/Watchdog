package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/NickNaskida/Watchdog/backend/configs"
	"github.com/NickNaskida/Watchdog/backend/pkg/models"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type Consumer struct{}

func (*Consumer) Setup(sarama.ConsumerGroupSession) error { return nil }

func (*Consumer) Cleanup(sarama.ConsumerGroupSession) error { return nil }

func (consumer *Consumer) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	var alertCount int
	for msg := range claim.Messages() {
		var alert models.Alert
		err := json.Unmarshal(msg.Value, &alert)
		if err != nil {
			log.Printf("failed to unmarshal notification: %v", err)
			continue
		}

		broadcastMessage(alert)

		alertCount++
		fmt.Println("[*] Received alert:", alert.Message, "Total received: ", alertCount)
		sess.MarkMessage(msg, "")
	}
	return nil
}

func initializeConsumerGroup() (sarama.ConsumerGroup, error) {
	config := sarama.NewConfig()

	consumerGroup, err := sarama.NewConsumerGroup(configs.KafkaBrokers, configs.KafkaConsumerGroup, config)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize consumer group: %w", err)
	}

	return consumerGroup, nil
}

func setupConsumerGroup(ctx context.Context) {
	consumerGroup, err := initializeConsumerGroup()
	if err != nil {
		log.Printf("initialization error: %v", err)
	}
	defer consumerGroup.Close()

	consumer := &Consumer{}

	for {
		err = consumerGroup.Consume(ctx, []string{configs.KafkaTopic}, consumer)
		if err != nil {
			log.Printf("error from consumer: %v", err)
		}
		if ctx.Err() != nil {
			return
		}
	}
}

var upgrader = websocket.Upgrader{
	// Allow all origins
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
var clients = make(map[*websocket.Conn]bool)

func Upgrade(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("%s, error while upgrading websocket connection\n", err.Error())
		return
	}
	clients[conn] = true

	fmt.Printf("Websocket connection established: %s\n", conn.RemoteAddr().String())

	defer conn.Close()

	for {
	}
}

func broadcastMessage(message models.Alert) {
	for conn := range clients {
		err := conn.WriteJSON(message)
		if err != nil {
			log.Printf("error while broadcasting message: %v", err)
			delete(clients, conn)
			conn.Close()
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go setupConsumerGroup(ctx)
	defer cancel()
	fmt.Printf("Kafka consumer (group: %s) started ...\n", configs.KafkaConsumerGroup)

	http.HandleFunc("/alerts", Upgrade)

	if err := http.ListenAndServe(configs.KafkaConsumerPort, nil); err != nil {
		log.Printf("failed to run the server: %v", err)
	}
}
