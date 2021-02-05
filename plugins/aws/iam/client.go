package iam

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/cloudquery/cloudquery/database"
	"github.com/cloudquery/cloudquery/providers/aws/resource"
	"go.uber.org/zap"
)

type Client struct {
	session   *session.Session
	db        *database.Database
	log       *zap.Logger
	accountID string
	svc       *iam.IAM
}

func NewClient(session *session.Session, awsConfig *aws.Config, db *database.Database, log *zap.Logger,
	accountID string, _ string) resource.ClientInterface {
	return &Client{
		session:   session,
		db:        db,
		log:       log,
		accountID: accountID,
		svc:       iam.New(session, awsConfig),
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
