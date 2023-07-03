package serviceusage

import (
	pb "cloud.google.com/go/serviceusage/apiv1/serviceusagepb"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func Services() *schema.Table {
	return &schema.Table{
		Name:        "gcp_serviceusage_services",
		Description: `https://cloud.google.com/service-usage/docs/reference/rest/v1/services#Service`,
		Resolver:    fetchServices,
		Multiplex:   client.ProjectMultiplexEnabledServices("serviceusage.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.Service{}, transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     arrow.BinaryTypes.String,
				Resolver: client.ResolveProject,
			},
		},
	}
}
