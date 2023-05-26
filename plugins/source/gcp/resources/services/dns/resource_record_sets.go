package dns

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
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
				Name:       "project_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveProject,
				PrimaryKey: true,
			},
		},
	}
}
