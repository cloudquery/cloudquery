package droplets

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func neighbors() *schema.Table {
	return &schema.Table{
		Name:      "digitalocean_droplet_neighbors",
		Resolver:  fetchDropletsNeighbors,
		Transform: transformers.TransformWithStruct(&NeighborWrapper{}),
		Columns: []schema.Column{
			{
				Name:       "neighbor_id",
				Type:       arrow.PrimitiveTypes.Int64,
				Resolver:   schema.PathResolver("NeighborId"),
				PrimaryKey: true,
			},
		},
	}
}
