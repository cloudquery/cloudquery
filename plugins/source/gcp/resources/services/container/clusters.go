package container

import (
	pb "cloud.google.com/go/container/apiv1/containerpb"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Clusters() *schema.Table {
	return &schema.Table{
		Name:        "gcp_container_clusters",
		Description: `https://cloud.google.com/kubernetes-engine/docs/reference/rest/v1/projects.locations.clusters#Cluster`,
		Resolver:    fetchClusters,
		Multiplex:   client.ProjectMultiplexEnabledServices("container.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.Cluster{}, transformers.WithPrimaryKeys("SelfLink")),
		Columns: []schema.Column{
			client.ProjectIDColumn(false),
		},
	}
}
