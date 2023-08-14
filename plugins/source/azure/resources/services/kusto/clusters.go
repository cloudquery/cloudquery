package kusto

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Clusters() *schema.Table {
	return &schema.Table{
		Name:                 "azure_kusto_clusters",
		Resolver:             fetchClusters,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/azurerekusto/clusters/list?tabs=HTTP#cluster",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_kusto_clusters", client.Namespacemicrosoft_kusto),
		Transform:            transformers.TransformWithStruct(&armkusto.Cluster{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchClusters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armkusto.NewClustersClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
