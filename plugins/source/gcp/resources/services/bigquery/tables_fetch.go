package bigquery

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"google.golang.org/api/bigquery/v2"
)

type tablesPreWrapper struct {
	datasetID string
	tableID   string
	svc       *bigquery.Service
}

func fetchTables(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	p := parent.Item.(*datasetWrapper)
	dsid := p.DatasetReference.DatasetId
	for {
		output, err := p.svc.Tables.List(c.ProjectId, dsid).PageToken(nextPageToken).Context(ctx).Do()
		if err != nil {
			return err
		}
		for i := range output.Tables {
			res <- &tablesPreWrapper{
				datasetID: dsid,
				tableID:   output.Tables[i].TableReference.TableId,
				svc:       p.svc,
			}
		}
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}

func tableGet(ctx context.Context, meta schema.ClientMeta, r *schema.Resource) error {
	c := meta.(*client.Client)
	wrapped := r.Item.(*tablesPreWrapper)
	item, err := wrapped.svc.Tables.Get(c.ProjectId, wrapped.datasetID, wrapped.tableID).Context(ctx).Do()
	if err != nil {
		return err
	}
	r.SetItem(item)
	return nil
}
