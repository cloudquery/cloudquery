package databox

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Jobs() *schema.Table {
	return &schema.Table{
		Name:                 "azure_databox_jobs",
		Resolver:             fetchJobs,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox@v1.0.0#JobResource",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_databox_jobs", client.Namespacemicrosoft_databox),
		Transform:            transformers.TransformWithStruct(&armdatabox.JobResource{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
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
