package kafka

import "github.com/Shopify/sarama"

type SyKafka interface {
	GetConfig() *sarama.Config
	GetProducer()
}
