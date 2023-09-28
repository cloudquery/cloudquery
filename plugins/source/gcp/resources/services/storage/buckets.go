package storage

import (
	pb "cloud.google.com/go/storage"
	"github.com/apache/arrow/go/v14/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Buckets() *schema.Table {
	return &schema.Table{
		Name:        "gcp_storage_buckets",
		Description: `https://pkg.go.dev/cloud.google.com/go/storage#BucketAttrs`,
		Resolver:    fetchBuckets,
		Multiplex:   client.ProjectMultiplexEnabledServices("storage.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.BucketAttrs{}, transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     arrow.BinaryTypes.String,
				Resolver: client.ResolveProject,
			},
		},
		Relations: []*schema.Table{
			BucketPolicies(),
		},
	}
}
