package dns

import (
	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	pb "google.golang.org/api/dns/v1"
)

func ManagedZones() *schema.Table {
	return &schema.Table{
		Name:        "gcp_dns_managed_zones",
		Description: `https://cloud.google.com/dns/docs/reference/v1/managedZones#resource`,
		Resolver:    fetchManagedZones,
		Multiplex:   client.ProjectMultiplexEnabledServices("dns.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.ManagedZone{}, transformers.WithPrimaryKeys("Id")),
		Columns: []schema.Column{
			client.ProjectIDColumn(false),
		},
		Relations: []*schema.Table{
			resourceRecordSets(),
		},
	}
}
