package kafka

import (
	"github.com/segmentio/kafka-go"
)

func (k Kafka) CreateTopics(topics []kafka.TopicConfig) {

	err := k.conn.CreateTopics(topics...)
	if err != nil {
		panic(err.Error())
	}
}
