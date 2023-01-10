package storage

import (
	pb "cloud.google.com/go/storage"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func Buckets() *schema.Table {
	return &schema.Table{
		Name:        "gcp_storage_buckets",
		Description: `https://cloud.google.com/storage/docs/json_api/v1/buckets#resource`,
		Resolver:    fetchBuckets,
		Multiplex:   client.ProjectMultiplexEnabledServices("storage.googleapis.com"),
		Transform:   transformers.TransformWithStruct(&pb.BucketAttrs{}, client.Options()...),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
		Relations: []*schema.Table{
			BucketPolicies(),
		},
	}
}
