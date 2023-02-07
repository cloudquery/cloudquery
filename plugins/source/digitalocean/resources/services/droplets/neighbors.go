package droplets

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func neighbors() *schema.Table {
	return &schema.Table{
		Name:      "digitalocean_droplet_neighbors",
		Resolver:  fetchDropletsNeighbors,
		Transform: transformers.TransformWithStruct(&NeighborWrapper{}),
		Columns: []schema.Column{
			{
				Name:     "neighbor_id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("NeighborId"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
