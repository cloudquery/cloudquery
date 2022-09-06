package storage

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
	"google.golang.org/api/iterator"
)

func fetchBuckets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	it := c.Services.StorageClient.Buckets(ctx, c.ProjectId)
	for {
		bucket, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return errors.WithStack(err)
		}
		res <- bucket
	}
	return nil
}
