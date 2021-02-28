package elbv2

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	elbv2 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/cloudquery/cloudquery/database"
	"github.com/cloudquery/cq-provider-aws/resources/resource"
	"github.com/hashicorp/go-hclog"
)

type Client struct {
	db        *database.Database
	log       hclog.Logger
	accountID string
	region    string
	svc       *elbv2.Client
}

func NewClient(awsConfig aws.Config, db *database.Database, log hclog.Logger,
	accountID string, region string) resource.ClientInterface {
	return &Client{
		db:        db,
		log:       log,
		accountID: accountID,
		region:    region,
		svc: elbv2.NewFromConfig(awsConfig, func(options *elbv2.Options) {
			options.Region = region
		}),
	}
}

func (c *Client) CollectResource(resource string, config interface{}) error {
	switch resource {
	case "load_balancers":
		return c.loadBalancers(config)
	case "target_groups":
		return c.targetGroups(config)
	default:
		return fmt.Errorf("unsupported resource elbv2.%s", resource)
	}
}
