package livedns

import (
	"github.com/cloudquery/cloudquery/plugins/source/gandi/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
)

func LiveDNSDomains() *schema.Table {
	return &schema.Table{
		Name:     "gandi_livedns_domains",
		Resolver: fetchLiveDNSDomains,
		Columns: []schema.Column{
			{
				Name:        "sharing_id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveSharingID,
				Description: `The Sharing ID of the resource.`,
			},
			{
				Name:     "fqdn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FQDN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "domain_href",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DomainHref"),
			},
			{
				Name:     "domain_keys_href",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DomainKeysHref"),
			},
			{
				Name:     "domain_records_href",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DomainRecordsHref"),
			},
			{
				Name:     "automatic_snapshots",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AutomaticSnapshots"),
			},
		},

		Relations: []*schema.Table{
			LiveDNSDomainSnapshots(),
		},
	}
}
