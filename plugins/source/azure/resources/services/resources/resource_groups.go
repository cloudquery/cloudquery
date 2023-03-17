package resources

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ResourceGroups() *schema.Table {
	return &schema.Table{
		Name:        "azure_resources_resource_groups",
		Resolver:    fetchResourceGroups,
		Description: "https://learn.microsoft.com/en-us/rest/api/resources/resource-groups/list#resourcegroup",
		Multiplex:   client.SubscriptionMultiplex,
		Transform:   transformers.TransformWithStruct(&armresources.ResourceGroup{}, transformers.WithPrimaryKeys("ID")),
		Columns:     schema.ColumnList{},
	}
}

func fetchResourceGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armresources.NewResourceGroupsClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListPager(nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
