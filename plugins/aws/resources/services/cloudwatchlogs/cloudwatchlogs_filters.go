package cloudwatchlogs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func CloudwatchlogsFilters() *schema.Table {
	return &schema.Table{
		Name:         "aws_cloudwatchlogs_filters",
		Description:  "Metric filters express how CloudWatch Logs would extract metric observations from ingested log events and transform them into metric data in a CloudWatch metric.",
		Resolver:     fetchCloudwatchlogsFilters,
		Multiplex:    client.ServiceAccountRegionMultiplexer("logs"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "region", "name", "log_group_name"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "creation_time",
				Description: "The creation time of the metric filter, expressed as the number of milliseconds after Jan 1, 1970 00:00:00 UTC.",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "name",
				Description: "The name of the metric filter.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("FilterName"),
			},
			{
				Name:        "pattern",
				Description: "A symbolic description of how CloudWatch Logs should interpret the data in each log event.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("FilterPattern"),
			},
			{
				Name:        "log_group_name",
				Description: "The name of the log group.",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_cloudwatchlogs_filter_metric_transformations",
				Description: "Indicates how to transform ingested log events to metric data in a CloudWatch metric.",
				Resolver:    fetchCloudwatchlogsFilterMetricTransformations,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"filter_cq_id", "metric_name"}},
				Columns: []schema.Column{
					{
						Name:        "filter_cq_id",
						Description: "Unique CloudQuery ID of aws_cloudwatchlogs_filters table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "metric_name",
						Description: "The name of the CloudWatch metric.",
						Type:        schema.TypeString,
					},
					{
						Name:        "metric_namespace",
						Description: "A custom namespace to contain your metric in CloudWatch.",
						Type:        schema.TypeString,
					},
					{
						Name:        "metric_value",
						Description: "The value to publish to the CloudWatch metric when a filter pattern matches a log event.",
						Type:        schema.TypeString,
					},
					{
						Name:        "default_value",
						Description: "(Optional) The value to emit when a filter pattern does not match a log event.",
						Type:        schema.TypeFloat,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchCloudwatchlogsFilters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config cloudwatchlogs.DescribeMetricFiltersInput
	c := meta.(*client.Client)
	svc := c.Services().CloudwatchLogs
	for {
		response, err := svc.DescribeMetricFilters(ctx, &config, func(options *cloudwatchlogs.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- response.MetricFilters
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
func fetchCloudwatchlogsFilterMetricTransformations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	res <- parent.Item.(types.MetricFilter).MetricTransformations
	return nil
}
