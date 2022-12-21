package codepipeline

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchCodepipelinePipelines(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Codepipeline
	config := codepipeline.ListPipelinesInput{}
	for {
		response, err := svc.ListPipelines(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.Pipelines

		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}

func getPipeline(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Codepipeline
	item := resource.Item.(types.PipelineSummary)
	response, err := svc.GetPipeline(ctx, &codepipeline.GetPipelineInput{Name: item.Name})
	if err != nil {
		return err
	}
	resource.Item = response
	return nil
}

func resolvePipelineTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	pipeline := resource.Item.(*codepipeline.GetPipelineOutput)

	cl := meta.(*client.Client)
	svc := cl.Services().Codepipeline
	response, err := svc.ListTagsForResource(ctx, &codepipeline.ListTagsForResourceInput{
		ResourceArn: pipeline.Metadata.PipelineArn,
	})
	if err != nil {
		return err
	}

	return resource.Set(c.Name, client.TagsToMap(response.Tags))
}

func resolvePipelineArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	pipeline := resource.Item.(*codepipeline.GetPipelineOutput)

	a := arn.ARN{
		Partition: cl.Partition,
		Service:   "codepipeline",
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  "pipelines/" + *pipeline.Metadata.PipelineArn,
	}

	return resource.Set(c.Name, a.String())
}
