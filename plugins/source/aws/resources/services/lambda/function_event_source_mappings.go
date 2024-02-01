package lambda

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	cqtypes "github.com/cloudquery/plugin-sdk/v4/types"
)

func functionEventSourceMappings() *schema.Table {
	tableName := "aws_lambda_function_event_source_mappings"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/lambda/latest/dg/API_EventSourceMappingConfiguration.html`,
		Resolver:    fetchLambdaFunctionEventSourceMappings,
		Transform: transformers.TransformWithStruct(&types.EventSourceMappingConfiguration{},
			transformers.WithPrimaryKeyComponents("FunctionArn", "EventSourceArn"), // FunctionArn here can also be a version
		),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "uuid",
				Type:                cqtypes.ExtensionTypes.UUID,
				Resolver:            schema.PathResolver("UUID"),
				PrimaryKeyComponent: true,
			},
		},
	}
}

func fetchLambdaFunctionEventSourceMappings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	p := parent.Item.(*lambda.GetFunctionOutput)
	if p.Configuration == nil {
		return nil
	}

	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceLambda).Lambda
	config := lambda.ListEventSourceMappingsInput{FunctionName: p.Configuration.FunctionArn}
	paginator := lambda.NewListEventSourceMappingsPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *lambda.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			if cl.IsNotFoundError(err) {
				return nil
			}
			return err
		}
		res <- page.EventSourceMappings
	}
	return nil
}
