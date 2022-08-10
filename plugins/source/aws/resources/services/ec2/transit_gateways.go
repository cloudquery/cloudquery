package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Ec2TransitGateways() *schema.Table {
	return &schema.Table{
		Name:          "aws_ec2_transit_gateways",
		Resolver:      fetchEc2TransitGateways,
		Multiplex:     client.ServiceAccountRegionMultiplexer("ec2"),
		IgnoreError:   client.IgnoreCommonErrors,
		DeleteFilter:  client.DeleteAccountRegionFilter,
		Options:       schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "id"}},
		IgnoreInTests: true,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "amazon_side_asn",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("Options.AmazonSideAsn"),
			},
			{
				Name:     "association_default_route_table_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Options.AssociationDefaultRouteTableId"),
			},
			{
				Name:     "auto_accept_shared_attachments",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Options.AutoAcceptSharedAttachments"),
			},
			{
				Name: "creation_time",
				Type: schema.TypeTimestamp,
			},
			{
				Name:     "default_route_table_association",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Options.DefaultRouteTableAssociation"),
			},
			{
				Name:     "default_route_table_propagation",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Options.DefaultRouteTablePropagation"),
			},
			{
				Name: "description",
				Type: schema.TypeString,
			},
			{
				Name:     "dns_support",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Options.DnsSupport"),
			},
			{
				Name:     "multicast_support",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Options.MulticastSupport"),
			},
			{
				Name: "owner_id",
				Type: schema.TypeString,
			},
			{
				Name:     "propagation_default_route_table_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Options.PropagationDefaultRouteTableId"),
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
				Name:     "transit_gateway_cidr_blocks",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Options.TransitGatewayCidrBlocks"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TransitGatewayId"),
			},
			{
				Name:     "vpn_ecmp_support",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Options.VpnEcmpSupport"),
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
						Name:     "association_state",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Association.State"),
					},
					{
						Name:     "association_route_table_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Association.TransitGatewayRouteTableId"),
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
						Name:     "appliance_mode_support",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Options.ApplianceModeSupport"),
					},
					{
						Name:     "dns_support",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Options.DnsSupport"),
					},
					{
						Name:     "ipv6_support",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Options.Ipv6Support"),
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
						Name:     "accepter_owner_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("AccepterTgwInfo.OwnerId"),
					},
					{
						Name:     "accepter_region",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("AccepterTgwInfo.Region"),
					},
					{
						Name:     "accepter_transit_gateway_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("AccepterTgwInfo.TransitGatewayId"),
					},
					{
						Name: "creation_time",
						Type: schema.TypeTimestamp,
					},
					{
						Name:     "requester_owner_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("RequesterTgwInfo.OwnerId"),
					},
					{
						Name:     "requester_region",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("RequesterTgwInfo.Region"),
					},
					{
						Name:     "requester_transit_gateway_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("RequesterTgwInfo.TransitGatewayId"),
					},
					{
						Name: "state",
						Type: schema.TypeString,
					},
					{
						Name:     "status_code",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Status.Code"),
					},
					{
						Name:     "status_message",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Status.Message"),
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
						Name:     "auto_accept_shared_associations",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Options.AutoAcceptSharedAssociations"),
					},
					{
						Name:     "igmpv2_support",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Options.Igmpv2Support"),
					},
					{
						Name:     "static_sources_support",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Options.StaticSourcesSupport"),
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
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchEc2TransitGateways(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config ec2.DescribeTransitGatewaysInput
	c := meta.(*client.Client)
	svc := c.Services().EC2
	for {
		output, err := svc.DescribeTransitGateways(ctx, &config, func(options *ec2.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
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
		output, err := svc.DescribeTransitGatewayAttachments(ctx, &config, func(options *ec2.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
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
		output, err := svc.DescribeTransitGatewayRouteTables(ctx, &config, func(options *ec2.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
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
		output, err := svc.DescribeTransitGatewayVpcAttachments(ctx, &config, func(options *ec2.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
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
		output, err := svc.DescribeTransitGatewayPeeringAttachments(ctx, &config, func(options *ec2.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
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
		output, err := svc.DescribeTransitGatewayMulticastDomains(ctx, &config, func(options *ec2.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- output.TransitGatewayMulticastDomains
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
