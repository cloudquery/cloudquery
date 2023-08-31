package oss

import (
	"context"
	"fmt"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/cloudquery/cloudquery/plugins/source/alicloud/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func fetchOssBucketStats(_ context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	// bucket stats must be called from the region the bucket is in
	bucket := parent.Item.(oss.BucketProperties)
	bucketRegion := strings.TrimPrefix(bucket.Location, "oss-")
	bucketClient := c.WithAccountIDAndRegion(c.AccountID, bucketRegion)
	services := bucketClient.Services()
	if services == nil {
		c.Logger().Warn().Str("bucket_name", bucket.Name).Str("region", bucketRegion).Msg("Skipping bucket stats, no client configured for region")
		return nil
	}
	bucketStat, err := services.OSS.GetBucketStat(bucket.Name)
	if err != nil {
		return fmt.Errorf("failed to get bucket stat for %s: %w", bucket.Name, err)
	}
	res <- bucketStat
	return nil
}
