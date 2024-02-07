package main

import (
	"context"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/NickNaskida/Watchdog/backend/configs"
	"github.com/gin-gonic/gin"
	"log"
)

type Consumer struct{}

func (*Consumer) Setup(sarama.ConsumerGroupSession) error   { return nil }
func (*Consumer) Cleanup(sarama.ConsumerGroupSession) error { return nil }

func (consumer *Consumer) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		log.Printf("received message: %s", string(msg.Value))
		//var notification models.Notification
		//err := json.Unmarshal(msg.Value, &notification)
		//if err != nil {
		//	log.Printf("failed to unmarshal notification: %v", err)
		//	continue
		//}

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

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go setupConsumerGroup(ctx)
	defer cancel()

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	fmt.Printf("Kafka consumer (group: %s) started ...\n", configs.KafkaConsumerGroup)

	if err := router.Run(configs.KafkaConsumerPort); err != nil {
		log.Printf("failed to run the server: %v", err)
	}
}
