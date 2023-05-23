package dns

import (
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/api/dns/v1"
)

func ResourceRecordSets() *schema.Table {
	return &schema.Table{
		Name:        "gcp_dns_resource_record_sets",
		Description: `https://cloud.google.com/dns/docs/reference/v1/resourceRecordSets`,
		Resolver:    fetchResourceRecordSets,
		Multiplex:   client.ProjectMultiplexEnabledServices("dns.googleapis.com"),
		Transform:   client.TransformWithStruct(&dns.ResourceRecordSet{}, transformers.WithPrimaryKeys("Name", "Type")),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
