package oss

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/alicloud/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/pkg/errors"
)

func fetchOssBuckets(_ context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	buckets, err := c.Services().OSS.ListBuckets()
	if err != nil {
		return errors.Wrap(err, "failed to list all buckets")
	}
	for _, bucket := range buckets.Buckets {
		c.Logger().Debug().Str("bucket", bucket.Name).Str("location", bucket.Location).Msg("Get bucket")
		res <- bucket
	}
	return nil
}
