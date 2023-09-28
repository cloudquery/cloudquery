package bigtableadmin

import (
	pb "cloud.google.com/go/bigtable"
	"github.com/apache/arrow/go/v14/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Instances() *schema.Table {
	return &schema.Table{
		Name:        "gcp_bigtableadmin_instances",
		Description: `https://cloud.google.com/bigtable/docs/reference/admin/rest/v2/projects.instances#Instance`,
		Resolver:    fetchInstances,
		Multiplex:   client.ProjectMultiplexEnabledServices("bigtableadmin.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.InstanceInfo{}, transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			{
				Name:       "project_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveProject,
				PrimaryKey: true,
			},
		},
		Relations: []*schema.Table{
			AppProfiles(), Clusters(), Tables(),
		},
	}
}
