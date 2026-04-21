package kafka

import (
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"
)

type Producer interface {
	Publish(ctx context.Context, key string, value interface{}) error
}

type producer struct {
	writer *kafka.Writer
}

func NewProducer(brokers []string, topic string) Producer {
	return &producer{
		writer: &kafka.Writer{
			Addr:     kafka.TCP(brokers...),
			Topic:    topic,
			Balancer: &kafka.LeastBytes{},
		},
	}
}

func (p *producer) Publish(ctx context.Context, key string, value interface{}) error {
	data, _ := json.Marshal(value)
	return p.writer.WriteMessages(ctx, kafka.Message{
		Key:   []byte(key),
		Value: data,
	})
}
