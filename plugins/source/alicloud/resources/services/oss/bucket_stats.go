package oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func BucketStats() *schema.Table {
	return &schema.Table{
		Name:     "alicloud_oss_bucket_stats",
		Resolver: fetchOssBucketStats,
		Transform: transformers.TransformWithStruct(
			&oss.BucketStat{},
			transformers.WithPrimaryKeys(
				"name",
			),
		),
		Columns: []schema.Column{},
	}
}
