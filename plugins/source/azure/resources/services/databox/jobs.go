package databox

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Jobs() *schema.Table {
	return &schema.Table{
		Name:        "azure_databox_jobs",
		Resolver:    fetchJobs,
		Description: "https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox@v1.0.0#JobResource",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_databox_jobs", client.Namespacemicrosoft_databox),
		Transform:   transformers.TransformWithStruct(&armdatabox.JobResource{}),
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

func fetchJobs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armdatabox.NewJobsClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
