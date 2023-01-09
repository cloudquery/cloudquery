package iam

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
	pb "google.golang.org/api/iam/v1"
)

func ServiceAccountKeys() *schema.Table {
	return &schema.Table{
		Name:        "gcp_iam_service_account_keys",
		Description: `https://cloud.google.com/iam/docs/reference/rest/v1/projects.serviceAccounts.keys#ServiceAccountKey`,
		Resolver:    fetchServiceAccountKeys,
		Multiplex:   client.ProjectMultiplexEnabledServices("iam.googleapis.com"),
		Transform:   transformers.TransformWithStruct(&pb.ServiceAccountKey{}, append(client.Options(), transformers.WithSkipFields("PrivateKeyData", "PrivateKeyType"))...),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name:     "service_account_unique_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("unique_id"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
