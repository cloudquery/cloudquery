package clusters

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/digitalocean/godo"
)

func Clusters() *schema.Table {
	return &schema.Table{
		Name:        "digitalocean_clusters",
		Description: "https://docs.digitalocean.com/reference/api/api-reference/#operation/kubernetes_list_clusters",
		Resolver:    fetchKubernetesClusters,
		Transform:   transformers.TransformWithStruct(&godo.KubernetesCluster{}),
		Columns: []schema.Column{
			{
				Name:       "id",
				Type:        arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
		},
	}
}
