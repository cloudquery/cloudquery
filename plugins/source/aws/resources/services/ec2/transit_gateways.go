package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Ec2TransitGateways() *schema.Table {
	return &schema.Table{
		Name:          "aws_ec2_transit_gateways",
		Resolver:      fetchEc2TransitGateways,
		Multiplex:     client.ServiceAccountRegionMultiplexer("ec2"),
		IgnoreInTests: true,
		Columns: []schema.Column{
			{
				Name:            "account_id",
				Type:            schema.TypeString,
				Resolver:        client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "options",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Options"),
			},
			{
				Name: "creation_time",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "description",
				Type: schema.TypeString,
			},
			{
				Name: "owner_id",
				Type: schema.TypeString,
			},
			{
				Name: "state",
				Type: schema.TypeString,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the resource.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TransitGatewayArn"),
			},
			{
				Name:            "id",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("TransitGatewayId"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
		},
		Relations: []*schema.Table{
			{
				Name:          "aws_ec2_transit_gateway_attachments",
				Resolver:      fetchEc2TransitGatewayAttachments,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:     "transit_gateway_cq_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "association",
						Type:     schema.TypeJSON,
						Resolver: schema.PathResolver("Association"),
					},
					{
						Name: "creation_time",
						Type: schema.TypeTimestamp,
					},
					{
						Name: "resource_id",
						Type: schema.TypeString,
					},
					{
						Name: "resource_owner_id",
						Type: schema.TypeString,
					},
					{
						Name: "resource_type",
						Type: schema.TypeString,
					},
					{
						Name: "state",
						Type: schema.TypeString,
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: client.ResolveTags,
					},
					{
						Name: "transit_gateway_owner_id",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:          "aws_ec2_transit_gateway_route_tables",
				Resolver:      fetchEc2TransitGatewayRouteTables,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:     "transit_gateway_cq_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "creation_time",
						Type: schema.TypeTimestamp,
					},
					{
						Name: "default_association_route_table",
						Type: schema.TypeBool,
					},
					{
						Name: "default_propagation_route_table",
						Type: schema.TypeBool,
					},
					{
						Name: "state",
						Type: schema.TypeString,
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: client.ResolveTags,
					},
					{
						Name: "transit_gateway_route_table_id",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:          "aws_ec2_transit_gateway_vpc_attachments",
				Resolver:      fetchEc2TransitGatewayVpcAttachments,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:     "transit_gateway_cq_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "creation_time",
						Type: schema.TypeTimestamp,
					},
					{
						Name:     "options",
						Type:     schema.TypeJSON,
						Resolver: schema.PathResolver("Options"),
					},
					{
						Name: "state",
						Type: schema.TypeString,
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: client.ResolveTags,
					},
					{
						Name: "transit_gateway_attachment_id",
						Type: schema.TypeString,
					},
					{
						Name: "vpc_id",
						Type: schema.TypeString,
					},
					{
						Name: "vpc_owner_id",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:          "aws_ec2_transit_gateway_peering_attachments",
				Resolver:      fetchEc2TransitGatewayPeeringAttachments,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:     "transit_gateway_cq_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "accepter_tgw_info",
						Type:     schema.TypeJSON,
						Resolver: schema.PathResolver("AccepterTgwInfo"),
					},
					{
						Name: "creation_time",
						Type: schema.TypeTimestamp,
					},
					{
						Name:     "requester_tgw_info",
						Type:     schema.TypeJSON,
						Resolver: schema.PathResolver("RequesterTgwInfo"),
					},
					{
						Name: "state",
						Type: schema.TypeString,
					},
					{
						Name:     "status",
						Type:     schema.TypeJSON,
						Resolver: schema.PathResolver("Status"),
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: client.ResolveTags,
					},
					{
						Name: "transit_gateway_attachment_id",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:          "aws_ec2_transit_gateway_multicast_domains",
				Resolver:      fetchEc2TransitGatewayMulticastDomains,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:     "transit_gateway_cq_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "creation_time",
						Type: schema.TypeTimestamp,
					},
					{
						Name:     "options",
						Type:     schema.TypeJSON,
						Resolver: schema.PathResolver("Options"),
					},
					{
						Name: "owner_id",
						Type: schema.TypeString,
					},
					{
						Name: "state",
						Type: schema.TypeString,
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: client.ResolveTags,
					},
					{
						Name: "transit_gateway_multicast_domain_arn",
						Type: schema.TypeString,
					},
					{
						Name: "transit_gateway_multicast_domain_id",
						Type: schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//
//	Table Resolver Functions
//
// ====================================================================================================================
func fetchEc2TransitGateways(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config ec2.DescribeTransitGatewaysInput
	c := meta.(*client.Client)
	svc := c.Services().EC2
	for {
		output, err := svc.DescribeTransitGateways(ctx, &config)
		if err != nil {
			return err
		}
		res <- output.TransitGateways
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}

func fetchEc2TransitGatewayAttachments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.TransitGateway)

	config := ec2.DescribeTransitGatewayAttachmentsInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("transit-gateway-id"),
				Values: []string{*r.TransitGatewayId},
			},
		},
	}
	c := meta.(*client.Client)
	svc := c.Services().EC2
	for {
		output, err := svc.DescribeTransitGatewayAttachments(ctx, &config)
		if err != nil {
			return err
		}
		res <- output.TransitGatewayAttachments
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}

func fetchEc2TransitGatewayRouteTables(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.TransitGateway)

	config := ec2.DescribeTransitGatewayRouteTablesInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("transit-gateway-id"),
				Values: []string{*r.TransitGatewayId},
			},
		},
	}
	c := meta.(*client.Client)
	svc := c.Services().EC2
	for {
		output, err := svc.DescribeTransitGatewayRouteTables(ctx, &config)
		if err != nil {
			return err
		}
		res <- output.TransitGatewayRouteTables
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}

func fetchEc2TransitGatewayVpcAttachments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.TransitGateway)

	config := ec2.DescribeTransitGatewayVpcAttachmentsInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("transit-gateway-id"),
				Values: []string{*r.TransitGatewayId},
			},
		},
	}
	c := meta.(*client.Client)
	svc := c.Services().EC2
	for {
		output, err := svc.DescribeTransitGatewayVpcAttachments(ctx, &config)
		if err != nil {
			return err
		}
		res <- output.TransitGatewayVpcAttachments
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}

func fetchEc2TransitGatewayPeeringAttachments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.TransitGateway)

	config := ec2.DescribeTransitGatewayPeeringAttachmentsInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("transit-gateway-id"),
				Values: []string{*r.TransitGatewayId},
			},
		},
	}

	c := meta.(*client.Client)
	svc := c.Services().EC2
	for {
		output, err := svc.DescribeTransitGatewayPeeringAttachments(ctx, &config)
		if err != nil {
			return err
		}
		res <- output.TransitGatewayPeeringAttachments
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}

func fetchEc2TransitGatewayMulticastDomains(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.TransitGateway)

	config := ec2.DescribeTransitGatewayMulticastDomainsInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("transit-gateway-id"),
				Values: []string{*r.TransitGatewayId},
			},
		},
	}

	c := meta.(*client.Client)
	svc := c.Services().EC2
	for {
		output, err := svc.DescribeTransitGatewayMulticastDomains(ctx, &config)
		if err != nil {
			return err
		}
		res <- output.TransitGatewayMulticastDomains
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
