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
)

func fetchGlueMlTransforms(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	input := glue.GetMLTransformsInput{}
	for {
		result, err := svc.GetMLTransforms(ctx, &input)
		if err != nil {
			return err
		}
		res <- result.Transforms
		if aws.ToString(result.NextToken) == "" {
			break
		}
		input.NextToken = result.NextToken
	}
	return nil
}
func resolveGlueMlTransformArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	r := resource.Item.(types.MLTransform)
	return resource.Set(c.Name, mlTransformARN(cl, &r))
}
func resolveGlueMlTransformTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	r := resource.Item.(types.MLTransform)
	result, err := svc.GetTags(ctx, &glue.GetTagsInput{
		ResourceArn: aws.String(mlTransformARN(cl, &r)),
	})
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	return resource.Set(c.Name, result.Tags)
}
func resolveMlTransformsSchema(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.MLTransform)
	j := make(map[string]string)
	for _, c := range r.Schema {
		j[*c.Name] = *c.DataType
	}
	return resource.Set(c.Name, j)
}
func fetchGlueMlTransformTaskRuns(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
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
