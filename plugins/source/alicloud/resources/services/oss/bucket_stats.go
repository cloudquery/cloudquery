package oss

import (
	"reflect"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/alicloud/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func BucketStats() *schema.Table {
	return &schema.Table{
		Name:     "alicloud_oss_bucket_stats",
		Title:    "Alibaba Cloud Object Storage Service (OSS) Bucket Stats",
		Resolver: fetchOssBucketStats,
		Transform: transformers.TransformWithStruct(
			&oss.BucketStat{},
			transformers.WithTypeTransformer(func(f reflect.StructField) (arrow.DataType, error) {
				if f.Name == "LastModifiedTime" {
					return arrow.FixedWidthTypes.Timestamp_us, nil
				}
				return nil, nil
			}),
		),
		Columns: []schema.Column{
			{
				Name:       "bucket_name",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("name"),
				PrimaryKey: true,
			},
			{
				Name:       "account_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveAccount,
				PrimaryKey: true,
			},
		},
	}
}
