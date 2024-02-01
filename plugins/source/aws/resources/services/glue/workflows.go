package glue

import (
	"context"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Workflows() *schema.Table {
	tableName := "aws_glue_workflows"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/glue/latest/webapi/API_Workflow.html`,
		Resolver:            fetchGlueWorkflows,
		PreResourceResolver: getWorkflow,
		Transform:           transformers.TransformWithStruct(&types.Workflow{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "glue"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            resolveGlueWorkflowArn,
				PrimaryKeyComponent: true,
			},
			tagsCol(func(cl *client.Client, resource *schema.Resource) string {
				return workflowARN(cl, aws.ToString(resource.Item.(*types.Workflow).Name))
			}),
		},
	}
}

func fetchGlueWorkflows(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceGlue).Glue
	paginator := glue.NewListWorkflowsPaginator(svc, &glue.ListWorkflowsInput{MaxResults: aws.Int32(25)})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *glue.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Workflows
	}
	return nil
}

func getWorkflow(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceGlue).Glue
	wf := resource.Item.(string)

	w, err := svc.GetWorkflow(ctx, &glue.GetWorkflowInput{Name: &wf}, func(options *glue.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}

	resource.Item = w.Workflow
	return nil
}

func resolveGlueWorkflowArn(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	return resource.Set(c.Name, workflowARN(cl, aws.ToString(resource.Item.(*types.Workflow).Name)))
}

func workflowARN(cl *client.Client, name string) string {
	return arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.GlueService),
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  "workflow/" + name,
	}.String()
}
