package cloudwatchlogs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func LogGroups() *schema.Table {
	return &schema.Table{
		Name:        "aws_cloudwatchlogs_log_groups",
		Description: "Represents a log group.",
		Resolver:    fetchCloudwatchlogsLogGroups,
		Multiplex:   client.ServiceAccountRegionMultiplexer("logs"),
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
				Name:        "tags",
				Description: "The tags for the log group.",
				Type:        schema.TypeJSON,
				Resolver:    ResolveCloudwatchlogsLogGroupTags,
			},
			{
				Name:            "arn",
				Description:     "The Amazon Resource Name (ARN) of the log group.",
				Type:            schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "creation_time",
				Description: "The creation time of the log group, expressed as the number of milliseconds after Jan 1, 1970 00:00:00 UTC.",
				Type:        schema.TypeInt,
			},
			{
				Name:          "kms_key_id",
				Description:   "The Amazon Resource Name (ARN) of the CMK to use when encrypting log data.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:        "log_group_name",
				Description: "The name of the log group.",
				Type:        schema.TypeString,
			},
			{
				Name:        "metric_filter_count",
				Description: "The number of metric filters.",
				Type:        schema.TypeInt,
			},
			{
				Name:          "retention_in_days",
				Description:   "The number of days to retain the log events in the specified log group",
				Type:          schema.TypeInt,
				IgnoreInTests: true,
			},
			{
				Name:        "stored_bytes",
				Description: "The number of bytes stored.",
				Type:        schema.TypeInt,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchCloudwatchlogsLogGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config cloudwatchlogs.DescribeLogGroupsInput
	c := meta.(*client.Client)
	svc := c.Services().CloudwatchLogs
	for {
		response, err := svc.DescribeLogGroups(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.LogGroups
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
func ResolveCloudwatchlogsLogGroupTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	lg := resource.Item.(types.LogGroup)
	cl := meta.(*client.Client)
	svc := cl.Services().CloudwatchLogs
	out, err := svc.ListTagsLogGroup(ctx, &cloudwatchlogs.ListTagsLogGroupInput{LogGroupName: lg.LogGroupName})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, out.Tags)
}
