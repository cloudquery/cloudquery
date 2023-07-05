package storage

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/digitalocean/godo"
)

func Volumes() *schema.Table {
	return &schema.Table{
		Name:        "digitalocean_storage_volumes",
		Description: "https://docs.digitalocean.com/reference/api/api-reference/#tag/Block-Storage",
		Resolver:    fetchStorageVolumes,
		Transform:   transformers.TransformWithStruct(&godo.Volume{}),
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
