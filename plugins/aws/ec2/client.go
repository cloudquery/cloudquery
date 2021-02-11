package ec2

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
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
	svc       *ec2.EC2
}

func NewClient(session *session.Session, awsConfig *aws.Config, db *database.Database, log *zap.Logger,
	accountID string, region string) resource.ClientInterface {
	return &Client{
		session:   session,
		db:        db,
		log:       log,
		accountID: accountID,
		region:    region,
		svc:       ec2.New(session, awsConfig),
	}
}

func (c *Client) CollectResource(resource string, config interface{}) error {
	switch resource {
	case "images":
		return c.images(config)
	case "instances":
		return c.instances(config)
	case "byoip_cidrs":
		return c.byoipCidrs(config)
	case "customer_gateways":
		return c.customerGateways(config)
	case "internet_gateways":
		return c.internetGateways(config)
	case "nat_gateways":
		return c.natGateways(config)
	case "network_acls":
		return c.networkAcls(config)
	case "route_tables":
		return c.routeTables(config)
	case "security_groups":
		return c.securityGroups(config)
	case "vpcs":
		return c.vpcs(config)
	case "subnets":
		return c.subnets(config)
	case "flow_logs":
		return c.FlowLogs(config)
	case "vpc_peering_connections":
		return c.vpcPeeringConnections(config)
	default:
		return fmt.Errorf("unsupported resource ec2.%s", resource)
	}
}
