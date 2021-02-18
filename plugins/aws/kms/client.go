package kms

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/cloudquery/cloudquery/database"
	"github.com/cloudquery/cq-provider-aws/resource"
	"go.uber.org/zap"
)

type Client struct {
	db        *database.Database
	log       *zap.Logger
	accountID string
	region    string
	svc       *kms.Client
}

func NewClient(awsConfig aws.Config, db *database.Database, log *zap.Logger,
	accountID string, region string) resource.ClientInterface {
	return &Client{
		db:        db,
		log:       log,
		accountID: accountID,
		region:    region,
		svc:       kms.NewFromConfig(awsConfig, func(options *kms.Options) {
			options.Region = region
		}),
	}
}

func (c *Client) CollectResource(resource string, config interface{}) error {
	switch resource {
	case "keys":
		return c.keys(config)
	default:
		return fmt.Errorf("unsupported resource kms.%s", resource)
	}
}
