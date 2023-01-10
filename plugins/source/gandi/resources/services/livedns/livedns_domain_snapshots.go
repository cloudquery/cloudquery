package livedns

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func LiveDNSDomainSnapshots() *schema.Table {
	return &schema.Table{
		Name:     "gandi_livedns_domain_snapshots",
		Resolver: fetchLiveDNSDomainSnapshots,
		Columns: []schema.Column{
			{
				Name:     "fqdn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("fqdn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "automatic",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Automatic"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
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
				Name:     "snapshot_href",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SnapshotHREF"),
			},
			{
				Name:     "zone_data",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ZoneData"),
			},
		},
	}
}
