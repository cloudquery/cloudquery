package apprunner

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/apprunner"
	"github.com/aws/aws-sdk-go-v2/service/apprunner/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func ObservabilityConfigurations() *schema.Table {
	tableName := "aws_apprunner_observability_configurations"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/apprunner/latest/api/API_ObservabilityConfiguration.html`,
		Resolver:            fetchApprunnerObservabilityConfigurations,
		PreResourceResolver: getObservabilityConfiguration,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "apprunner"),
		Transform:           transformers.TransformWithStruct(&types.ObservabilityConfiguration{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ObservabilityConfigurationArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveApprunnerTags("ObservabilityConfigurationArn"),
			},
		},
	}
}

func fetchApprunnerObservabilityConfigurations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config apprunner.ListObservabilityConfigurationsInput
	svc := meta.(*client.Client).Services().Apprunner
	paginator := apprunner.NewListObservabilityConfigurationsPaginator(svc, &config)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- output.ObservabilityConfigurationSummaryList
	}
	return nil
}

func getObservabilityConfiguration(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Apprunner
	service := resource.Item.(types.ObservabilityConfigurationSummary)

	describeTaskDefinitionOutput, err := svc.DescribeObservabilityConfiguration(ctx, &apprunner.DescribeObservabilityConfigurationInput{ObservabilityConfigurationArn: service.ObservabilityConfigurationArn})
	if err != nil {
		return err
	}

	resource.Item = describeTaskDefinitionOutput.ObservabilityConfiguration
	return nil
}
