package vpcaccess

import (
	pb "cloud.google.com/go/vpcaccess/apiv1/vpcaccesspb"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func Connectors() *schema.Table {
	return &schema.Table{
		Name:        "gcp_vpcaccess_connectors",
		Description: `https://cloud.google.com/vpc/docs/reference/vpcaccess/rest/v1/projects.locations.connectors`,
		Resolver:    fetchConnectors,
		Multiplex:   client.ProjectMultiplexEnabledServices("vpcaccess.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.Connector{}, transformers.WithPrimaryKeys("Name")),
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
