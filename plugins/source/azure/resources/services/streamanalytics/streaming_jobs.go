package streamanalytics

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func StreamingJobs() *schema.Table {
	return &schema.Table{
		Name:                 "azure_streamanalytics_streaming_jobs",
		Resolver:             fetchStreamingJobs,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/streamanalytics/2020-03-01/streaming-jobs/list?tabs=HTTP#streamingjob",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_streamanalytics_streaming_jobs", client.Namespacemicrosoft_streamanalytics),
		Transform:            transformers.TransformWithStruct(&armstreamanalytics.StreamingJob{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchStreamingJobs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armstreamanalytics.NewStreamingJobsClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
