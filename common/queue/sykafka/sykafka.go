package sykafka

import (
	"context"
	"github.com/Shopify/sarama"
	"strings"
)

type EasySyKafka interface {
	GetCfg() *sarama.Config
	InitProducer() error
	SendMessage(message *sarama.ProducerMessage) error
	SendMessages(messages []*sarama.ProducerMessage) error
	GetTopics() []string
	GetServiceName() string
	Consumer(ctx context.Context)
}

var GEasySyKafkaManager *EasySyKafkaManager

func init() {
	GEasySyKafkaManager = new(EasySyKafkaManager)
	GEasySyKafkaManager.db = make(map[string]EasySyKafka)
}

type EasySyKafkaManager struct {
	db map[string]EasySyKafka
}

// Register 暂时无锁
func (e *EasySyKafkaManager) Register(v EasySyKafka) {
	e.db[strings.Join(v.GetTopics(), "-")] = v
}

func (e *EasySyKafkaManager) StartProducer() {

}
