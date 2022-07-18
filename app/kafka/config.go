package kafka

import (
	"os"
)

var KafkaHost = os.Getenv("KAFKA_BOOTSTRAP_SERVER")
