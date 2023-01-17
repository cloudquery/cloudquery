package oss

import (
	"reflect"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/cloudquery/cloudquery/plugins/source/alicloud/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func BucketStats() *schema.Table {
	return &schema.Table{
		Name:     "alicloud_oss_bucket_stats",
		Resolver: fetchOssBucketStats,
		Transform: transformers.TransformWithStruct(
			&oss.BucketStat{},
			transformers.WithTypeTransformer(func(f reflect.StructField) (schema.ValueType, error) {
				if f.Name == "LastModifiedTime" {
					return schema.TypeTimestamp, nil
				}
				return transformers.DefaultTypeTransformer(f)
			}),
		),
		Columns: []schema.Column{
			{
				Name:     "bucket_name",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("name"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAccount,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
