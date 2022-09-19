// Code generated by codegen; DO NOT EDIT.

package ec2

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func TransitGatewayAttachments() *schema.Table {
	return &schema.Table{
		Name:      "aws_ec2_transit_gateway_attachments",
		Resolver:  fetchEc2TransitGatewayAttachments,
		Multiplex: client.ServiceAccountRegionMultiplexer("ec2"),
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
				Name:     "transit_gateway_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentResourceFieldResolver("arn"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
			{
				Name:     "association",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Association"),
			},
			{
				Name:     "creation_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreationTime"),
			},
			{
				Name:     "resource_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceId"),
			},
			{
				Name:     "resource_owner_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceOwnerId"),
			},
			{
				Name:     "resource_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceType"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "transit_gateway_attachment_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TransitGatewayAttachmentId"),
			},
			{
				Name:     "transit_gateway_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TransitGatewayId"),
			},
			{
				Name:     "transit_gateway_owner_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TransitGatewayOwnerId"),
			},
		},
	}
}
