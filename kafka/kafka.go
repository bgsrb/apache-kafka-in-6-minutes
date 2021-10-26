package kafka

import (
	"net"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
)

type Kafka struct {
	Config Optons
	conn   *kafka.Conn
	broker kafka.Broker
}

func New(brokers []string, topic string, options ...func(*Kafka)) *Kafka {
	var err error

	//Kafka
	k := &Kafka{
		Config: Optons{
			Brokers:      brokers,
			Topic:        topic,
			Partitions:   3,
			Replications: 1,
		},
	}

	//options
	for _, option := range options {
		option(k)
	}

	//conn
	k.conn, err = kafka.Dial("tcp", k.Config.Brokers[0])
	if err != nil {
		panic(err.Error())
	}

	//broker
	k.broker, err = k.conn.Controller()
	if err != nil {
		panic(err.Error())
	}
	k.conn, err = kafka.Dial("tcp", net.JoinHostPort(k.broker.Host, strconv.Itoa(k.broker.Port)))
	if err != nil {
		panic(err.Error())
	}

	//topics
	k.CreateTopics([]kafka.TopicConfig{
		{
			Topic:             k.Config.Topic,
			NumPartitions:     k.Config.Partitions,
			ReplicationFactor: k.Config.Replications,
		},
	})

	return k
}

func (k Kafka) NewConsumer(groupId string) *Consumer {

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: k.Config.Brokers,
		Topic:   k.Config.Topic,
		GroupID: groupId,
		MaxWait: time.Second,
	})

	consumer := &Consumer{
		reader: r,
	}

	return consumer
}

func (k Kafka) Consumer(groupId string, consume func(consumer *Consumer)) *Consumer {

	consumer := k.NewConsumer(groupId)
	consume(consumer)

	return consumer
}

func (k Kafka) NewProducer() *Producer {

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  k.Config.Brokers,
		Topic:    k.Config.Topic,
		Balancer: &kafka.Hash{},
	})

	producer := &Producer{
		writer: w,
	}

	return producer
}

func (k Kafka) Producer(produce func(producer *Producer)) *Producer {

	producer := k.NewProducer()
	produce(producer)

	return producer
}

func (k Kafka) Close() {
	k.conn.Close()
}
