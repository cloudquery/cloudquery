package directconnect

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	"github.com/cloudquery/cloudquery/database"
	"github.com/cloudquery/cq-provider-aws/resources/resource"
	"github.com/hashicorp/go-hclog"
)

type Client struct {
	db        *database.Database
	log       hclog.Logger
	accountID string
	region    string
	svc       *directconnect.Client
}

func NewClient(awsConfig aws.Config, db *database.Database, log hclog.Logger,
	accountID string, region string) resource.ClientInterface {
	return &Client{
		db:        db,
		log:       log,
		accountID: accountID,
		region:    region,
		svc: directconnect.NewFromConfig(awsConfig, func(options *directconnect.Options) {
			options.Region = region
		}),
	}
}

func (c *Client) CollectResource(resource string, config interface{}) error {
	switch resource {
	case "gateways":
		return c.gateways(config)
	default:
		return fmt.Errorf("unsupported resource directconnect.%s", resource)
	}
}
