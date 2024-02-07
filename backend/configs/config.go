package configs

const KafkaTopic string = "alerts"

const KafkaConsumerPort string = ":8080"
const KafkaConsumerGroup string = "alerts-group"

var KafkaBrokers []string = []string{"kafka:9092"}
