package dns

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
	pb "google.golang.org/api/dns/v1"
)

func ManagedZones() *schema.Table {
	return &schema.Table{
		Name:        "gcp_dns_managed_zones",
		Description: `https://cloud.google.com/dns/docs/reference/v1/managedZones#resource`,
		Resolver:    fetchManagedZones,
		Multiplex:   client.ProjectMultiplexEnabledServices("dns.googleapis.com"),
		Transform:   transformers.TransformWithStruct(&pb.ManagedZone{}, client.Options()...),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name:     "id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Id"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
