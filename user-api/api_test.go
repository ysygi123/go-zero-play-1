package main

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"go-zero-play-1/common/symysql"
	"go-zero-play-1/common/syredis"
	user_model "go-zero-play-1/model/mysql/user-model"
	"strconv"
	"sync"
	"testing"
	"time"
)

func Test_apid(t *testing.T) {
	err := symysql.InitSyMysql("root:1qazxsw2@tcp(127.0.0.1:3306)/scrm?charset=utf8mb4&parseTime=True&loc=Local&timeout=1000ms", []string{"root:1qazxsw2@tcp(127.0.0.1:3306)/scrm?charset=utf8mb4&parseTime=True&loc=Local&timeout=1000ms", "scrm:Chanke!2022@tcp(10.255.11.118:3306)/scrm?charset=utf8mb4&parseTime=True&loc=Local&timeout=1000ms"})
	if err != nil {
		fmt.Println(err)
		return
	}

	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			x := new(user_model.ScCorpUser)
			err = symysql.GetDbSession(context.Background()).Table("sc_corp_user").
				Where("id=14").First(&x).Error
			fmt.Println(err)
			fmt.Println(x.Name)
		}()
	}
	wg.Wait()
}

func TestProduceKafka(t *testing.T) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          //follow同步数据后返回
	config.Producer.Partitioner = sarama.NewRandomPartitioner //随机分配分区 partition
	config.Producer.Return.Successes = true
	config.Producer.Interceptors = []sarama.ProducerInterceptor{}
	//当出现消息失败的时候 需要看是 dail fail 哪个名称失败了，然后mac是去 /etc/hosts 里面加上 ip 域名 这样就可以了
	client, err := sarama.NewSyncProducer([]string{"192.168.3.36:9092"}, config)
	if err != nil {
		fmt.Println("建立链接就错误了", err)
		return
	}
	defer client.Close()
	msg := &sarama.ProducerMessage{
		Topic:     "t1",
		Key:       nil,
		Value:     sarama.StringEncoder("有点东西4"),
		Headers:   nil,
		Metadata:  nil,
		Offset:    0,
		Partition: 0,
		Timestamp: time.Time{},
	}
	partition, offset, err := client.SendMessage(msg)
	fmt.Println("查看发送情况", partition, offset, err)

	clientV2, err := sarama.NewAsyncProducer([]string{"192.168.3.36:9092"}, config)
	clientV2.Errors()
	clientV2.Input()
}

func TestConsumeKafka(t *testing.T) {
	consumer, err := sarama.NewConsumer([]string{"192.168.3.36:9092"}, nil)
	if err != nil {
		fmt.Println("链接客户端错误", err)
		return
	}
	partitionList, err := consumer.Partitions("t1")
	if err != nil {
		fmt.Println("获取分区错误", err)
		return
	}
	fmt.Println("分区？", partitionList)
	for _, partition := range partitionList {
		pc, errs := consumer.ConsumePartition("testlog", partition, sarama.OffsetNewest)
		if errs != nil {
			fmt.Println("消费的时候好想有问题", errs)
			continue
		}
		fmt.Printf("我看下pc %+v\n", pc)
		ch := pc.Messages()
		for {
			value, ok := <-ch
			if !ok {
				break
			}
			fmt.Println("接收消息啦 : ", "消息offset", value.Offset, string(value.Value), string(value.Key))
		}
	}
	fmt.Println("直接睡觉了？？离谱吧")
	time.Sleep(100 * time.Second)
}

func Test_Redis(t *testing.T) {
	cf := cache.CacheConf{}
	cf = append(cf, cache.NodeConf{
		RedisConf: redis.RedisConf{
			Host: "127.0.0.1:6380",
			Type: "cluster",
			Pass: "",
			Tls:  false,
		},
		Weight: 100,
	})
	cf = append(cf, cache.NodeConf{
		RedisConf: redis.RedisConf{
			Host: "127.0.0.1:6381",
			Type: "cluster",
			Pass: "",
			Tls:  false,
		},
		Weight: 100,
	})
	cf = append(cf, cache.NodeConf{
		RedisConf: redis.RedisConf{
			Host: "127.0.0.1:6382",
			Type: "cluster",
			Pass: "",
			Tls:  false,
		},
		Weight: 100,
	})
	syredis.InitSyRedis(cf)
	var err error
	for i := int64(5000000); i < 100000000; i++ {
		_, err = syredis.GsyRedis.Zadd("test:big:range", i, strconv.Itoa(int(i)))
		if err != nil {
			fmt.Println(err)
		}
		tt := time.Now()
		if i%10000 == 0 {
			fmt.Println("执行了次数", i, "执行时间", time.Since(tt))
		}
	}
}

func Test_Select(t *testing.T) {
	output := make(chan string, 1)
	ttimer := time.Tick(time.Second)
	go func() {
		for {
			select {
			case output <- "你好":
				time.Sleep(time.Second)
			case <-ttimer:
				fmt.Println("接收定时器哦")
			}
		}
	}()

	for x := range output {
		fmt.Println(x)
	}
}