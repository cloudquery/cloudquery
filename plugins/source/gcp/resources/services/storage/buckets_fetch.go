package storage

import (
	"context"

	"cloud.google.com/go/storage"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/api/iterator"
)

func fetchBuckets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	storageClient, err := storage.NewClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := storageClient.Buckets(ctx, c.ProjectId)
	for {
		bucket, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		res <- bucket
	}
	return nil
}
