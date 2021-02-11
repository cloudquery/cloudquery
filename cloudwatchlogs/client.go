package cloudwatchlogs

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/cloudquery/cloudquery/database"
	"github.com/cloudquery/cq-provider-aws/resource"
	"go.uber.org/zap"
)

type Client struct {
	session   *session.Session
	db        *database.Database
	log       *zap.Logger
	accountID string
	region    string
	svc       *cloudwatchlogs.CloudWatchLogs
}

func NewClient(session *session.Session, awsConfig *aws.Config, db *database.Database, log *zap.Logger,
	accountID string, region string) resource.ClientInterface {
	return &Client{
		session:   session,
		db:        db,
		log:       log,
		accountID: accountID,
		region:    region,
		svc:       cloudwatchlogs.New(session, awsConfig),
	}
}

func (c *Client) CollectResource(resource string, config interface{}) error {
	switch resource {
	case "metric_filters":
		return c.metricFilters(config)
	default:
		return fmt.Errorf("unsupported resource autoscaling.%s", resource)
	}
}
