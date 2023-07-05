package websecurityscanner

import (
	pb "cloud.google.com/go/websecurityscanner/apiv1/websecurityscannerpb"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func Findings() *schema.Table {
	return &schema.Table{
		Name:        "gcp_websecurityscanner_scan_config_scan_run_findings",
		Description: `https://cloud.google.com/security-command-center/docs/reference/web-security-scanner/rest/v1/projects.scanConfigs.scanRuns.findings`,
		Resolver:    fetchFindings,
		Multiplex:   client.ProjectMultiplexEnabledServices("websecurityscanner.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.Finding{}, transformers.WithPrimaryKeys("Name")),
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
