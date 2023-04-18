package cloudwatchlogs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func dataProtectionPolicy() *schema.Table {
	tableName := "aws_cloudwatchlogs_log_group_data_protection_policies"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_GetDataProtectionPolicy.html`,
		Resolver:    fetchDataProtectionPolicy,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "logs"),
		Transform:   transformers.TransformWithStruct(&cloudwatchlogs.GetDataProtectionPolicyOutput{}, transformers.WithPrimaryKeys("LogGroupIdentifier"), transformers.WithSkipFields("ResultMetadata")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:        "log_group_arn",
				Description: "The Amazon Resource Name (ARN) of the log group.",
				Type:        schema.TypeString,
				Resolver:    schema.ParentColumnResolver("arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
func fetchDataProtectionPolicy(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	config := cloudwatchlogs.GetDataProtectionPolicyInput{
		LogGroupIdentifier: parent.Item.(types.LogGroup).Arn,
	}
	c := meta.(*client.Client)
	svc := c.Services().Cloudwatchlogs
	resp, err := svc.GetDataProtectionPolicy(ctx, &config)

	if err != nil {
		return err
	}
	res <- resp

	return nil
}
