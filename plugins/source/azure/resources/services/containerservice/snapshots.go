package containerservice

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Snapshots() *schema.Table {
	return &schema.Table{
		Name:                 "azure_containerservice_snapshots",
		Resolver:             fetchSnapshots,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/aks/snapshots/list?tabs=HTTP#snapshot",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_containerservice_snapshots", client.Namespacemicrosoft_containerservice),
		Transform:            transformers.TransformWithStruct(&armcontainerservice.Snapshot{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchSnapshots(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armcontainerservice.NewSnapshotsClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
