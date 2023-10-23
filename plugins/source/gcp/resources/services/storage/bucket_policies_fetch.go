package storage

import (
	"context"

	"cloud.google.com/go/storage"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/scalar"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func fetchBucketPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	bkt := parent.Get("name").(*scalar.String).Value
	storageClient, err := storage.NewClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	output, err := storageClient.Bucket(bkt).IAM().V3().Policy(ctx)
	if err != nil {
		return err
	}
	res <- output
	return nil
}
