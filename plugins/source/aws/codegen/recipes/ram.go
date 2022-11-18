package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/ram"
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func RAMResources() []*Resource {
	mx := `client.ServiceAccountRegionMultiplexer("ram")`
	resources := []*Resource{
		{
			SubService:   "principals",
			Struct:       new(types.Principal),
			Multiplex:    mx,
			PKColumns:    []string{"id", "account_id"},
			ExtraColumns: defaultRegionalColumns,
		},
		{
			SubService:   "resources",
			Struct:       new(types.Resource),
			Multiplex:    mx,
			PKColumns:    []string{"arn"},
			ExtraColumns: defaultRegionalColumns,
		},
		{
			SubService:   "resource_shares",
			Struct:       new(types.ResourceShare),
			Multiplex:    mx,
			PKColumns:    []string{"arn"},
			ExtraColumns: defaultRegionalColumns,
			Relations: []string{
				"ResourceSharePermissions()",
			},
			NameTransformer: CreateReplaceTransformer(map[string]string{"resource_share_arn": "arn"}),
		},
		{
			SubService:   "resource_share_associations",
			Struct:       new(types.ResourceShareAssociation),
			Multiplex:    mx,
			PKColumns:    []string{"associated_entity", "resource_share_arn"},
			ExtraColumns: defaultRegionalColumns,
		},
		{
			SubService:                  "resource_share_invitations",
			Struct:                      new(types.ResourceShareInvitation),
			Multiplex:                   mx,
			PKColumns:                   []string{"arn"},
			ExtraColumns:                defaultRegionalColumns,
			ResolverAndMockTestTemplate: "describe_resources_1",
			NameTransformer:             CreateReplaceTransformer(map[string]string{"resource_share_invitation_arn": "arn"}),
			CustomDescribeInput:         `getResourceShareInvitationsInput()`,
		},
		{
			SubService: "resource_share_permissions",
			Struct:     new(types.ResourceSharePermissionSummary),
			Multiplex:  "", // it's a relation for resource_shares
			PKColumns:  []string{"arn"},
			ExtraColumns: append(defaultRegionalColumns,
				codegen.ColumnDefinition{
					// grabbed from types.ResourceSharePermissionDetail
					Name:     "permission",
					Type:     schema.TypeJSON,
					Resolver: `resolveResourceSharePermissionDetailPermission`,
				},
			),
			CustomListInput: `listResourceSharePermissionsInput()`,
		},
		{
			SubService:      "resource_types",
			Struct:          new(types.ServiceNameAndResourceType),
			Multiplex:       mx,
			PKColumns:       []string{"account_id", "resource_type", "service_name"},
			ExtraColumns:    defaultRegionalColumns,
			CustomListInput: `listResourceTypesInput()`,
		},
	}
	for _, r := range resources {
		r.Service = "ram"
		r.Description = "https://docs.aws.amazon.com/ram/latest/APIReference/API_" + r.StructName() + ".html"
		if len(r.ResolverAndMockTestTemplate) > 0 {
			r.ShouldGenerateResolverAndMockTest = true
			r.Client = &ram.Client{}
		}
	}
	return resources
}
