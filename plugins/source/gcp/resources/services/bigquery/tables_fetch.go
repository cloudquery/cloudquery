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
	bigqueryService, err := bigquery.NewService(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	for {
		output, err := bigqueryService.Tables.List(c.ProjectId, parent.Item.(*bigquery.Dataset).DatasetReference.DatasetId).PageToken(nextPageToken).Do()
		if err != nil {
			return err
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
	bigqueryService, err := bigquery.NewService(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	item, err := bigqueryService.Tables.Get(c.ProjectId, r.Parent.Item.(*bigquery.Dataset).DatasetReference.DatasetId, r.Item.(*bigquery.TableListTables).TableReference.TableId).Do()
	if err != nil {
		return errors.WithStack(err)
	}
	r.SetItem(item)
	return nil
}
