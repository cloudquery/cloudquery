package bigquery

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
	"google.golang.org/api/bigquery/v2"
)

func fetchDatasets(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		output, err := c.Services.BigqueryService.Datasets.List(c.ProjectId).PageToken(nextPageToken).Do()
		if err != nil {
			return errors.WithStack(err)
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
	item, err := c.Services.BigqueryService.Datasets.Get(c.ProjectId, datasetListDataset.DatasetReference.DatasetId).Do()
	if err != nil {
		return errors.WithStack(err)
	}
	r.SetItem(item)
	return nil
}
