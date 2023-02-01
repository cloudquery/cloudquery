package bigtableadmin

import (
	pb "cloud.google.com/go/bigtable"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func Clusters() *schema.Table {
	return &schema.Table{
		Name:        "gcp_bigtableadmin_clusters",
		Description: `https://cloud.google.com/bigtable/docs/reference/admin/rest/v2/projects.instances.clusters#Cluster`,
		Resolver:    fetchClusters,
		Multiplex:   client.ProjectMultiplexEnabledServices("bigtableadmin.googleapis.com"),
		Transform:   transformers.TransformWithStruct(&pb.ClusterInfo{}, append(client.Options(), transformers.WithPrimaryKeys("Name"))...),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
		Relations: []*schema.Table{
			Backups(),
		},
	}
}
