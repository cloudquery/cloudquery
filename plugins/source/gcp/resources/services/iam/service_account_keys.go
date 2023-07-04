package iam

import (
	pb "cloud.google.com/go/iam/admin/apiv1/adminpb"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func ServiceAccountKeys() *schema.Table {
	return &schema.Table{
		Name:        "gcp_iam_service_account_keys",
		Description: `https://cloud.google.com/iam/docs/reference/rest/v1/projects.serviceAccounts.keys#ServiceAccountKey`,
		Resolver:    fetchServiceAccountKeys,
		Multiplex:   client.ProjectMultiplexEnabledServices("iam.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.ServiceAccountKey{}, transformers.WithSkipFields("PrivateKeyData", "PrivateKeyType"), transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			{
				Name:       "project_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveProject,
				PrimaryKey: true,
			},
			{
				Name:     "service_account_unique_id",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("unique_id"),
			},
		},
	}
}
