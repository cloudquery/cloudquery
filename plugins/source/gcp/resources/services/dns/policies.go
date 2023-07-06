package dns

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
	pb "google.golang.org/api/dns/v1"
)

func Policies() *schema.Table {
	return &schema.Table{
		Name:        "gcp_dns_policies",
		Description: `https://cloud.google.com/dns/docs/reference/v1/policies#resource`,
		Resolver:    fetchPolicies,
		Multiplex:   client.ProjectMultiplexEnabledServices("dns.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.Policy{}, transformers.WithPrimaryKeys("Id")),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     arrow.BinaryTypes.String,
				Resolver: client.ResolveProject,
			},
		},
	}
}
