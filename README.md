# Apache Kafka in 6 minutes
A quick introduction to how Apache Kafka works and differs from other messaging systems using an example application. In this video I explain partitioning, consumer offsets, replication and many other concepts found in Kafka.

[![Apache Kafka in 6 minutes](http://img.youtube.com/vi/Ch5VhJzaoaI/0.jpg)](https://www.youtube.com/watch?v=Ch5VhJzaoaI "Apache Kafka in 6 minutes")


Example
```go
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
```
Output
```
producer_1 -> America Canada : Querter Start
producer_1 -> Malta Portugal : Foul
producer_1 -> America Canada : Score 39-46
producer_1 -> Brazil Australia : Free Throw
producer_1 -> America Canada : Score 41-46
producer_1 -> Brazil Australia : Querter End
--------------------------------------------
mobile_consumer_2 -> America Canada : Querter Start
mobile_consumer_3 -> Malta Portugal : Foul
mobile_consumer_1 -> America Canada : Score 39-46
mobile_consumer_2 -> Brazil Australia : Free Throw
mobile_consumer_1 -> Brazil Australia : Querter End
mobile_consumer_3 -> America Canada : Score 41-46
computer_consumer_2 -> America Canada : Querter Start
computer_consumer_2 -> Brazil Australia : Free Throw
computer_consumer_3 -> America Canada : Score 39-46
computer_consumer_1 -> Malta Portugal : Foul
computer_consumer_3 -> Brazil Australia : Querter End
computer_consumer_1 -> America Canada : Score 41-46
```
