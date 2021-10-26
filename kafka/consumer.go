package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader *kafka.Reader
}

func (c Consumer) ReadMessage() (kafka.Message, error) {

	return c.reader.ReadMessage(context.Background())
}

func (c Consumer) Config() kafka.ReaderConfig {

	return c.reader.Config()
}
