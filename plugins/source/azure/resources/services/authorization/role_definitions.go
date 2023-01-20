package authorization

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func RoleDefinitions() *schema.Table {
	return &schema.Table{
		Name:        "azure_authorization_role_definitions",
		Resolver:    fetchRoleDefinitions,
		Description: "https://learn.microsoft.com/en-us/rest/api/authorization/role-definitions/list?tabs=HTTP#roledefinition",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_authorization_role_definitions", client.Namespacemicrosoft_authorization),
		Transform:   transformers.TransformWithStruct(&armauthorization.RoleDefinition{}),
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
