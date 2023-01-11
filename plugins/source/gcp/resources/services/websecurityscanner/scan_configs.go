package websecurityscanner

import (
	pb "cloud.google.com/go/websecurityscanner/apiv1/websecurityscannerpb"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func ScanConfigs() *schema.Table {
	return &schema.Table{
		Name:        "gcp_websecurityscanner_scan_configs",
		Description: `https://cloud.google.com/security-command-center/docs/reference/web-security-scanner/rest/v1/projects.scanConfigs#resource:-scanconfig`,
		Resolver:    fetchScanConfigs,
		Multiplex:   client.ProjectMultiplexEnabledServices("websecurityscanner.googleapis.com"),
		Transform:   transformers.TransformWithStruct(&pb.ScanConfig{}, client.Options()...),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
		Relations: []*schema.Table{
			ScanRuns(),
		},
	}
}
