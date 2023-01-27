package connectedvmware

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Clusters() *schema.Table {
	return &schema.Table{
		Name:        "azure_connectedvmware_clusters",
		Resolver:    fetchClusters,
		Description: "https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware@v0.1.0#Cluster",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_connectedvmware_clusters", client.Namespacemicrosoft_connectedvmwarevsphere),
		Transform:   transformers.TransformWithStruct(&armconnectedvmware.Cluster{}),
		Columns:     schema.ColumnList{client.SubscriptionID, client.IDColumn},
	}
}

func fetchClusters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armconnectedvmware.NewClustersClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
