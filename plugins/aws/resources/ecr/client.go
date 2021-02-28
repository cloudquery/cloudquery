package ecr

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/cloudquery/cloudquery/database"
	"github.com/cloudquery/cq-provider-aws/resources/resource"
	"github.com/hashicorp/go-hclog"
)

type Client struct {
	db        *database.Database
	log       hclog.Logger
	accountID string
	region    string
	svc       *ecr.Client
}

func NewClient(awsConfig aws.Config, db *database.Database, log hclog.Logger,
	accountID string, region string) resource.ClientInterface {
	return &Client{
		db:        db,
		log:       log,
		accountID: accountID,
		region:    region,
		svc: ecr.NewFromConfig(awsConfig, func(options *ecr.Options) {
			options.Region = region
		}),
	}
}

func (c *Client) CollectResource(resource string, config interface{}) error {
	switch resource {
	case "images":
		return c.images(config)
	default:
		return fmt.Errorf("unsupported resource ecr.%s", resource)
	}
}
