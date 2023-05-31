package oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/alicloud/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
)

func Buckets() *schema.Table {
	return &schema.Table{
		Name:      "alicloud_oss_buckets",
		Title:     "Alibaba Cloud Object Storage Service (OSS) Buckets",
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
				Name:       "account_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveAccount,
				PrimaryKey: true,
			},
		},
		Relations: []*schema.Table{
			BucketStats(),
		},
	}
}
