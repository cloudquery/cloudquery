package ec2

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/cloudquery/cloudquery/database"
	"github.com/cloudquery/cq-provider-aws/resources/resource"
	"github.com/hashicorp/go-hclog"
)

type Client struct {
	db        *database.Database
	log       hclog.Logger
	accountID string
	region    string
	svc       *ec2.Client
}

func NewClient(awsConfig aws.Config, db *database.Database, log hclog.Logger,
	accountID string, region string) resource.ClientInterface {
	return &Client{
		db:        db,
		log:       log,
		accountID: accountID,
		region:    region,
		svc: ec2.NewFromConfig(awsConfig, func(options *ec2.Options) {
			options.Region = region
		}),
	}
}

func (c *Client) CollectResource(ctx context.Context, resource string, config interface{}) error {
	switch resource {
	case "images":
		return c.images(ctx, config)
	case "instances":
		return c.instances(ctx, config)
	case "byoip_cidrs":
		return c.byoipCidrs(ctx, config)
	case "customer_gateways":
		return c.customerGateways(ctx, config)
	case "internet_gateways":
		return c.internetGateways(ctx, config)
	case "nat_gateways":
		return c.natGateways(ctx, config)
	case "network_acls":
		return c.networkAcls(ctx, config)
	case "route_tables":
		return c.routeTables(ctx, config)
	case "security_groups":
		return c.securityGroups(ctx, config)
	case "vpcs":
		return c.vpcs(ctx, config)
	case "subnets":
		return c.subnets(ctx, config)
	case "flow_logs":
		return c.FlowLogs(ctx, config)
	case "vpc_peering_connections":
		return c.vpcPeeringConnections(ctx, config)
	default:
		return fmt.Errorf("unsupported resource ec2.%s", resource)
	}
}
