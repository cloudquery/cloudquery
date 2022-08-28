// Code generated by codegen; DO NOT EDIT.

package logging

import (
	"context"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
)

func Metrics() *schema.Table {
	return &schema.Table{
		Name:      "gcp_logging_metrics",
		Resolver:  fetchMetrics,
		Multiplex: client.ProjectMultiplex,
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name:     "bucket_options",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("BucketOptions"),
			},
			{
				Name:     "create_time",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CreateTime"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "disabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Disabled"),
			},
			{
				Name:     "filter",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Filter"),
			},
			{
				Name:     "label_extractors",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LabelExtractors"),
			},
			{
				Name:     "metric_descriptor",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("MetricDescriptor"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "update_time",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("UpdateTime"),
			},
			{
				Name:     "value_extractor",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ValueExtractor"),
			},
			{
				Name:     "version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Version"),
			},
		},
	}
}

func fetchMetrics(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		output, err := c.Services.Logging.Projects.Metrics.List("projects/" + c.ProjectId).PageToken(nextPageToken).Do()
		if err != nil {
			return errors.WithStack(err)
		}
		res <- output.Metrics

		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
