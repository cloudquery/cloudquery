package livedns

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/gandi/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

func LiveDNSDomains() *schema.Table {
	return &schema.Table{
		Name:     "gandi_livedns_domains",
		Resolver: fetchLiveDNSDomains,
		Columns: []schema.Column{
			{
				Name:        "sharing_id",
				Type:        arrow.BinaryTypes.String,
				Resolver:    client.ResolveSharingID,
				Description: `The Sharing ID of the resource.`,
			},
			{
				Name:       "fqdn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("FQDN"),
				PrimaryKey: true,
			},
			{
				Name:     "domain_href",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("DomainHref"),
			},
			{
				Name:     "domain_keys_href",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("DomainKeysHref"),
			},
			{
				Name:     "domain_records_href",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("DomainRecordsHref"),
			},
			{
				Name:     "automatic_snapshots",
				Type:     arrow.FixedWidthTypes.Boolean,
				Resolver: schema.PathResolver("AutomaticSnapshots"),
			},
		},

		Relations: []*schema.Table{
			LiveDNSDomainSnapshots(),
		},
	}
}
