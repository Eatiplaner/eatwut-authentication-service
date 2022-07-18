package kafka_config

import (
	"os"
)

var KafkaHost = os.Getenv("KAFKA_BOOTSTRAP_SERVER")
