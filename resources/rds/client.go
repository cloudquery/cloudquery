package rds

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/cloudquery/cloudquery/database"
	"github.com/cloudquery/cq-provider-aws/resources/resource"
	"github.com/hashicorp/go-hclog"
)

type Client struct {
	db        *database.Database
	log       hclog.Logger
	accountID string
	region    string
	svc       *rds.Client
}

func NewClient(awsConfig aws.Config, db *database.Database, log hclog.Logger,
	accountID string, region string) resource.ClientInterface {
	return &Client{
		db:        db,
		log:       log,
		accountID: accountID,
		region:    region,
		svc: rds.NewFromConfig(awsConfig, func(options *rds.Options) {
			options.Region = region
		}),
	}
}

func (c *Client) CollectResource(ctx context.Context, resource string, config interface{}) error {
	switch resource {
	case "certificates":
		return c.certificates(ctx, config)
	case "clusters":
		return c.clusters(ctx, config)
	case "db_subnet_groups":
		return c.dbSubnetGroups(ctx, config)
	default:
		return fmt.Errorf("unsupported resource rds.%s", resource)
	}
}
