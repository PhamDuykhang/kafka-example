package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/Shopify/sarama"

	"github.com/sirupsen/logrus"
)

const (
	TOPIC = "demo_topic"
)

func sub(ctx context.Context, gid int, wg *sync.WaitGroup, transactionConsumer sarama.Consumer) {
	defer wg.Done()
	//subscrib topic
	partitionConsumer, err := transactionConsumer.ConsumePartition(TOPIC, 0, sarama.OffsetNewest)
	if err != nil {
		logrus.Errorf("can't subscribe the topic")
		return
	}
	for {
		select {
		// When cancel function is called this case will run and destroy the goroutine
		case <-ctx.Done():
			logrus.Info("the goroutine is canceled")
			partitionConsumer.AsyncClose()
			return
		case data := <-partitionConsumer.Messages():
			logrus.Infof("data: %v \n", string(data.Value))
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

	transactionConsumer, err := sarama.NewConsumer([]string{"localhost:9092"}, consumerConfig)
	if err != nil {
		logrus.Errorf("can't create consumer err: %v", err)
	}
	ctx := context.Background()
	// trap Ctrl+C and call cancel on the context
	ctx, cancel := context.WithCancel(ctx)
	go func(ctx context.Context, transactionConsumer sarama.Consumer) {
		var wg sync.WaitGroup
		for i := 1; i <= 5; i++ {
			wg.Add(1)
			go sub(ctx, i, &wg, transactionConsumer)
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
