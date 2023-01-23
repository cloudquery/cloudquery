package cloudwatch

import (
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Alarms() *schema.Table {
	return &schema.Table{
		Name:        "aws_cloudwatch_alarms",
		Description: `https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_MetricAlarm.html`,
		Resolver:    fetchCloudwatchAlarms,
		Multiplex:   client.ServiceAccountRegionMultiplexer("logs"),
		Transform:   transformers.TransformWithStruct(&types.MetricAlarm{}),
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
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveCloudwatchAlarmTags,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AlarmArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "dimensions",
				Type:     schema.TypeJSON,
				Resolver: resolveCloudwatchAlarmDimensions,
			},
		},
	}
}
