package storage

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
)

func fetchBucketPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	bkt := parent.Get("name").(*schema.Text).Str
	output, err := c.Services.StorageClient.Bucket(bkt).IAM().V3().Policy(ctx)
	if err != nil {
		return errors.WithStack(err)
	}
	res <- output
	return nil
}
