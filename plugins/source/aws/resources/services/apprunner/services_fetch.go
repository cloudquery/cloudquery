package apprunner

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/apprunner"
	"github.com/aws/aws-sdk-go-v2/service/apprunner/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchApprunnerServices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config apprunner.ListServicesInput
	svc := meta.(*client.Client).Services().Apprunner
	paginator := apprunner.NewListServicesPaginator(svc, &config)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- output.ServiceSummaryList
	}
	return nil
}

func getService(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Apprunner
	service := resource.Item.(types.ServiceSummary)

	describeTaskDefinitionOutput, err := svc.DescribeService(ctx, &apprunner.DescribeServiceInput{ServiceArn: service.ServiceArn})
	if err != nil {
		return err
	}

	resource.Item = describeTaskDefinitionOutput.Service
	return nil
}
