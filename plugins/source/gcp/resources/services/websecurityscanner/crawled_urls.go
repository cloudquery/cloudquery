package websecurityscanner

import (
	pb "cloud.google.com/go/websecurityscanner/apiv1/websecurityscannerpb"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func CrawledUrls() *schema.Table {
	return &schema.Table{
		Name:        "gcp_websecurityscanner_scan_config_scan_run_crawled_urls",
		Description: `https://cloud.google.com/security-command-center/docs/reference/web-security-scanner/rest/v1/projects.scanConfigs.scanRuns.crawledUrls/list#CrawledUrl`,
		Resolver:    fetchCrawledUrls,
		Multiplex:   client.ProjectMultiplexEnabledServices("websecurityscanner.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.CrawledUrl{}),
		Columns: []schema.Column{
			{
				Name:       "project_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveProject,
				PrimaryKey: true,
			},
			{
				Name:       "scan_run_name",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("name"),
				PrimaryKey: true,
			},
			{
				Name:       "http_method",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("HttpMethod"),
				PrimaryKey: true,
			},
			{
				Name:       "url",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("Url"),
				PrimaryKey: true,
			},
		},
	}
}
