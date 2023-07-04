package container

import (
	pb "cloud.google.com/go/container/apiv1/containerpb"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func Clusters() *schema.Table {
	return &schema.Table{
		Name:        "gcp_container_clusters",
		Description: `https://cloud.google.com/kubernetes-engine/docs/reference/rest/v1/projects.locations.clusters#Cluster`,
		Resolver:    fetchClusters,
		Multiplex:   client.ProjectMultiplexEnabledServices("container.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.Cluster{}, transformers.WithPrimaryKeys("SelfLink")),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     arrow.BinaryTypes.String,
				Resolver: client.ResolveProject,
			},
		},
	}
}
