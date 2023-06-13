package livedns

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/types"
)

func LiveDNSDomainSnapshots() *schema.Table {
	return &schema.Table{
		Name:     "gandi_livedns_domain_snapshots",
		Resolver: fetchLiveDNSDomainSnapshots,
		Columns: []schema.Column{
			{
				Name:       "fqdn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("fqdn"),
				PrimaryKey: true,
			},
			{
				Name:     "automatic",
				Type:     arrow.FixedWidthTypes.Boolean,
				Resolver: schema.PathResolver("Automatic"),
			},
			{
				Name:     "created_at",
				Type:     arrow.FixedWidthTypes.Timestamp_us,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
			{
				Name:     "name",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "snapshot_href",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("SnapshotHREF"),
			},
			{
				Name:     "zone_data",
				Type:     types.ExtensionTypes.JSON,
				Resolver: schema.PathResolver("ZoneData"),
			},
		},
	}
}
