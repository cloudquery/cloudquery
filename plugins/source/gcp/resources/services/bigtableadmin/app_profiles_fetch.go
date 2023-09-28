package bigtableadmin

import (
	"context"

	"google.golang.org/api/iterator"

	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"

	"cloud.google.com/go/bigtable"
)

func fetchAppProfiles(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	instance := parent.Item.(*bigtable.InstanceInfo)
	gcpClient, err := bigtable.NewInstanceAdminClient(ctx, c.ProjectId, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.ListAppProfiles(ctx, instance.Name)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		res <- resp
	}
	return nil
}
