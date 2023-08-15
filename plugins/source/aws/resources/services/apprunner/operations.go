package apprunner

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/apprunner"
	"github.com/aws/aws-sdk-go-v2/service/apprunner/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func operations() *schema.Table {
	return &schema.Table{
		Name:        "aws_apprunner_operations",
		Description: `https://docs.aws.amazon.com/apprunner/latest/api/API_OperationSummary.html`,
		Resolver:    fetchApprunnerOperations,
		Transform:   transformers.TransformWithStruct(&types.OperationSummary{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
		},
	}
}

func fetchApprunnerOperations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	paginator := apprunner.NewListOperationsPaginator(cl.Services().Apprunner,
		&apprunner.ListOperationsInput{ServiceArn: parent.Item.(*types.Service).ServiceArn})
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx, func(options *apprunner.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- output.OperationSummaryList
	}
	return nil
}
