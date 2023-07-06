package apprunner

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/apprunner"
	"github.com/aws/aws-sdk-go-v2/service/apprunner/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
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
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ObservabilityConfigurationArn"),
				PrimaryKey: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveApprunnerTags("ObservabilityConfigurationArn"),
			},
		},
	}
}

func fetchApprunnerObservabilityConfigurations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config apprunner.ListObservabilityConfigurationsInput
	cl := meta.(*client.Client)
	svc := cl.Services().Apprunner
	paginator := apprunner.NewListObservabilityConfigurationsPaginator(svc, &config)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx, func(options *apprunner.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- output.ObservabilityConfigurationSummaryList
	}
	return nil
}

func getObservabilityConfiguration(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Apprunner
	service := resource.Item.(types.ObservabilityConfigurationSummary)

	describeTaskDefinitionOutput, err := svc.DescribeObservabilityConfiguration(ctx, &apprunner.DescribeObservabilityConfigurationInput{ObservabilityConfigurationArn: service.ObservabilityConfigurationArn}, func(options *apprunner.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}

	resource.Item = describeTaskDefinitionOutput.ObservabilityConfiguration
	return nil
}
