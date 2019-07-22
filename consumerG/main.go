package main

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"sync"
	"time"
)

const (
	TOPIC1 = "notifications"
	TOPIC2 = "bulk"
)

type ConsumerGroupHandlerImp struct {
	OutChan chan string
	//StopCh chan  struct{}
}

func NewConsumerGroupeImp(a chan string) ConsumerGroupHandlerImp {
	return ConsumerGroupHandlerImp{
		OutChan: a,
		//StopCh:stop,
	}
}
func (ch ConsumerGroupHandlerImp) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (ch ConsumerGroupHandlerImp) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (ch ConsumerGroupHandlerImp) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case data := <-claim.Messages():
			//ch.OutChan <- fmt.Sprintf("data from : %s : %s",data.Topic,string(data.Value))
			logrus.Infof("data from : %s : %s", data.Topic, string(data.Value))
		}
	}
}

func main() {
	// to read addition config, you can see at https://godoc.org/github.com/Shopify/sarama#Config
	consumerConfig := sarama.NewConfig()
	consumerConfig.Consumer.MaxWaitTime = time.Duration(1 * time.Minute)
	// enable error chanel to count error message
	consumerConfig.Consumer.Return.Errors = true
	consumerConfig.Consumer.Offsets.Retry.Max = 10
	consumerConfig.Version = sarama.V2_0_1_0

	transactionConsumer, err := sarama.NewConsumerGroup([]string{"localhost:9092"}, "krp", consumerConfig)
	if err != nil {
		logrus.Errorf("can't create consumer err: %v", err)
		return
	}
	//defer func() { _ = transactionConsumer.Close() }()
	//go func() {
	//	for err := range transactionConsumer.Errors() {
	//		fmt.Println("ERROR", err)
	//	}
	//}()

	//if err != nil {
	//	logrus.Errorf("can't consume err: %v", err)
	//	return
	//}

	//a := make(chan string)
	//b := make(chan struct{})
	ctx := context.Background()
	// trap Ctrl+C and call cancel on the context

	ctx, cancel := context.WithCancel(ctx)
	go func(ctx context.Context, consumerG sarama.ConsumerGroup) {
		var wg sync.WaitGroup
		for i := 1; i <= 10; i++ {
			wg.Add(1)
			consume(ctx, wg, i, consumerG)
		}
		wg.Wait()
	}(ctx, transactionConsumer)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	cancel()
	logrus.Info("gracefully shutdown all goroutines")
	time.Sleep(700 * time.Nanosecond)
	logrus.Info("all goroutine is shutdowned ")

}

func consume(ctx context.Context, wg sync.WaitGroup, wid int, consume sarama.ConsumerGroup) {
	defer wg.Done()
	handler := ConsumerGroupHandlerImp{}
	for {
		select {
		case <-ctx.Done():
			logrus.Info("the goroutine is canceled")
			//shutdown producer
			//wait to ensure data is seen before closing producer
			time.Sleep(5 * time.Second)
			consume.Close()
			return
		default:
			err := consume.Consume(ctx, []string{TOPIC1, TOPIC2}, handler)
			if err != nil {
				panic(err)
			}
		}
	}
}
