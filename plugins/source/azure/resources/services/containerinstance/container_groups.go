package containerinstance

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerinstance/armcontainerinstance/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func ContainerGroups() *schema.Table {
	return &schema.Table{
		Name:                 "azure_containerinstance_container_groups",
		Resolver:             fetchContainerGroups,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/container-instances/container-groups/list?tabs=HTTP#containergroup",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_containerinstance_container_groups", client.Namespacemicrosoft_containerinstance),
		Transform:            transformers.TransformWithStruct(&armcontainerinstance.ContainerGroup{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchContainerGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armcontainerinstance.NewContainerGroupsClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
