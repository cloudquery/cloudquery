package cloudwatchlogs

import (
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func MetricFilters() *schema.Table {
	return &schema.Table{
		Name:        "aws_cloudwatchlogs_metric_filters",
		Description: `https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_MetricFilter.html`,
		Resolver:    fetchCloudwatchlogsMetricFilters,
		Multiplex:   client.ServiceAccountRegionMultiplexer("logs"),
		Transform:   transformers.TransformWithStruct(&types.MetricFilter{}),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveMetricFilterArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
