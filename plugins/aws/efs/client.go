package efs

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/efs"
	"github.com/cloudquery/cloudquery/database"
	"github.com/cloudquery/cq-provider-aws/resource"
	"go.uber.org/zap"
)

type Client struct {
	db        *database.Database
	log       *zap.Logger
	accountID string
	region    string
	svc       *efs.Client
}

func NewClient(awsConfig aws.Config, db *database.Database, log *zap.Logger,
	accountID string, region string) resource.ClientInterface {
	return &Client{
		db:        db,
		log:       log,
		accountID: accountID,
		region:    region,
		svc:       efs.NewFromConfig(awsConfig, func(options *efs.Options) {
			options.Region = region
		}),
	}
}

func (c *Client) CollectResource(resource string, config interface{}) error {
	switch resource {
	case "filesystems":
		return c.fileSystems(config)
	default:
		return fmt.Errorf("unsupported resource ecr.%s", resource)
	}
}
