package storage

import (
	pb "cloud.google.com/go/storage"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func Buckets() *schema.Table {
	return &schema.Table{
		Name:        "gcp_storage_buckets",
		Description: `https://cloud.google.com/storage/docs/json_api/v1/buckets#resource`,
		Resolver:    fetchBuckets,
		Multiplex:   client.ProjectMultiplexEnabledServices("storage.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.BucketAttrs{}, transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
		},
		Relations: []*schema.Table{
			BucketPolicies(),
		},
	}
}
