package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"github.com/Shopify/sarama"
)

const (
	TOPIC = "notifications"
)

func main() {
	kafkaConf := sarama.NewConfig()
	kafkaConf.Producer.RequiredAcks = sarama.WaitForAll
	kafkaConf.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, kafkaConf)
	if err != nil {
		logrus.Errorf("can't create producer err: %v ", err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			logrus.Errorf("can't close producer ")
		}
	}()
	for {
		uuid := uuid.New()
		id := uuid.String()
		msg := &sarama.ProducerMessage{
			Topic: TOPIC,
			Value: sarama.StringEncoder(fmt.Sprintf("notifications-transcation-id-sync:%v", id)),
		}
		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			logrus.Errorf("FAILED to send message: %s\n", err)
		} else {
			logrus.Infof("> message sent to partition %d at offset %d\n", partition, offset)
		}
		//time.Sleep(2 * time.Millisecond)
	}

}
func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
