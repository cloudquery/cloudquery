package client

import (
	"bufio"
	"bytes"
	"context"
	"strings"
	"sync/atomic"

	"github.com/Shopify/sarama"
	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/schema"
)

func (c *Client) createTopics(_ context.Context, tables schema.Tables) error {
	c.conf.Version = sarama.V2_0_0_0
	admin, err := sarama.NewClusterAdmin(c.pluginSpec.Brokers, c.conf)
	if err != nil {
		return err
	}
	defer admin.Close()
	for _, table := range tables.FlattenTables() {
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

func (c *Client) Write(ctx context.Context, tables schema.Tables, res <-chan *destination.ClientResource) error {
	if err := c.createTopics(ctx, tables); err != nil {
		return err
	}

	messages := make([]*sarama.ProducerMessage, 0, c.spec.BatchSize)
	for r := range res {
		var b bytes.Buffer
		w := bufio.NewWriter(&b)
		table := tables.Get(r.TableName)
		if err := c.Client.WriteTableBatchFile(w, table, [][]any{r.Data}); err != nil {
			return err
		}
		w.Flush()
		messages = append(messages, &sarama.ProducerMessage{
			Topic: r.TableName,
			Key:   nil,
			Value: sarama.ByteEncoder(b.Bytes()),
		})
		if len(messages) >= c.spec.BatchSize {
			if err := c.producer.SendMessages(messages); err != nil {
				return err
			}
			atomic.AddUint64(&c.metrics.Writes, uint64(c.spec.BatchSize))
			messages = make([]*sarama.ProducerMessage, 0, c.spec.BatchSize)
		}
	}

	if len(messages) > 0 {
		if err := c.producer.SendMessages(messages); err != nil {
			return err
		}
		atomic.AddUint64(&c.metrics.Writes, uint64(len(messages)))
	}

	return nil
}
