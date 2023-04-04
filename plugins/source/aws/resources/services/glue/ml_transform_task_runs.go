package glue

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func mlTransformTaskRuns() *schema.Table {
	tableName := "aws_glue_ml_transform_task_runs"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/glue/latest/webapi/API_TaskRun.html`,
		Resolver:    fetchGlueMlTransformTaskRuns,
		Transform:   transformers.TransformWithStruct(&types.TaskRun{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "glue"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "ml_transform_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}

func fetchGlueMlTransformTaskRuns(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(types.MLTransform)
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	input := glue.GetMLTaskRunsInput{
		TransformId: r.TransformId,
	}
	for {
		result, err := svc.GetMLTaskRuns(ctx, &input)
		if err != nil {
			return err
		}
		res <- result.TaskRuns
		if aws.ToString(result.NextToken) == "" {
			break
		}
		input.NextToken = result.NextToken
	}
	return nil
}

func mlTransformARN(cl *client.Client, tr *types.MLTransform) string {
	return arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.GlueService),
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  fmt.Sprintf("mlTransform/%s", aws.ToString(tr.TransformId)),
	}.String()
}
