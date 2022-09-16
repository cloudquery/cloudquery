// Code generated by codegen; DO NOT EDIT.

package storage

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func Volumes() *schema.Table {
	return &schema.Table{
		Name:     "digitalocean_storage_volumes",
		Resolver: fetchStorageVolumes,
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "droplet_ids",
				Type:     schema.TypeIntArray,
				Resolver: schema.PathResolver("DropletIDs"),
			},
			{
				Name:     "region",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Region"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "size_giga_bytes",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("SizeGigaBytes"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "filesystem_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FilesystemType"),
			},
			{
				Name:     "filesystem_label",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FilesystemLabel"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Tags"),
			},
		},
	}
}
