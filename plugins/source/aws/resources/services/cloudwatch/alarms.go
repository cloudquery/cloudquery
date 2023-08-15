package cloudwatch

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Alarms() *schema.Table {
	tableName := "aws_cloudwatch_alarms"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_MetricAlarm.html`,
		Resolver:    fetchCloudwatchAlarms,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "monitoring"),
		Transform:   transformers.TransformWithStruct(&types.MetricAlarm{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveCloudwatchAlarmTags,
			},
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("AlarmArn"),
				PrimaryKey: true,
			},
			{
				Name:     "dimensions",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveCloudwatchAlarmDimensions,
			},
		},
	}
}

func fetchCloudwatchAlarms(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config cloudwatch.DescribeAlarmsInput
	cl := meta.(*client.Client)
	svc := cl.Services().Cloudwatch
	paginator := cloudwatch.NewDescribeAlarmsPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *cloudwatch.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.MetricAlarms
	}
	return nil
}
func resolveCloudwatchAlarmDimensions(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	alarm := resource.Item.(types.MetricAlarm)
	dimensions := make(map[string]*string)
	for _, d := range alarm.Dimensions {
		dimensions[*d.Name] = d.Value
	}
	return resource.Set("dimensions", dimensions)
}

func resolveCloudwatchAlarmTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Cloudwatch
	alarm := resource.Item.(types.MetricAlarm)

	input := cloudwatch.ListTagsForResourceInput{
		ResourceARN: alarm.AlarmArn,
	}
	output, err := svc.ListTagsForResource(ctx, &input, func(options *cloudwatch.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, client.TagsToMap(output.Tags))
}
