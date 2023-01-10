package storage

import (
	pb "cloud.google.com/go/iam"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func BucketPolicies() *schema.Table {
	return &schema.Table{
		Name:        "gcp_storage_bucket_policies",
		Description: `https://cloud.google.com/iam/docs/reference/rest/v1/Policy`,
		Resolver:    fetchBucketPolicies,
		Multiplex:   client.ProjectMultiplexEnabledServices("storage.googleapis.com"),
		Transform:   transformers.TransformWithStruct(&pb.Policy3{}, client.Options()...),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name:     "bucket_name",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("name"),
			},
		},
	}
}
