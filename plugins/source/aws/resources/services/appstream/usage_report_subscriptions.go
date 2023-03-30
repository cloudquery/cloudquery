package appstream

import (
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func UsageReportSubscriptions() *schema.Table {
	tableName := "aws_appstream_usage_report_subscriptions"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/appstream2/latest/APIReference/API_UsageReportSubscription.html`,
		Resolver:    fetchAppstreamUsageReportSubscriptions,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "appstream2"),
		Transform:   transformers.TransformWithStruct(&types.UsageReportSubscription{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:     "s3_bucket_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("S3BucketName"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
