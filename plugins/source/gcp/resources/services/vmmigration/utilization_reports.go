package vmmigration

import (
	pb "cloud.google.com/go/vmmigration/apiv1/vmmigrationpb"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func UtilizationReports() *schema.Table {
	return &schema.Table{
		Name:        "gcp_vmmigration_source_utilization_reports",
		Description: `https://cloud.google.com/migrate/virtual-machines/docs/5.0/reference/rest/v1/projects.locations.sources.utilizationReports`,
		Resolver:    fetchUtilizationReports,
		Multiplex:   client.ProjectMultiplexEnabledServices("vmmigration.googleapis.com"),
		Transform:   transformers.TransformWithStruct(&pb.UtilizationReport{}, append(client.Options(), transformers.WithPrimaryKeys("Name"))...),
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
