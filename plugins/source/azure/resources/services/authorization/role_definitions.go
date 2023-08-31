package authorization

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func RoleDefinitions() *schema.Table {
	return &schema.Table{
		Name:                 "azure_authorization_role_definitions",
		Resolver:             fetchRoleDefinitions,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/authorization/role-definitions/list?tabs=HTTP#roledefinition",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_authorization_role_definitions", client.Namespacemicrosoft_authorization),
		Transform:            transformers.TransformWithStruct(&armauthorization.RoleDefinition{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}
