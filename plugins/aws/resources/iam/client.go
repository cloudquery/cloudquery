package iam

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/cloudquery/cloudquery/database"
	"github.com/cloudquery/cq-provider-aws/resources/resource"
	"github.com/hashicorp/go-hclog"
)

type Client struct {
	db        *database.Database
	log       hclog.Logger
	accountID string
	svc       *iam.Client
}

func NewClient(awsConfig aws.Config, db *database.Database, log hclog.Logger,
	accountID string, _ string) resource.ClientInterface {
	return &Client{
		db:        db,
		log:       log,
		accountID: accountID,
		svc: iam.NewFromConfig(awsConfig, func(options *iam.Options) {
			options.Region = "us-east-1"
		}),
	}
}

func (c *Client) CollectResource(resource string, config interface{}) error {
	switch resource {
	case "users":
		return c.users(config)
	case "groups":
		return c.groups(config)
	case "policies":
		return c.policies(config)
	case "roles":
		return c.roles(config)
	case "password_policies":
		return c.passwordPolicies(config)
	case "virtual_mfa_devices":
		return c.virtualMFADevices(config)
	default:
		return fmt.Errorf("unsupported resource iam.%s", resource)
	}
}
