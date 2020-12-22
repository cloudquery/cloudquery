package iam

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/cloudquery/cloudquery/providers/aws/resource"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Client struct {
	session   *session.Session
	db        *gorm.DB
	log       *zap.Logger
	accountID string
	svc       *iam.IAM
}

func NewClient(session *session.Session, awsConfig *aws.Config, db *gorm.DB, log *zap.Logger,
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
	default:
		return fmt.Errorf("unsupported resource iam.%s", resource)
	}
}
