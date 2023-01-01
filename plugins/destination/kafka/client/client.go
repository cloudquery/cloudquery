package client

import (
	"context"
	"crypto/tls"
	"fmt"
	"time"

	"github.com/cloudquery/filetypes/csv"
	"github.com/cloudquery/filetypes/json"
	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/specs"

	"github.com/Shopify/sarama"
	"github.com/rs/zerolog"
)

type Client struct {
	destination.UnimplementedManagedWriter

	conf     *sarama.Config
	producer sarama.SyncProducer

	logger     zerolog.Logger
	spec       specs.Destination
	pluginSpec Spec
	metrics    destination.Metrics

	csvTransformer         *csv.Transformer
	csvReverseTransformer  *csv.ReverseTransformer
	jsonTransformer        *json.Transformer
	jsonReverseTransformer *json.ReverseTransformer
}

func New(ctx context.Context, logger zerolog.Logger, spec specs.Destination) (destination.Client, error) {
	if spec.WriteMode != specs.WriteModeAppend {
		return nil, fmt.Errorf("destination only supports append mode")
	}
	c := &Client{
		logger:                 logger.With().Str("module", "dest-kafka").Logger(),
		csvTransformer:         &csv.Transformer{},
		jsonTransformer:        &json.Transformer{},
		csvReverseTransformer:  &csv.ReverseTransformer{},
		jsonReverseTransformer: &json.ReverseTransformer{},
	}

	c.spec = spec
	if err := spec.UnmarshalSpec(&c.pluginSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
	}
	c.pluginSpec.SetDefaults()
	if err := c.pluginSpec.Validate(); err != nil {
		return nil, err
	}
	if c.pluginSpec.Verbose {
		sarama.Logger = NewSaramaLoggerAdapter(logger)
	}

	c.conf = sarama.NewConfig()
	if c.pluginSpec.MaxMetadataRetries != 0 {
		c.conf.Metadata.Retry.Max = c.pluginSpec.MaxMetadataRetries
	}
	c.conf.Metadata.Retry.Backoff = time.Millisecond * 500
	c.conf.Producer.Retry.Max = 1
	c.conf.Producer.RequiredAcks = sarama.WaitForAll
	c.conf.Producer.Return.Successes = true
	c.conf.Metadata.Full = true
	c.conf.Version = sarama.V1_0_0_0
	c.conf.Metadata.Full = true
	c.conf.ClientID = "cq-destination-kafka-" + c.spec.Name

	if c.pluginSpec.SaslUsername != "" {
		c.conf.Net.SASL.Enable = true
		c.conf.Net.SASL.User = c.pluginSpec.SaslUsername
		c.conf.Net.SASL.Password = c.pluginSpec.SaslPassword
		c.conf.Net.TLS.Enable = true
		c.conf.Net.TLS.Config = &tls.Config{InsecureSkipVerify: true}
		c.conf.Net.SASL.Handshake = true
	}

	var err error
	c.producer, err = sarama.NewSyncProducer(c.pluginSpec.Brokers, c.conf)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Client) Close(ctx context.Context) error {
	return c.producer.Close()
}
