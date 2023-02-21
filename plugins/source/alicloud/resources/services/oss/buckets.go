package oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/cloudquery/cloudquery/plugins/source/alicloud/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Buckets() *schema.Table {
	return &schema.Table{
		Name:      "alicloud_oss_buckets",
		Multiplex: client.AccountMultiplex,
		Resolver:  fetchOssBuckets,
		Transform: transformers.TransformWithStruct(
			&oss.BucketProperties{},
			transformers.WithPrimaryKeys(
				"Name",
			),
		),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAccount,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
		Relations: []*schema.Table{
			BucketStats(),
		},
	}
}
