package apprunner

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/apprunner"
	"github.com/aws/aws-sdk-go-v2/service/apprunner/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

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
