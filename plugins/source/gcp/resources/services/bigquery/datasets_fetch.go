package bigquery

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"google.golang.org/api/bigquery/v2"
)

type datasetPreWrapper struct {
	datasetID string
	svc       *bigquery.Service
}

func fetchDatasets(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	bigqueryService, err := bigquery.NewService(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	for {
		output, err := bigqueryService.Datasets.List(c.ProjectId).PageToken(nextPageToken).Do()
		if err != nil {
			return err
		}
		for i := range output.Datasets {
			res <- &datasetPreWrapper{
				datasetID: output.Datasets[i].DatasetReference.DatasetId,
				svc:       bigqueryService,
			}
		}

		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}

func datasetGet(ctx context.Context, meta schema.ClientMeta, r *schema.Resource) error {
	c := meta.(*client.Client)
	wrapped := r.Item.(*datasetPreWrapper)
	item, err := wrapped.svc.Datasets.Get(c.ProjectId, wrapped.datasetID).Context(ctx).Do()
	if err != nil {
		return err
	}
	r.SetItem(&datasetWrapper{
		Dataset: item,
		svc:     wrapped.svc,
	})
	return nil
}
