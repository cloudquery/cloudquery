package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func DirectConnectResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "connections",
			Struct:     &types.Connection{},
			SkipFields: []string{"ConnectionId", "Tags", "Region"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveConnectionARN()`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "id",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("ConnectionId")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `client.ResolveTags`,
					},
				}...),
		},
		{
			SubService: "gateways",
			Struct:     &types.DirectConnectGateway{},
			SkipFields: []string{"DirectConnectGatewayId"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveGatewayARN()`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "id",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("DirectConnectGatewayId")`,
					},
				}...),
			Relations: []string{
				"GatewayAssociations()",
				"GatewayAttachments()",
			},
		},
		{
			SubService: "gateway_associations",
			Struct:     &types.DirectConnectGatewayAssociation{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "gateway_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("arn")`,
					},
					{
						Name:     "gateway_id",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("id")`,
					},
				}...),
		},
		{
			SubService: "gateway_attachments",
			Struct:     &types.DirectConnectGatewayAttachment{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "gateway_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("arn")`,
					},
					{
						Name:     "gateway_id",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("id")`,
					},
				}...),
		},
		{
			SubService: "lags",
			Struct:     &types.Lag{},
			SkipFields: []string{"LagId", "Tags", "Region"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveLagARN()`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "id",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("LagId")`,
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `client.ResolveTags`,
					},
				}...),
		},
		{
			SubService: "virtual_gateways",
			Struct:     &types.VirtualGateway{},
			SkipFields: []string{"VirtualGatewayId"},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "account_id",
					Type:     schema.TypeString,
					Resolver: `client.ResolveAWSAccount`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:     "region",
					Type:     schema.TypeString,
					Resolver: "client.ResolveAWSRegion",
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:     "id",
					Type:     schema.TypeString,
					Resolver: `schema.PathResolver("VirtualGatewayId")`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
			},
		},
		{
			SubService: "virtual_interfaces",
			Struct:     &types.VirtualInterface{},
			SkipFields: []string{"Tags", "VirtualInterfaceId", "Region"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveVirtualInterfaceARN()`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `client.ResolveTags`,
					},
					{
						Name:     "id",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("VirtualInterfaceId")`,
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "directconnect"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("directconnect")`
	}
	return resources
}
