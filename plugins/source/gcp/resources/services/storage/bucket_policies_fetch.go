package storage

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
)

func fetchBucketPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	output, err := c.Services.StorageClient.Bucket(parent.Data["name"].(string)).IAM().Policy(ctx)
	if err != nil {
		return errors.WithStack(err)
	}
	res <- output
	return nil
}
