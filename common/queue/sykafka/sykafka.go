package sykafka

import "github.com/Shopify/sarama"

type EasySyKafka interface {
	GetCfg() *sarama.Config
	SendMessage(message *sarama.ProducerMessage) error
	SendMessages(messages []*sarama.ProducerMessage)
}

type EasySyKafkaManager struct {
	db map[string]EasySyKafka
}

// Register 暂时无锁
func (e *EasySyKafkaManager) Register(key string, v EasySyKafka) {
	e.db[key] = v
}
