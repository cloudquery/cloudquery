package appflow

import (
	"context"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/service/appflow"
	"github.com/aws/aws-sdk-go-v2/service/appflow/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Flows() *schema.Table {
	tableName := "aws_appflow_flows"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/appflow/1.0/APIReference/API_DescribeFlow.html`,
		Resolver:            fetchFlows,
		PreResourceResolver: getFlow,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "appflow"),
		Transform:           transformers.TransformWithStruct(&appflow.DescribeFlowOutput{}, transformers.WithSkipFields("ResultMetadata")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("FlowArn"),
				PrimaryKeyComponent: true,
			},
		},
	}
}

func fetchFlows(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceAppflow).Appflow
	paginator := appflow.NewListFlowsPaginator(svc, nil)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx, func(o *appflow.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- output.Flows
	}
	return nil
}

func getFlow(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceAppflow).Appflow
	input := appflow.DescribeFlowInput{FlowName: resource.Item.(types.FlowDefinition).FlowName}
	output, err := svc.DescribeFlow(ctx, &input, func(o *appflow.Options) { o.Region = cl.Region })
	if err != nil {
		return err
	}
	resource.Item = output
	return nil
}
