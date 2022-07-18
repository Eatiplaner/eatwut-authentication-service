package procedures

import (
	kafka_config "Eatiplan-Auth/app/kafka"

	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func SendNotification(msg map[string]interface{}) {
	topic := "send_notification"
	partition := 0

	b, _ := json.Marshal(msg)

	conn, err := kafka.DialLeader(context.Background(), "tcp", kafka_config.KafkaHost, topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.Write(b)

	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
