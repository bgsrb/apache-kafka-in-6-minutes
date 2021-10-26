package main

import (
	"encoding/json"
	"fmt"

	k "github.com/segmentio/kafka-go"

	"github.com/sohamkamani/golang-kafka-example/kafka"
)

func main() {

	k := kafka.New([]string{"localhost:9092"}, "test",
		kafka.WithPartitions(3),
		kafka.WithReplications(1),
		kafka.WithBalancer(&k.Hash{}),
	)

	//producers
	k.Producer(produce("producer_1"))

	fmt.Println("--------------------------------------------")

	//consumers
	go k.Consumer("mobile", consume("mobile_consumer_1"))
	go k.Consumer("mobile", consume("mobile_consumer_2"))
	go k.Consumer("mobile", consume("mobile_consumer_3"))

	go k.Consumer("computer", consume("computer_consumer_1"))
	go k.Consumer("computer", consume("computer_consumer_2"))
	go k.Consumer("computer", consume("computer_consumer_3"))

	exit()
}

type Event struct {
	Event string
	Key   string
	Team1 string
	Team2 string
}

func produce(id string) func(producer *kafka.Producer) {

	return func(producer *kafka.Producer) {
		events := []Event{{
			Event: "Querter Start",
			Key:   "America_Canada",
			Team1: "America",
			Team2: "Canada",
		}, {
			Event: "Foul",
			Key:   "Malta_Portugal",
			Team1: "Malta",
			Team2: "Portugal",
		}, {
			Event: "Score 39-46",
			Key:   "America_Canada",
			Team1: "America",
			Team2: "Canada",
		}, {
			Event: "Free Throw",
			Key:   "Brazil_Australia",
			Team1: "Brazil",
			Team2: "Australia",
		}, {
			Event: "Score 41-46",
			Key:   "America_Canada",
			Team1: "America",
			Team2: "Canada",
		}, {
			Event: "Querter End",
			Key:   "Brazil_Australia",
			Team1: "Brazil",
			Team2: "Australia",
		}}
		for _, event := range events {
			value, _ := json.Marshal(event)
			err := producer.Produce(event.Key, value)
			if err != nil {
				fmt.Println("could not write message " + err.Error())
			}
			fmt.Printf("%s -> %s %s : %s\n", id, event.Team1, event.Team2, event.Event)
		}
	}
}

func consume(id string) func(consumer *kafka.Consumer) {
	return func(consumer *kafka.Consumer) {
		for {
			msg, err := consumer.ReadMessage()
			if err != nil {
				fmt.Println("could not read message " + err.Error())
			}
			var event Event
			json.Unmarshal(msg.Value, &event)

			fmt.Printf("%s -> %s %s : %s\n", id, event.Team1, event.Team2, event.Event)
		}
	}
}

func exit() {
	exit := make(chan int)
	<-exit
}
