package service

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"go-zero-play-1/common/queue/sykafka"
	"go-zero-play-1/common/utils"
	"sync"
)

type Fsp struct {
	producer sarama.SyncProducer
}

func init() {
	fsp := new(Fsp)
	sykafka.GEasySyKafkaManager.Register(fsp)
}

func (f *Fsp) GetTopics() []string {
	return []string{"t1"}
}

func (f *Fsp) GetServiceName() string {
	return "t1"
}

func (f *Fsp) GetCfg() *sarama.Config {
	cfg := sarama.NewConfig()
	cfg.Producer.RequiredAcks = sarama.WaitForAll          //follow同步数据后返回
	cfg.Producer.Partitioner = sarama.NewRandomPartitioner //随机分配分区 partition
	cfg.Producer.Return.Successes = true
	cfg.Producer.Interceptors = []sarama.ProducerInterceptor{}

	cfg.Consumer.Offsets.Initial = sarama.OffsetOldest
	return cfg
}

func (f *Fsp) InitProducer() error {
	var err error
	f.producer, err = sarama.NewSyncProducer([]string{"192.168.15.38:9092"}, f.GetCfg())
	if err != nil {
		return err
	}
	return nil
}

func (f *Fsp) SendMessage(message *sarama.ProducerMessage) error {
	partition, offset, err := f.producer.SendMessage(message)
	fmt.Println("partion, offset", partition, offset)
	return err
}

func (f *Fsp) SendMessages(messages []*sarama.ProducerMessage) error {
	err := f.producer.SendMessages(messages)
	return err
}

func (f *Fsp) GetConsumerNum() int {
	return 5
}

func (f *Fsp) Consumer(ctx context.Context) {
	fmt.Println("开始消费")
	consumerGroup, err := sarama.NewConsumerGroup([]string{"192.168.15.38:9092"}, "gpid", f.GetCfg())
	if err != nil {
		return
	}
	fmt.Println("消费函数")
	wg := sync.WaitGroup{}
	for i := 0; i < f.GetConsumerNum(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err = consumerGroup.Consume(ctx, f.GetTopics(), &easy{})
			fmt.Println(err)
		}()
	}
	wg.Wait()
}

type easy struct{}

func (easy) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (easy) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (e easy) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	fmt.Println("哎嘿嘿", claim.Messages())
	for msg := range claim.Messages() {
		fmt.Printf("Message topic:%q partition:%d offset:%d, value:%s\n", msg.Topic, msg.Partition, msg.Offset, utils.B2S(msg.Value))
		sess.MarkMessage(msg, "")
	}
	return nil
}
