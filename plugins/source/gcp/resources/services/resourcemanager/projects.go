package resourcemanager

import (
	pb "cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func Projects() *schema.Table {
	return &schema.Table{
		Name:        "gcp_resourcemanager_projects",
		Description: `https://cloud.google.com/resource-manager/reference/rest/v3/projects#Project`,
		Resolver:    fetchProjects,
		Multiplex:   client.ProjectMultiplexEnabledServices("cloudresourcemanager.googleapis.com"),
		Transform:   transformers.TransformWithStruct(&pb.Project{}, client.Options()...),
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
