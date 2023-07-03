package storage

import (
	pb "cloud.google.com/go/iam"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func BucketPolicies() *schema.Table {
	return &schema.Table{
		Name:        "gcp_storage_bucket_policies",
		Description: `https://cloud.google.com/iam/docs/reference/rest/v1/Policy`,
		Resolver:    fetchBucketPolicies,
		Multiplex:   client.ProjectMultiplexEnabledServices("storage.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.Policy3{}),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     arrow.BinaryTypes.String,
				Resolver: client.ResolveProject,
			},
			{
				Name:     "bucket_name",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("name"),
			},
		},
	}
}
