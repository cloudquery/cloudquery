package ec2

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/cloudquery/cloudquery/providers/aws/internal/resource"
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
	svc              *ec2.EC2
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
		svc:              ec2.New(session),
	}
}

func (c *Client) CollectResource(resource string, config interface{}) error {
	switch resource {
	case "images":
		return c.Images(config)
	case "instances":
		return c.Instances(config)
	case "byoip_cidrs":
		return c.ByoipCidrs(config)
	case "customer_gateways":
		return c.CustomerGateways(config)
	case "internet_gateways":
		return c.InternetGateways(config)
	case "nat_gateways":
		return c.NatGateways(config)
	case "network_acls":
		return c.NetworkAcls(config)
	case "route_tables":
		return c.RouteTables(config)
	case "security_groups":
		return c.SecurityGroups(config)
	case "vpcs":
		return c.Vpcs(config)
	case "subnets":
		return c.Subnets(config)
	default:
		return fmt.Errorf("unsupported resource ec2.%s", resource)
	}
}
