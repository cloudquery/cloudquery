package vmmigration

import (
	pb "cloud.google.com/go/vmmigration/apiv1/vmmigrationpb"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func MigratingVms() *schema.Table {
	return &schema.Table{
		Name:        "gcp_vmmigration_source_migrating_vms",
		Description: `https://cloud.google.com/migrate/virtual-machines/docs/5.0/reference/rest/v1/projects.locations.sources.migratingVms`,
		Resolver:    fetchMigratingVms,
		Multiplex:   client.ProjectMultiplexEnabledServices("vmmigration.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.MigratingVm{}, transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     arrow.BinaryTypes.String,
				Resolver: client.ResolveProject,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},

		Relations: []*schema.Table{
			CloneJobs(),
			CutoverJobs(),
		},
	}
}
