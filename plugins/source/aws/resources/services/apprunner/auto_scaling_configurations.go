package apprunner

import (
	"github.com/aws/aws-sdk-go-v2/service/apprunner/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func AutoScalingConfigurations() *schema.Table {
	tableName := "aws_apprunner_auto_scaling_configurations"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/apprunner/latest/api/API_AutoScalingConfiguration.html`,
		Resolver:            fetchApprunnerAutoScalingConfigurations,
		PreResourceResolver: getAutoScalingConfiguration,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "apprunner"),
		Transform:           transformers.TransformWithStruct(&types.AutoScalingConfiguration{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AutoScalingConfigurationArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveApprunnerTags("AutoScalingConfigurationArn"),
			},
		},
	}
}
