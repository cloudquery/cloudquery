package emr

import (
	"context"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/service/emr"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func steps() *schema.Table {
	tableName := "aws_emr_steps"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/emr/latest/APIReference/API_Step.html`,
		Resolver:            fetchEmrSteps,
		PreResourceResolver: getStep,
		Transform:           transformers.TransformWithStruct(&types.Step{}, transformers.WithPrimaryKeyComponents("Id")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "cluster_arn",
				Description:         "The Amazon Resource Name (ARN) of the EMR Cluster.",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("arn"),
				PrimaryKeyComponent: true,
			},
		},
	}
}

func fetchEmrSteps(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceEmr).Emr
	p := parent.Item.(*types.Cluster)
	paginator := emr.NewListStepsPaginator(svc, &emr.ListStepsInput{ClusterId: p.Id})
	for paginator.HasMorePages() {
		response, err := paginator.NextPage(ctx, func(options *emr.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- response.Steps
	}
	return nil
}

func getStep(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceEmr).Emr
	p := resource.Parent.Item.(*types.Cluster)
	stepSummary := resource.Item.(types.StepSummary)
	response, err := svc.DescribeStep(ctx, &emr.DescribeStepInput{ClusterId: p.Id, StepId: stepSummary.Id}, func(options *emr.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	resource.Item = response.Step
	return nil
}
