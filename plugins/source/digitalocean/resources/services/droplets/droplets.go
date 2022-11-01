// Code generated by codegen; DO NOT EDIT.

package droplets

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func Droplets() *schema.Table {
	return &schema.Table{
		Name:     "digitalocean_droplets",
		Resolver: fetchDropletsDroplets,
		Columns: []schema.Column{
			{
				Name:     "backup_ids",
				Type:     schema.TypeIntArray,
				Resolver: schema.PathResolver("BackupIDs"),
			},
			{
				Name:     "snapshot_ids",
				Type:     schema.TypeIntArray,
				Resolver: schema.PathResolver("SnapshotIDs"),
			},
			{
				Name:     "volume_ids",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("VolumeIDs"),
			},
			{
				Name:     "id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "memory",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Memory"),
			},
			{
				Name:     "vcpus",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Vcpus"),
			},
			{
				Name:     "disk",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Disk"),
			},
			{
				Name:     "region",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Region"),
			},
			{
				Name:     "image",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Image"),
			},
			{
				Name:     "size",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Size"),
			},
			{
				Name:     "size_slug",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SizeSlug"),
			},
			{
				Name:     "next_backup_window",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("NextBackupWindow"),
			},
			{
				Name:     "features",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Features"),
			},
			{
				Name:     "locked",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Locked"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "networks",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Networks"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Created"),
			},
			{
				Name:     "kernel",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Kernel"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Tags"),
			},
			{
				Name:     "vpc_uuid",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VPCUUID"),
			},
		},

		Relations: []*schema.Table{
			Neighbors(),
		},
	}
}
