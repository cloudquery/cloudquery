package cloudtrail

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	"github.com/cloudquery/cloudquery/database"
	"github.com/cloudquery/cq-provider-aws/resources/resource"
	"github.com/hashicorp/go-hclog"
)

type Client struct {
	db        *database.Database
	log  hclog.Logger
	accountID string
	region    string
	svc       *cloudtrail.Client
}

func NewClient(awsConfig aws.Config, db *database.Database, log hclog.Logger,
	accountID string, region string) resource.ClientInterface {
	return &Client{
		db:        db,
		log:       log,
		accountID: accountID,
		region:    region,
		svc:       cloudtrail.NewFromConfig(awsConfig, func(o *cloudtrail.Options) {
			o.Region = region
		}),
	}
}

func (c *Client) CollectResource(resource string, config interface{}) error {
	switch resource {
	case "trails":
		return c.trails(config)
	default:
		return fmt.Errorf("unsupported resource autoscaling.%s", resource)
	}
}
