package vmmigration

import (
	pb "cloud.google.com/go/vmmigration/apiv1/vmmigrationpb"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func DatacenterConnectors() *schema.Table {
	return &schema.Table{
		Name:        "gcp_vmmigration_source_datacenter_connectors",
		Description: `https://cloud.google.com/migrate/virtual-machines/docs/5.0/reference/rest/v1/projects.locations.sources.datacenterConnectors`,
		Resolver:    fetchDatacenterConnectors,
		Multiplex:   client.ProjectMultiplexEnabledServices("vmmigration.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.DatacenterConnector{}, transformers.WithPrimaryKeys("Name")),
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
