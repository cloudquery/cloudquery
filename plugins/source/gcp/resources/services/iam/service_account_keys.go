package iam

import (
	pb "cloud.google.com/go/iam/admin/apiv1/adminpb"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func ServiceAccountKeys() *schema.Table {
	return &schema.Table{
		Name:        "gcp_iam_service_account_keys",
		Description: `https://cloud.google.com/iam/docs/reference/rest/v1/projects.serviceAccounts.keys#ServiceAccountKey`,
		Resolver:    fetchServiceAccountKeys,
		Multiplex:   client.ProjectMultiplexEnabledServices("iam.googleapis.com"),
		Transform:   transformers.TransformWithStruct(&pb.ServiceAccountKey{}, append(client.Options(), transformers.WithSkipFields("PrivateKeyData", "PrivateKeyType"), transformers.WithPrimaryKeys("Name"))...),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "service_account_unique_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("unique_id"),
			},
		},
	}
}
