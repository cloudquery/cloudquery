package snapshots

import (
	"github.com/apache/arrow/go/v13/arrow"
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
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
		},
	}
}
