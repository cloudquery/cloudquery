package client

import (
	"bufio"
	"bytes"
	"context"
	"strings"
	"sync/atomic"

	"github.com/Shopify/sarama"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v2/schema"
)

func (c *Client) createTopics(_ context.Context, tables schema.Schemas) error {
	c.conf.Version = sarama.V2_0_0_0
	admin, err := sarama.NewClusterAdmin(c.pluginSpec.Brokers, c.conf)
	if err != nil {
		return err
	}
	defer admin.Close()
	for _, table := range tables {
		tableName := schema.TableName(table)
		err := admin.CreateTopic(tableName, &sarama.TopicDetail{
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

func (c *Client) Write(ctx context.Context, tables schema.Schemas, res <-chan arrow.Record) error {
	if err := c.createTopics(ctx, tables); err != nil {
		return err
	}

	messages := make([]*sarama.ProducerMessage, 0, c.spec.BatchSize)
	for r := range res {
		var b bytes.Buffer
		w := bufio.NewWriter(&b)
		sc := r.Schema()
		tableName := schema.TableName(sc)
		if err := c.Client.WriteTableBatchFile(w, sc, []arrow.Record{r}); err != nil {
			return err
		}
		w.Flush()
		messages = append(messages, &sarama.ProducerMessage{
			Topic: tableName,
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
