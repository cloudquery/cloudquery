package client

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/cloudquery/cloudquery/plugins/destination/mongodb/v2/client/spec"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/writers/batchwriter"
	"github.com/goccy/go-json"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
		var cfg aws.Config
		var err error
		configFns := []func(*config.LoadOptions) error{}
		if c.spec.AWSCredentials.Default {
			// Use default AWS credentials
			cfg, err = config.LoadDefaultConfig(ctx)
			if err != nil {
				return nil, err
			}
		} else {
			configFns = append(configFns, config.WithDefaultRegion("us-east-1"))
			if c.spec.AWSCredentials != nil && c.spec.AWSCredentials.LocalProfile != "" {
				configFns = append(configFns, config.WithSharedConfigProfile(c.spec.AWSCredentials.LocalProfile))
			}

			cfg, err = config.LoadDefaultConfig(ctx, configFns...)
			if err != nil {
				return nil, fmt.Errorf("unable to load AWS SDK config: %w", err)
			}

			if c.spec.AWSCredentials != nil && c.spec.AWSCredentials.RoleARN != "" {
				opts := make([]func(*stscreds.AssumeRoleOptions), 0, 1)
				if c.spec.AWSCredentials.ExternalID != "" {
					opts = append(opts, func(opts *stscreds.AssumeRoleOptions) {
						opts.ExternalID = &c.spec.AWSCredentials.ExternalID
					})
				}
				if c.spec.AWSCredentials.RoleSessionName != "" {
					opts = append(opts, func(opts *stscreds.AssumeRoleOptions) {
						opts.RoleSessionName = c.spec.AWSCredentials.RoleSessionName
					})
				}
				stsClient := sts.NewFromConfig(cfg)
				provider := stscreds.NewAssumeRoleProvider(stsClient, c.spec.AWSCredentials.RoleARN, opts...)

				cfg.Credentials = aws.NewCredentialsCache(provider)
			}
		}
		awsCreds, err := cfg.Credentials.Retrieve(ctx)
		if err != nil {
			return nil, err
		}
		assumeRoleCredential := options.Credential{
			AuthMechanism: "MONGODB-AWS",
			Username:      awsCreds.AccessKeyID,
			Password:      awsCreds.SecretAccessKey,
		}
		if awsCreds.SessionToken != "" {
			assumeRoleCredential.AuthMechanismProperties = map[string]string{
				"AWS_SESSION_TOKEN": awsCreds.SessionToken,
			}
		}

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
