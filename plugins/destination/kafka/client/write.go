package client

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"strings"

	"github.com/IBM/sarama"
	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func (c *Client) Write(ctx context.Context, res <-chan message.WriteMessage) error {
	var tables schema.Tables

	messages := make([]*sarama.ProducerMessage, 0, c.spec.BatchSize)
	rows := int64(0)
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
			rows += m.Record.NumRows()
			if rows >= c.spec.BatchSize {
				if err := c.producer.SendMessages(messages); err != nil {
					return err
				}
				// TODO(v4): Increment metrics
				messages, rows = messages[:0], 0
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
			NumPartitions:     int32(c.spec.TopicDetails.NumPartitions),
			ReplicationFactor: int16(c.spec.TopicDetails.ReplicationFactor),
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
