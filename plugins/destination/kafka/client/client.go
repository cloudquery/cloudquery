package client

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"time"

	"github.com/IBM/sarama"
	"github.com/cloudquery/cloudquery/plugins/destination/kafka/client/spec"
	internalPlugin "github.com/cloudquery/cloudquery/plugins/destination/kafka/resources/plugin"
	"github.com/cloudquery/filetypes/v4"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/rs/zerolog"
)

type Client struct {
	plugin.UnimplementedSource

	conf     *sarama.Config
	producer sarama.SyncProducer

	logger zerolog.Logger
	spec   *spec.Spec

	*filetypes.Client
}

func New(_ context.Context, logger zerolog.Logger, s []byte, opts plugin.NewClientOptions) (plugin.Client, error) {
	c := &Client{
		logger: logger.With().Str("module", "dest-kafka").Logger(),
	}
	if opts.NoConnection {
		return c, nil
	}

	if err := json.Unmarshal(s, &c.spec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
	}
	if err := c.spec.Validate(); err != nil {
		return nil, err
	}
	c.spec.SetDefaults()

	if c.spec.Verbose {
		sarama.Logger = NewSaramaLoggerAdapter(logger)
	}

	c.conf = sarama.NewConfig()
	c.conf.Metadata.Retry.Backoff = time.Millisecond * 500
	c.conf.Producer.Retry.Max = 1
	c.conf.Producer.RequiredAcks = sarama.WaitForAll
	c.conf.Producer.Return.Successes = true
	c.conf.Metadata.Full = true
	c.conf.Version = sarama.V1_0_0_0
	c.conf.Metadata.Full = true
	c.conf.ClientID = `cwc|1c04a227-aef8-47a9-9353-e20bbb6a9616|cq-destination-kafka|` + internalPlugin.Version

	if c.spec.SASLUsername != "" {
		c.conf.Net.SASL.Enable = true
		c.conf.Net.SASL.User = c.spec.SASLUsername
		c.conf.Net.SASL.Password = c.spec.SASLPassword
		c.conf.Net.TLS.Enable = true
		c.conf.Net.TLS.Config = &tls.Config{InsecureSkipVerify: true}
		c.conf.Net.SASL.Handshake = true
	}

	var err error
	c.producer, err = sarama.NewSyncProducer(c.spec.Brokers, c.conf)
	if err != nil {
		return nil, err
	}

	filetypesClient, err := filetypes.NewClient(&c.spec.FileSpec)
	if err != nil {
		return nil, fmt.Errorf("failed to create filetypes client: %w", err)
	}
	c.Client = filetypesClient

	return c, nil
}

func (c *Client) Close(_ context.Context) error {
	return c.producer.Close()
}
