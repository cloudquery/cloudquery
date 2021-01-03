package s3

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/cloudquery/cloudquery/database"
	"github.com/cloudquery/cloudquery/providers/aws/resource"
	"go.uber.org/zap"
)

type Client struct {
	session   *session.Session
	db        *database.Database
	log       *zap.Logger
	accountID string
	region    string
	svc       *s3.S3
	awsConfig *aws.Config
}

func NewClient(sess *session.Session, awsConfig *aws.Config, db *database.Database, log *zap.Logger,
	accountID string, _ string) resource.ClientInterface {
	globalRegion := "us-east-1"
	awsConfig.Region = &globalRegion
	return &Client{
		session:   sess,
		db:        db,
		log:       log,
		accountID: accountID,
		region:    "us-east-1",
		svc:       s3.New(sess, awsConfig),
		awsConfig: awsConfig,
	}
}

func (c *Client) CollectResource(resource string, config interface{}) error {
	switch resource {
	case "buckets":
		return c.buckets(config)
	default:
		return fmt.Errorf("unsupported resource buckets.%s", resource)
	}
}
