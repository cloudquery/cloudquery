package oss

import (
	"context"
	"fmt"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/cloudquery/cloudquery/plugins/source/alicloud/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchOssBucketStats(_ context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	bucket := parent.Item.(oss.BucketProperties)
	bucketStat, err := c.Services().OSS.GetBucketStat(bucket.Name)
	if err != nil {
		return fmt.Errorf("failed to get bucket stat for %s: %w", bucket.Name, err)
	}
	res <- bucketStat
	return nil
}
