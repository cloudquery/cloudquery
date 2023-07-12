package redhatopenshift

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redhatopenshift/armredhatopenshift"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func OpenShiftClusters() *schema.Table {
	return &schema.Table{
		Name:                 "azure_redhatopenshift_open_shift_clusters",
		Resolver:             fetchOpenShiftClusters,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/openshift/open-shift-clusters/list?tabs=HTTP#openshiftcluster",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_redhatopenshift_open_shift_clusters", client.Namespacemicrosoft_redhatopenshift),
		Transform:            transformers.TransformWithStruct(&armredhatopenshift.OpenShiftCluster{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchOpenShiftClusters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armredhatopenshift.NewOpenShiftClustersClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
