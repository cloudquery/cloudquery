package client

import (
	"bytes"
	"context"
	"time"

	"github.com/Shopify/sarama"
	"github.com/cloudquery/plugin-sdk/schema"
)

const (
	maxWaitTime = 3 * time.Second
)

func (c *Client) Read(ctx context.Context, table *schema.Table, sourceName string, res chan<- []any) error {
	consumer, err := sarama.NewConsumer(c.pluginSpec.Brokers, c.conf)
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
			if err := c.Client.Read(bytes.NewReader(msg.Value), table, sourceName, res); err != nil {
				return err
			}
		case err := <-partitionConsumer.Errors():
			return err.Err
		case <-time.After(maxWaitTime):
			return nil
		}
	}
}
