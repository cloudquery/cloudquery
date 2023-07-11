package bigtableadmin

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/api/iterator"

	"cloud.google.com/go/bigtable"
)

func fetchBackups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	instance := parent.Parent.Item.(*bigtable.InstanceInfo)
	cluster := parent.Item.(*bigtable.ClusterInfo)
	gcpClient, err := bigtable.NewAdminClient(ctx, c.ProjectId, instance.Name, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.Backups(ctx, cluster.Name)
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
