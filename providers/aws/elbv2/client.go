package elbv2

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/cloudquery/cloudquery/providers/aws/resource"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Client struct {
	session          *session.Session
	db               *gorm.DB
	log              *zap.Logger
	accountID        string
	region           string
	resourceMigrated map[string]bool
	svc              *elbv2.ELBV2
}

func NewClient(session *session.Session, db *gorm.DB, log *zap.Logger,
	accountID string, region string) resource.ClientInterface {
	return &Client{
		session:          session,
		db:               db,
		log:              log,
		accountID:        accountID,
		region:           region,
		resourceMigrated: map[string]bool{},
		svc:              elbv2.New(session),
	}
}

func (c *Client) CollectResource(resource string, config interface{}) error {
	switch resource {
	case "load_balancers":
		return c.loadBalancers(config)
	default:
		return fmt.Errorf("unsupported resource elbv2.%s", resource)
	}
}
