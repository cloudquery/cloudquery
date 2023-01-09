package iam

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
	pb "google.golang.org/api/iam/v1"
)

func ServiceAccounts() *schema.Table {
	return &schema.Table{
		Name:        "gcp_iam_service_accounts",
		Description: `https://cloud.google.com/iam/docs/reference/rest/v1/projects.serviceAccounts#ServiceAccount`,
		Resolver:    fetchServiceAccounts,
		Multiplex:   client.ProjectMultiplexEnabledServices("iam.googleapis.com"),
		Transform:   transformers.TransformWithStruct(&pb.ServiceAccount{}, client.Options()...),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name:     "unique_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("UniqueId"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
		Relations: []*schema.Table{
			ServiceAccountKeys(),
		},
	}
}
