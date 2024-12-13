package client

import (
	"bytes"
	"context"
	"time"

	"github.com/IBM/sarama"
	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

const (
	maxWaitTime = 3 * time.Second
)

func (c *Client) Read(ctx context.Context, table *schema.Table, res chan<- arrow.Record) error {
	consumer, err := sarama.NewConsumer(c.spec.Brokers, c.conf)
	if err != nil {
		return err
	}
	defer consumer.Close()
	partitionConsumer, err := consumer.ConsumePartition(table.Name, 0, sarama.OffsetOldest)
	if err != nil {
		return err
	}
	defer partitionConsumer.Close()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case msg := <-partitionConsumer.Messages():
			if err := c.Client.Read(bytes.NewReader(msg.Value), table, res); err != nil {
				return err
			}
		case err := <-partitionConsumer.Errors():
			return err.Err
		case <-time.After(maxWaitTime):
			return nil
		}
	}
}
