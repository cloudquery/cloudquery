package client

import (
	"context"
	"errors"
	"fmt"

	"github.com/cloudquery/cloudquery/plugins/destination/mongodb/v2/client/spec"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/writers/batchwriter"
	"github.com/goccy/go-json"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/auth"
)

type Client struct {
	plugin.UnimplementedSource
	batchwriter.UnimplementedDeleteRecord
	logger zerolog.Logger
	spec   *spec.Spec
	client *mongo.Client
	writer *batchwriter.BatchWriter
}

var errInvalidSpec = errors.New("invalid spec")
var errConnectionFailed = errors.New("failed to connect to MongoDB")

func New(ctx context.Context, logger zerolog.Logger, specByte []byte, _ plugin.NewClientOptions) (plugin.Client, error) {
	var err error
	c := &Client{
		logger: logger.With().Str("module", "mongo-dest").Logger(),
	}
	if err := json.Unmarshal(specByte, &c.spec); err != nil {
		return nil, errors.Join(errInvalidSpec, err)
	}
	if err := c.spec.Validate(); err != nil {
		return nil, errors.Join(errInvalidSpec, err)
	}

	mongoDBClientOptions := options.Client().ApplyURI(c.spec.ConnectionString).SetRegistry(getRegistry())
	if c.spec.AWSCredentials != nil {
		auth.RegisterAuthenticatorFactory(MongoDBCQAWS, NewAuthenticator)
		assumeRoleCredential := options.Credential{
			AuthMechanism: MongoDBCQAWS,
			AuthMechanismProperties: map[string]string{
				"LocalProfile":    c.spec.AWSCredentials.LocalProfile,
				"RoleARN":         c.spec.AWSCredentials.RoleARN,
				"RoleSessionName": c.spec.AWSCredentials.RoleSessionName,
				"ExternalID":      c.spec.AWSCredentials.ExternalID,
				"Default":         fmt.Sprintf("%t", c.spec.AWSCredentials.Default),
			},
		}
		// According to the docs: if ApplyURI is called before SetAuth, the Credential from SetAuth will overwrite the values from the connection string
		mongoDBClientOptions = mongoDBClientOptions.SetAuth(assumeRoleCredential)
	}

	c.client, err = mongo.Connect(context.Background(), mongoDBClientOptions)
	if err != nil {
		return nil, errors.Join(errConnectionFailed, err)
	}
	if err := c.client.Ping(ctx, nil); err != nil {
		return nil, err
	}
	c.writer, err = batchwriter.New(c, batchwriter.WithBatchSize(c.spec.BatchSize), batchwriter.WithBatchSizeBytes(c.spec.BatchSizeBytes), batchwriter.WithLogger(c.logger))
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Client) Close(ctx context.Context) error {
	if err := c.writer.Close(ctx); err != nil {
		_ = c.client.Disconnect(ctx)
		return fmt.Errorf("failed to close writer: %w", err)
	}
	return c.client.Disconnect(ctx)
}
