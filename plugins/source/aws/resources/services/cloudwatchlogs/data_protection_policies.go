package cloudwatchlogs

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func dataProtectionPolicy() *schema.Table {
	tableName := "aws_cloudwatchlogs_log_group_data_protection_policies"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_GetDataProtectionPolicy.html`,
		Resolver:    fetchDataProtectionPolicy,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "logs"),
		Transform:   transformers.TransformWithStruct(&cloudwatchlogs.GetDataProtectionPolicyOutput{}, transformers.WithSkipFields("ResultMetadata")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:        "log_group_arn",
				Description: "The Amazon Resource Name (ARN) of the log group.",
				Type:        arrow.BinaryTypes.String,
				Resolver:    schema.ParentColumnResolver("arn"),
				PrimaryKey:  true,
			},
		},
	}
}
func fetchDataProtectionPolicy(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	lg := parent.Item.(types.LogGroup)
	if lg.DataProtectionStatus == "" { // Inactive Data Protection policy, don't attempt to fetch
		return nil
	}

	config := cloudwatchlogs.GetDataProtectionPolicyInput{
		LogGroupIdentifier: lg.LogGroupName,
	}
	cl := meta.(*client.Client)
	svc := cl.Services().Cloudwatchlogs
	resp, err := svc.GetDataProtectionPolicy(ctx, &config, func(options *cloudwatchlogs.Options) {
		options.Region = cl.Region
	})

	if err != nil {
		return err
	}
	res <- resp

	return nil
}
