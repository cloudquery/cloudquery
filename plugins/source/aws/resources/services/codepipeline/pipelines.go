package codepipeline

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Pipelines() *schema.Table {
	tableName := "aws_codepipeline_pipelines"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/codepipeline/latest/APIReference/API_GetPipeline.html`,
		Resolver:            fetchCodepipelinePipelines,
		PreResourceResolver: getPipeline,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "codepipeline"),
		Transform:           transformers.TransformWithStruct(&codepipeline.GetPipelineOutput{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolvePipelineArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolvePipelineTags,
			},
		},
	}
}

func fetchCodepipelinePipelines(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Codepipeline
	config := codepipeline.ListPipelinesInput{}
	paginator := codepipeline.NewListPipelinesPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.Pipelines
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
	paginator := codepipeline.NewListTagsForResourcePaginator(svc, &codepipeline.ListTagsForResourceInput{
		ResourceArn: pipeline.Metadata.PipelineArn,
	})
	var tags []types.Tag
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		tags = append(tags, page.Tags...)
	}
	return resource.Set(c.Name, client.TagsToMap(tags))
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
