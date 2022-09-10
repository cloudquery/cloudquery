package codepipeline

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Pipelines() *schema.Table {
	return &schema.Table{
		Name:        "aws_codepipeline_pipelines",
		Description: "Represents the output of a GetPipeline action",
		Resolver:    fetchCodepipelinePipelines,
		Multiplex:   client.ServiceAccountRegionMultiplexer("codepipeline"),
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "tags",
				Description: "The tags associated with the pipeline.",
				Type:        schema.TypeJSON,
				Resolver:    resolveCodepipelinePipelineTags,
			},
			{
				Name:     "metadata",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Metadata"),
			},
			{
				Name:     "pipeline",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Pipeline"),
			},
			{
				Name:     "stages",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Stages"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchCodepipelinePipelines(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().CodePipeline
	config := codepipeline.ListPipelinesInput{}
	for {
		response, err := svc.ListPipelines(ctx, &config)
		if err != nil {
			return err
		}
		for i := range response.Pipelines {
			response, err := svc.GetPipeline(ctx, &codepipeline.GetPipelineInput{Name: response.Pipelines[i].Name})
			if err != nil {
				if c.IsNotFoundError(err) {
					continue
				}
				return err
			}
			res <- response
		}

		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
func resolveCodepipelinePipelineTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	pipeline := resource.Item.(*codepipeline.GetPipelineOutput)

	cl := meta.(*client.Client)
	svc := cl.Services().CodePipeline
	response, err := svc.ListTagsForResource(ctx, &codepipeline.ListTagsForResourceInput{
		ResourceArn: pipeline.Metadata.PipelineArn,
	})
	if err != nil {
		return err
	}

	return resource.Set(c.Name, client.TagsToMap(response.Tags))
}
