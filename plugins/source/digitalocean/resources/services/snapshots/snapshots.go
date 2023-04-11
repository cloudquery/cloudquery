package snapshots

import (
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/digitalocean/godo"
)

func Snapshots() *schema.Table {
	return &schema.Table{
		Name:      "digitalocean_snapshots",
		Resolver:  fetchSnapshotsSnapshots,
		Transform: transformers.TransformWithStruct(&godo.Snapshot{}),
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
