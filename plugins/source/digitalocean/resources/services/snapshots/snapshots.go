package snapshots

import (
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"github.com/digitalocean/godo"
)

func Snapshots() *schema.Table {
	return &schema.Table{
		Name:        "digitalocean_snapshots",
		Description: "https://docs.digitalocean.com/reference/api/api-reference/#tag/Snapshots",
		Resolver:    fetchSnapshotsSnapshots,
		Transform:   transformers.TransformWithStruct(&godo.Snapshot{}),
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
