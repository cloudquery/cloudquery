package client

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"strings"

	"github.com/Shopify/sarama"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func (c *Client) Write(ctx context.Context, res <-chan message.WriteMessage) error {
	var tables schema.Tables

	messages := make([]*sarama.ProducerMessage, 0, c.spec.BatchSize)
	for r := range res {
		switch m := r.(type) {
		case *message.WriteMigrateTable:
			tables = append(tables, m.Table)
		case *message.WriteDeleteStale:
			continue
		case *message.WriteInsert:
			if err := c.createTopics(ctx, tables); err != nil {
				return fmt.Errorf("failed to create topics: %w", err)
			}
			tables = nil

			table := m.GetTable()
			var b bytes.Buffer
			w := bufio.NewWriter(&b)
			if err := c.Client.WriteTableBatchFile(w, table, []arrow.Record{m.Record}); err != nil {
				return err
			}
			if err := w.Flush(); err != nil {
				return fmt.Errorf("failed to flush buffer: %w", err)
			}

			messages = append(messages, &sarama.ProducerMessage{
				Topic: table.Name,
				Key:   nil,
				Value: sarama.ByteEncoder(b.Bytes()),
			})
			if len(messages) >= c.spec.BatchSize {
				if err := c.producer.SendMessages(messages); err != nil {
					return err
				}
				// TODO(v4): Increment metrics
				messages = messages[:0]
			}

		default:
			return fmt.Errorf("unhandled message type: %T", m)
		}
	}

	if len(messages) > 0 {
		if err := c.producer.SendMessages(messages); err != nil {
			return err
		}
		// TODO(v4): Increment metrics
	}

	return nil
}

func (c *Client) createTopics(_ context.Context, tables schema.Tables) error {
	c.conf.Version = sarama.V2_0_0_0
	admin, err := sarama.NewClusterAdmin(c.spec.Brokers, c.conf)
	if err != nil {
		return err
	}
	defer admin.Close()
	for _, table := range tables {
		err := admin.CreateTopic(table.Name, &sarama.TopicDetail{
			NumPartitions:     1,
			ReplicationFactor: 1,
		}, false)
		if err != nil {
			if strings.Contains(err.Error(), "Topic with this name already exists") {
				continue
			}
			return err
		}
	}
	return nil
}
