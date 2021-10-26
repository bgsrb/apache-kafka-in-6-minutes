package kafka

import "github.com/segmentio/kafka-go"

type Optons struct {
	Brokers      []string
	Topic        string
	Partitions   int
	Replications int
	Balancer     kafka.Balancer
}

func WithPartitions(partitions int) func(*Kafka) {
	return func(s *Kafka) {
		s.Config.Partitions = partitions
	}
}

func WithReplications(replications int) func(*Kafka) {
	return func(s *Kafka) {
		s.Config.Replications = replications
	}
}

func WithBalancer(balancer kafka.Balancer) func(*Kafka) {
	return func(s *Kafka) {
		s.Config.Balancer = balancer
	}
}
