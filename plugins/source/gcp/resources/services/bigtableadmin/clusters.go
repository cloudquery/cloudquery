package bigtableadmin

import (
	pb "cloud.google.com/go/bigtable"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func Clusters() *schema.Table {
	return &schema.Table{
		Name:        "gcp_bigtableadmin_clusters",
		Description: `https://cloud.google.com/bigtable/docs/reference/admin/rest/v2/projects.instances.clusters#Cluster`,
		Resolver:    fetchClusters,
		Multiplex:   client.ProjectMultiplexEnabledServices("bigtableadmin.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.ClusterInfo{}, transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			{
				Name:       "project_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveProject,
				PrimaryKey: true,
			},
		},
		Relations: []*schema.Table{
			Backups(),
		},
	}
}
