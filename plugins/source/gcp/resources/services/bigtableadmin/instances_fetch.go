package bigtableadmin

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugins/source/gcp/client"

	"cloud.google.com/go/bigtable"
)

func fetchInstances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	gcpClient, err := bigtable.NewInstanceAdminClient(ctx, c.ProjectId, c.ClientOptions...)
	if err != nil {
		return err
	}
	resp, err := gcpClient.Instances(ctx)
	if err != nil {
		return err
	}

	res <- resp
	return nil
}
