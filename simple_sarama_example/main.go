package main

import (
	"fmt"
	"time"

	"github.com/Shopify/sarama"
)

func main() {
	brokers := []string{"172.22.0.4:19092", "172.22.0.5:29092"}
	topic := "foo"

	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		panic(err)
	}

	// partition 0 and offset 0
	p1, err := consumer.ConsumePartition(topic, 0, 0)
	if err != nil {
		panic(err)
	}

	// partition 1 and offset 0
	p2, err := consumer.ConsumePartition(topic, 1, 0)
	if err != nil {
		panic(err)
	}

	go func() {
		for m := range p1.Messages() {
			fmt.Println(m.Partition, m.Key, string(m.Value))
		}
	}()

	go func() {
		for m := range p2.Messages() {
			fmt.Println(m.Partition, m.Key, string(m.Value))
		}
	}()

	time.Sleep(100 * time.Second)

}
