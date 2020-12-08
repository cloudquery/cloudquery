package ecr

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
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
	svc              *ecr.ECR
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
		svc:              ecr.New(session),
	}
}

func (c *Client) CollectResource(resource string, config interface{}) error {
	switch resource {
	case "images":
		return c.imageIdentifiers(config)
	default:
		return fmt.Errorf("unsupported resource ecr.%s", resource)
	}
}
