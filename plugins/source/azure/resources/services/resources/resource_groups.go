package resources

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func ResourceGroups() *schema.Table {
	return &schema.Table{
		Name:                 "azure_resources_resource_groups",
		Resolver:             fetchResourceGroups,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/resources/resource-groups/list#resourcegroup",
		Multiplex:            client.SubscriptionMultiplex,
		Transform:            transformers.TransformWithStruct(&armresources.ResourceGroup{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchResourceGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	// We already fetched the resource groups for this subscription, no need to fetch again
	res <- cl.ResourceGroups[cl.SubscriptionId]
	return nil
}
