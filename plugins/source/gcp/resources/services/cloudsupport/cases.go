package cloudsupport

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
	pb "google.golang.org/api/cloudsupport/v2beta"
)

func Cases() *schema.Table {
	return &schema.Table{
		Name:        "gcp_cloudsupport_cases",
		Description: `https://cloud.google.com/support/docs/reference/rest/v2beta/cases#Case`,
		Resolver:    fetchCases,
		Multiplex:   client.ProjectMultiplexEnabledServices("cloudsupport.googleapis.com"),
		Transform:   transformers.TransformWithStruct(&pb.Case{}, client.Options()...),
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
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
