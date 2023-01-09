package storage

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/digitalocean/godo"
)

func Volumes() *schema.Table {
	return &schema.Table{
		Name:      "digitalocean_storage_volumes",
		Resolver:  fetchStorageVolumes,
		Transform: transformers.TransformWithStruct(&godo.Volume{}),
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
