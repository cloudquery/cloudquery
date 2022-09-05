package bigquery

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
	"google.golang.org/api/bigquery/v2"
)

func fetchTables(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		output, err := c.Services.BigqueryService.Tables.List(c.ProjectId, parent.Item.(*bigquery.Dataset).DatasetReference.DatasetId).PageToken(nextPageToken).Do()
		if err != nil {
			return errors.WithStack(err)
		}
		res <- output.Tables

		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}

func tableGet(ctx context.Context, meta schema.ClientMeta, r *schema.Resource) error {
	c := meta.(*client.Client)
	item, err := c.Services.BigqueryService.Tables.Get(c.ProjectId, r.Parent.Item.(*bigquery.Dataset).DatasetReference.DatasetId, r.Item.(*bigquery.TableListTables).TableReference.TableId).Do()
	if err != nil {
		return errors.WithStack(err)
	}
	r.SetItem(item)
	return nil
}
