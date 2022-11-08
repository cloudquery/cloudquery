package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func RAMResources() []*Resource {
	mx := `client.ServiceAccountRegionMultiplexer("ram")`
	pk := schema.ColumnCreationOptions{PrimaryKey: true}
	resources := []*Resource{
		{
			SubService: "principals",
			Struct:     new(types.Principal),
			Multiplex:  mx,
			SkipFields: []string{"Id"},
			ExtraColumns: append(defaultRegionalColumns,
				codegen.ColumnDefinition{
					Name:     "id",
					Type:     schema.TypeString,
					Resolver: `schema.PathResolver("Id")`,
					Options:  pk,
				},
			),
			ResolverAndMockTestTemplate: "list_resources_paginated_1",
		},
		{
			SubService: "resources",
			Struct:     new(types.Resource),
			Multiplex:  mx,
			SkipFields: []string{"Arn"},
			ExtraColumns: append(defaultRegionalColumns,
				codegen.ColumnDefinition{
					Name:     "arn",
					Type:     schema.TypeString,
					Resolver: `schema.PathResolver("Arn")`,
					Options:  pk,
				},
			),
			ResolverAndMockTestTemplate: "list_resources_paginated_1",
		},
		{
			SubService: "resource_shares",
			Struct:     new(types.ResourceShare),
			Multiplex:  mx,
			SkipFields: []string{"ResourceShareArn"},
			ExtraColumns: append(defaultRegionalColumns,
				codegen.ColumnDefinition{
					Name:     "arn",
					Type:     schema.TypeString,
					Resolver: `schema.PathResolver("ResourceShareArn")`,
					Options:  pk,
				},
			),
			Relations: []string{
				"ResourceShareAssociatedPrincipals()",
				"ResourceShareAssociatedResources()",
			},
		},
		{
			SubService: "resource_share_associated_principals",
			Struct:     new(types.ResourceShareAssociation),
			Multiplex:  "", // it's a relation for resource_shares
			SkipFields: []string{"AssociatedEntity", "ResourceShareArn"},
			ExtraColumns: append(defaultRegionalColumns,
				codegen.ColumnDefinition{
					Name:     "associated_entity",
					Type:     schema.TypeString,
					Resolver: `schema.PathResolver("AssociatedEntity")`,
					Options:  pk,
				},
				codegen.ColumnDefinition{
					Name:        "resource_share_arn",
					Type:        schema.TypeString,
					Resolver:    `schema.PathResolver("ResourceShareArn")`,
					Description: "Resource Share ARN",
					Options:     pk,
				},
			),
		},
		{
			SubService: "resource_share_associated_resources",
			Struct:     new(types.ResourceShareAssociation),
			Multiplex:  "", // it's a relation for resource_shares
			SkipFields: []string{"AssociatedEntity", "ResourceShareArn"},
			ExtraColumns: append(defaultRegionalColumns,
				codegen.ColumnDefinition{
					Name:     "associated_entity",
					Type:     schema.TypeString,
					Resolver: `schema.PathResolver("AssociatedEntity")`,
					Options:  pk,
				},
				codegen.ColumnDefinition{
					Name:        "resource_share_arn",
					Type:        schema.TypeString,
					Resolver:    `schema.PathResolver("ResourceShareArn")`,
					Description: "Resource Share ARN",
					Options:     pk,
				},
			),
		},
		{
			SubService: "resource_share_invitations",
			Struct:     new(types.ResourceShareInvitation),
			Multiplex:  mx,
			SkipFields: []string{"ResourceShareInvitationArn"},
			ExtraColumns: append(defaultRegionalColumns,
				codegen.ColumnDefinition{
					Name:     "arn",
					Type:     schema.TypeString,
					Resolver: `schema.PathResolver("ResourceShareInvitationArn")`,
					Options:  pk,
				},
			),
			ResolverAndMockTestTemplate: "get_resources_paginated_1",
		},
		{
			SubService: "resource_share_permissions",
			Struct:     new(types.ResourceSharePermissionSummary),
			Multiplex:  mx,
			SkipFields: []string{"Arn"},
			ExtraColumns: append(defaultRegionalColumns,
				codegen.ColumnDefinition{
					Name:     "arn",
					Type:     schema.TypeString,
					Resolver: `schema.PathResolver("Arn")`,
					Options:  pk,
				},
				codegen.ColumnDefinition{
					// grabbed from types.ResourceSharePermissionDetail
					Name:     "permission",
					Type:     schema.TypeJSON,
					Resolver: `resolveResourceSharePermissionDetailPermission`,
				},
			),
		},
		{
			SubService: "resource_types",
			Struct:     new(types.ServiceNameAndResourceType),
			Multiplex:  mx,
			SkipFields: []string{"ResourceType", "ServiceName"},
			ExtraColumns: append(defaultRegionalColumns,
				codegen.ColumnDefinition{
					Name:     "resource_type",
					Type:     schema.TypeString,
					Resolver: `schema.PathResolver("ResourceType")`,
					Options:  pk,
				},
				codegen.ColumnDefinition{
					Name:     "service_name",
					Type:     schema.TypeString,
					Resolver: `schema.PathResolver("ServiceName")`,
					Options:  pk,
				},
			),
		},
	}
	for _, r := range resources {
		r.Service = "ram"
		r.Description = "https://docs.aws.amazon.com/ram/latest/APIReference/API_" + r.StructName() + ".html"
		if len(r.ResolverAndMockTestTemplate) > 0 {
			r.ShouldGenerateResolverAndMockTest = true
			r.MaxResults = 500
		}
	}
	return resources
}
