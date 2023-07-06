package glue

import (
	"context"
	"fmt"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
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
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   resolveGlueWorkflowArn,
				PrimaryKey: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveGlueWorkflowTags,
			},
		},
	}
}

func fetchGlueWorkflows(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
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
	svc := cl.Services().Glue
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

func resolveGlueWorkflowArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	return resource.Set(c.Name, workflowARN(cl, aws.ToString(resource.Item.(*types.Workflow).Name)))
}

func resolveGlueWorkflowTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	result, err := svc.GetTags(ctx, &glue.GetTagsInput{
		ResourceArn: aws.String(workflowARN(cl, aws.ToString(resource.Item.(*types.Workflow).Name))),
	}, func(options *glue.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	return resource.Set(c.Name, result.Tags)
}

// ====================================================================================================================
//                                                  User Defined Helpers
// ====================================================================================================================

func workflowARN(cl *client.Client, name string) string {
	return arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.GlueService),
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  fmt.Sprintf("workflow/%s", name),
	}.String()
}
