package resources

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
		Resolver:     fetchCloudwatchlogsFilters,
		Multiplex:    client.AccountRegionMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
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
				Name: "creation_time",
				Type: schema.TypeBigInt,
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FilterName"),
			},
			{
				Name:     "pattern",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FilterPattern"),
			},
			{
				Name: "log_group_name",
				Type: schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "aws_cloudwatchlogs_filter_metric_transformations",
				Resolver: fetchCloudwatchlogsFilterMetricTransformations,
				Columns: []schema.Column{
					{
						Name:     "filter_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "metric_name",
						Type: schema.TypeString,
					},
					{
						Name: "metric_namespace",
						Type: schema.TypeString,
					},
					{
						Name: "metric_value",
						Type: schema.TypeString,
					},
					{
						Name: "default_value",
						Type: schema.TypeFloat,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchCloudwatchlogsFilters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
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
func fetchCloudwatchlogsFilterMetricTransformations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	res <- parent.Item.(types.MetricFilter).MetricTransformations
	return nil

}
