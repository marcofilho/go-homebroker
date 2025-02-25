package kafka

import ckakfa "github.com/confluentinc/confluent-kafka-go/kafka"

type Consumer struct {
	ConfigMap *ckakfa.ConfigMap
	Topics    []string
}

func NewConsumer(configMap *ckakfa.ConfigMap, topics []string) *Consumer {
	return &Consumer{
		ConfigMap: configMap,
		Topics:    topics,
	}
}

func (c *Consumer) Consume(msgChan chan *ckakfa.Message) error {
	consumer, err := ckakfa.NewConsumer(c.ConfigMap)
	if err != nil {
		panic(err)
	}

	err = consumer.SubscribeTopics(c.Topics, nil)
	if err != nil {
		panic(err)
	}

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			msgChan <- msg
		}
	}
}
