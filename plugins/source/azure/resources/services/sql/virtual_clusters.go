package sql

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func VirtualClusters() *schema.Table {
	return &schema.Table{
		Name:                 "azure_sql_virtual_clusters",
		Resolver:             fetchVirtualClusters,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/sql/2021-11-01/virtual-clusters/list?tabs=HTTP#virtualcluster",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_sql_virtual_clusters", client.Namespacemicrosoft_sql),
		Transform:            transformers.TransformWithStruct(&armsql.VirtualCluster{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchVirtualClusters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armsql.NewVirtualClustersClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
