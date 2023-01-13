package oss

import (
	"context"
	"fmt"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/cloudquery/cloudquery/plugins/source/alicloud/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/pkg/errors"
)

func fetchOssBucketStats(_ context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cli := meta.(*client.Client)
	bucket := parent.Item.(oss.BucketProperties)
	ossCli, err := cli.GetOSSClient(bucket.Location)
	if err != nil {
		return errors.Wrap(err, "Failed to get oss client")
	}
	bucketStat, err := ossCli.GetBucketStat(bucket.Name)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Failed to get bucket: %s stat", bucket.Name))
	}
	res <- bucketStat
	return nil
}

func fetchOssBuckets(_ context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cli := meta.(*client.Client)
	buckets, err := cli.OSSClient.ListBuckets()
	if err != nil {
		return errors.Wrap(err, "failed to list all buckets")
	}
	for _, bucket := range buckets.Buckets {
		cli.Logger().Info().Str("bucket: ", bucket.Name).Str("location: ", bucket.Location).Msg("Get bucket")
		res <- bucket
	}
	return nil
}
