package bigquery

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/api/bigquery/v2"
)

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
		res <- output.Datasets

		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}

func datasetGet(ctx context.Context, meta schema.ClientMeta, r *schema.Resource) error {
	c := meta.(*client.Client)
	datasetListDataset := r.Item.(*bigquery.DatasetListDatasets)
	bigqueryService, err := bigquery.NewService(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	item, err := bigqueryService.Datasets.Get(c.ProjectId, datasetListDataset.DatasetReference.DatasetId).Do()
	if err != nil {
		return err
	}
	r.SetItem(item)
	return nil
}
