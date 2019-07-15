package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/Shopify/sarama"
	"github.com/google/uuid"

	"github.com/sirupsen/logrus"
)

const (
	TOPIC_NAME    = "demo_topic"
	PARTITION_NUM = 2
)

func publish(ctx context.Context, gid int, wg *sync.WaitGroup, txnProducer sarama.AsyncProducer) {
	defer wg.Done()
	var enqueued, errors int
	for {
		uuid := uuid.New()
		id := uuid.String()

		//
		transcationMessage := &sarama.ProducerMessage{
			// The value must is encoder to serialize by sarama'encoder
			// Partition number that data will be slipt across kafka cluster
			Value:     sarama.StringEncoder(fmt.Sprintf("transactions-id:%s", id)),
			Topic:     TOPIC_NAME,
			Partition: PARTITION_NUM,
		}
		select {
		// When cancel function is called this case will run and destroy the goroutine
		case <-ctx.Done():
			logrus.Info("the goroutine is canceled")
			logrus.Infof("Enqueued: %d; errors: %d\n", enqueued, errors)
			//shutdown producer
			txnProducer.AsyncClose()
			return
		// This is normal case, message is seen to kafka cluster
		case txnProducer.Input() <- transcationMessage:
			enqueued++
			logrus.Infof("goroutine #%d seen id: %s \n", gid, uuid.String())
		// calculate errors and success
		case err := <-txnProducer.Errors():
			log.Println("Failed to produce message", err)
			errors++
		}

	}
}
func main() {
	// init config for producer
	kafkaConf := sarama.NewConfig()

	kafkaConf.Producer.Compression = sarama.CompressionGZIP
	// partitioner is func, the func will decide the partition of data
	kafkaConf.Producer.Partitioner = sarama.NewRandomPartitioner
	//RequiredAcks the parameter to decide kakfa producer, does producer have to wait for a notification
	//before sending new message
	// to read addition config, you can see at https://godoc.org/github.com/Shopify/sarama#Config
	kafkaConf.Producer.RequiredAcks = sarama.NoResponse
	// the address of broker should be a list of borken
	// ex: []string{"broker1":broke_port1","broker2":broke_port2","brokern":broke_portn"}
	producer, err := sarama.NewAsyncProducer([]string{"localhost:9092"}, kafkaConf)
	if err != nil {
		logrus.Errorf("can't create kafka producer err:%v", err)
	}
	ctx := context.Background()
	// trap Ctrl+C and call cancel on the context
	ctx, cancel := context.WithCancel(ctx)
	go func(ctx context.Context, producer sarama.AsyncProducer) {
		var wg sync.WaitGroup
		for i := 1; i <= 5; i++ {
			wg.Add(1)
			go publish(ctx, i, &wg, producer)
		}
		wg.Wait()
	}(ctx, producer)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	cancel()
	logrus.Info("gracefully shutdown all goroutines")
	time.Sleep(700 * time.Nanosecond)
	logrus.Info("all goroutine is shutdowned ")
}
