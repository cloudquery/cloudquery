package bigtableadmin

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
	pb "google.golang.org/genproto/googleapis/bigtable/admin/v2"
)

func AppProfiles() *schema.Table {
	return &schema.Table{
		Name:        "gcp_bigtableadmin_app_profiles",
		Description: `https://cloud.google.com/bigtable/docs/reference/admin/rest/v2/projects.instances.appProfiles#AppProfile`,
		Resolver:    fetchAppProfiles,
		Multiplex:   client.ProjectMultiplexEnabledServices("bigtableadmin.googleapis.com"),
		Transform:   transformers.TransformWithStruct(&pb.AppProfile{}, client.Options()...),
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
