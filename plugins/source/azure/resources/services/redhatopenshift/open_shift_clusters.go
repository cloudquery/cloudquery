package redhatopenshift

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redhatopenshift/armredhatopenshift"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func OpenShiftClusters() *schema.Table {
	return &schema.Table{
		Name:        "azure_redhatopenshift_open_shift_clusters",
		Resolver:    fetchOpenShiftClusters,
		Description: "https://learn.microsoft.com/en-us/rest/api/openshift/open-shift-clusters/list?tabs=HTTP#openshiftcluster",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_redhatopenshift_open_shift_clusters", client.Namespacemicrosoft_redhatopenshift),
		Transform:   transformers.TransformWithStruct(&armredhatopenshift.OpenShiftCluster{}),
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
