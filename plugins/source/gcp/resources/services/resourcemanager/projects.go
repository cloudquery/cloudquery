package resourcemanager

import (
	pb "cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	"github.com/apache/arrow/go/v14/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Projects() *schema.Table {
	return &schema.Table{
		Name:        "gcp_resourcemanager_projects",
		Description: `https://cloud.google.com/resource-manager/reference/rest/v3/projects#Project`,
		Resolver:    fetchProjects,
		Multiplex:   client.ProjectMultiplexEnabledServices("cloudresourcemanager.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.Project{}, transformers.WithPrimaryKeys("Name")),
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
