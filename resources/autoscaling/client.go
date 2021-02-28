package autoscaling

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/cloudquery/cloudquery/database"
	"github.com/cloudquery/cq-provider-aws/resources/resource"
	"github.com/hashicorp/go-hclog"
)

type Client struct {
	db        *database.Database
	log  hclog.Logger
	accountID string
	region    string
	svc       *autoscaling.Client
}

func NewClient(awsConfig aws.Config, db *database.Database, log hclog.Logger,
	accountID string, region string) resource.ClientInterface {
	return &Client{
		db:        db,
		log:       log,
		accountID: accountID,
		region:    region,
		svc:       autoscaling.NewFromConfig(awsConfig, func(o *autoscaling.Options) {
			o.Region = region
		}),
	}
}

func (c *Client) CollectResource(resource string, config interface{}) error {
	switch resource {
	case "launch_configurations":
		return c.launchConfigurations(config)
	default:
		return fmt.Errorf("unsupported resource autoscaling.%s", resource)
	}
}
