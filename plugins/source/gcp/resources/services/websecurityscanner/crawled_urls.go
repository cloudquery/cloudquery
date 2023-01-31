package websecurityscanner

import (
	pb "cloud.google.com/go/websecurityscanner/apiv1/websecurityscannerpb"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func CrawledUrls() *schema.Table {
	return &schema.Table{
		Name:        "gcp_websecurityscanner_scan_config_scan_run_crawled_urls",
		Description: `https://cloud.google.com/security-command-center/docs/reference/web-security-scanner/rest/v1/projects.scanConfigs.scanRuns.crawledUrls/list#CrawledUrl`,
		Resolver:    fetchCrawledUrls,
		Multiplex:   client.ProjectMultiplexEnabledServices("websecurityscanner.googleapis.com"),
		Transform:   transformers.TransformWithStruct(&pb.CrawledUrl{}, client.Options()...),
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
				Name:     "scan_run_name",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("name"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "http_method",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HttpMethod"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Url"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
