package cloudwatch

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/cloudquery/cloudquery/database"
	"github.com/cloudquery/cq-provider-aws/resources/resource"
	"github.com/hashicorp/go-hclog"
)

type Client struct {
	db        *database.Database
	log  hclog.Logger
	accountID string
	region    string
	svc       *cloudwatch.Client
}

func NewClient(awsConfig aws.Config, db *database.Database, log hclog.Logger,
	accountID string, region string) resource.ClientInterface {
	return &Client{
		db:        db,
		log:       log,
		accountID: accountID,
		region:    region,
		svc:       cloudwatch.NewFromConfig(awsConfig, func(o *cloudwatch.Options) {
			o.Region = region
		}),
	}
}

func (c *Client) CollectResource(resource string, config interface{}) error {
	switch resource {
	case "alarms":
		return c.alarms(config)
	default:
		return fmt.Errorf("unsupported resource autoscaling.%s", resource)
	}
}
