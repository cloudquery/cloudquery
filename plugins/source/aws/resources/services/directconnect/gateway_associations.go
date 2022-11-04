// Code generated by codegen; DO NOT EDIT.

package directconnect

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func GatewayAssociations() *schema.Table {
	return &schema.Table{
		Name:        "aws_directconnect_gateway_associations",
		Description: `https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DirectConnectGatewayAssociation.html`,
		Resolver:    fetchDirectconnectGatewayAssociations,
		Multiplex:   client.ServiceAccountRegionMultiplexer("directconnect"),
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
				Name:     "gateway_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "gateway_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
			},
			{
				Name:     "allowed_prefixes_to_direct_connect_gateway",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AllowedPrefixesToDirectConnectGateway"),
			},
			{
				Name:     "associated_gateway",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AssociatedGateway"),
			},
			{
				Name:     "association_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AssociationId"),
			},
			{
				Name:     "association_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AssociationState"),
			},
			{
				Name:     "direct_connect_gateway_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DirectConnectGatewayId"),
			},
			{
				Name:     "direct_connect_gateway_owner_account",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DirectConnectGatewayOwnerAccount"),
			},
			{
				Name:     "state_change_error",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StateChangeError"),
			},
			{
				Name:     "virtual_gateway_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VirtualGatewayId"),
			},
			{
				Name:     "virtual_gateway_owner_account",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VirtualGatewayOwnerAccount"),
			},
			{
				Name:     "virtual_gateway_region",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VirtualGatewayRegion"),
			},
		},
	}
}
