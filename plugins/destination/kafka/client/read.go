package client

import (
	"bytes"
	"context"
	"time"

	"github.com/Shopify/sarama"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

const (
	maxWaitTime = 3 * time.Second
)

func (c *Client) Read(ctx context.Context, table *arrow.Schema, sourceName string, res chan<- arrow.Record) error {
	consumer, err := sarama.NewConsumer(c.pluginSpec.Brokers, c.conf)
	if err != nil {
		return err
	}
	defer consumer.Close()
	tableName := schema.TableName(table)
	partitionConsumer, err := consumer.ConsumePartition(tableName, 0, sarama.OffsetOldest)
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
