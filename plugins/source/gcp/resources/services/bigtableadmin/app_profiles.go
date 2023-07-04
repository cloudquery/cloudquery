package bigtableadmin

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
	pb "google.golang.org/genproto/googleapis/bigtable/admin/v2"
)

func AppProfiles() *schema.Table {
	return &schema.Table{
		Name:        "gcp_bigtableadmin_app_profiles",
		Description: `https://cloud.google.com/bigtable/docs/reference/admin/rest/v2/projects.instances.appProfiles#AppProfile`,
		Resolver:    fetchAppProfiles,
		Multiplex:   client.ProjectMultiplexEnabledServices("bigtableadmin.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.AppProfile{}, transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			{
				Name:       "project_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveProject,
				PrimaryKey: true,
			},
		},
	}
}
