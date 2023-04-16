package kms

import (
	pb "cloud.google.com/go/kms/apiv1/kmspb"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func ImportJobs() *schema.Table {
	return &schema.Table{
		Name:        "gcp_kms_import_jobs",
		Description: `https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.importJobs#ImportJob`,
		Resolver:    fetchImportJobs,
		Multiplex:   client.ProjectMultiplexEnabledServices("cloudkms.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.ImportJob{}, transformers.WithPrimaryKeys("Name")),
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
