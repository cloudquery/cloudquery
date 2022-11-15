package storage

import (
	"cloud.google.com/go/storage"
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

func resolveBucketSelfLink(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(*storage.BucketAttrs)
	return resource.Set(c.Name, "https://www.googleapis.com/storage/v1/b/"+p.Name)
}
