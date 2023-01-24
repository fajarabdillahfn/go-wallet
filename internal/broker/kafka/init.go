package kafka

import (
	"log"

	"github.com/Shopify/sarama"
	"github.com/lovoo/goka"
)

var (
	brokers                         = []string{"localhost:9092"}
	topic               goka.Stream = "deposits"
	aboveThresholdGroup goka.Group  = "aboveThreshold"
	balanceGroup        goka.Group  = "balance"

	tmc *goka.TopicManagerConfig
)

func Init() {
	tmc = goka.NewTopicManagerConfig()
	tmc.Table.Replication = 1
	tmc.Stream.Replication = 1

	config := goka.DefaultConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	goka.ReplaceGlobalConfig(config)

	tm, err := goka.NewTopicManager(brokers, goka.DefaultConfig(), tmc)
	if err != nil {
		log.Fatalf("Error creating topic manager: %v", err)
	}
	defer tm.Close()
	err = tm.EnsureStreamExists(string(topic), 8)
	if err != nil {
		log.Printf("Error creating kafka topic %s: %v", topic, err)
	}
}
