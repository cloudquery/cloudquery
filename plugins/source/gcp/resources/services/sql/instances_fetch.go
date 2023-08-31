package sql

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	sql "google.golang.org/api/sqladmin/v1beta4"
)

func fetchInstances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	sqlClient, err := sql.NewService(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	for {
		output, err := sqlClient.Instances.List(c.ProjectId).MaxResults(1000).PageToken(nextPageToken).Do()
		if err != nil {
			return err
		}
		res <- output.Items
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
