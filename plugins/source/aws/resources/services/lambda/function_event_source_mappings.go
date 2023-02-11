package lambda

import (
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func FunctionEventSourceMappings() *schema.Table {
	return &schema.Table{
		Name:        "aws_lambda_function_event_source_mappings",
		Description: `https://docs.aws.amazon.com/lambda/latest/dg/API_EventSourceMappingConfiguration.html`,
		Resolver:    fetchLambdaFunctionEventSourceMappings,
		Transform:   transformers.TransformWithStruct(&types.EventSourceMappingConfiguration{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("lambda"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "function_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}
