package procedures

import (
	kafka_config "Eatiplan-Auth/app/kafka"
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func SendNotification(msg map[string]interface{}) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": kafka_config.KafkaHost})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	topic := "send_notification"

	b, _ := json.Marshal(msg)

	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          b,
	}, nil)

	// Wait for message deliveries before shutting down
	p.Flush(100)
}
