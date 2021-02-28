package organizations

import (
	"fmt"
	"github.com/hashicorp/go-hclog"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/cloudquery/cloudquery/database"
	"github.com/cloudquery/cq-provider-aws/resources/resource"
)

type Client struct {
	db        *database.Database
	log       hclog.Logger
	accountID string
	svc       *organizations.Client
}

func NewClient(awsConfig aws.Config, db *database.Database, log hclog.Logger,
	accountID string, _ string) resource.ClientInterface {
	return &Client{
		db:        db,
		log:       log,
		accountID: accountID,
		svc: organizations.NewFromConfig(awsConfig, func(options *organizations.Options) {
			options.Region = "us-east-1"
		}),
	}
}

func (c *Client) CollectResource(resource string, config interface{}) error {
	switch resource {
	case "accounts":
		return c.accounts(config)
	default:
		return fmt.Errorf("unsupported resource organizations.%s", resource)
	}
}
