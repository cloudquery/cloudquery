package sagemaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

type WrappedSageMakerModel struct {
	*sagemaker.DescribeModelOutput
	ModelArn  *string
	ModelName *string
}

func fetchSagemakerModels(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Sagemaker
	config := sagemaker.ListModelsInput{}
	for {
		response, err := svc.ListModels(ctx, &config)
		if err != nil {
			return err
		}

		res <- response.Models

		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}

func getModel(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Sagemaker
	n := resource.Item.(types.ModelSummary)

	response, err := svc.DescribeModel(ctx, &sagemaker.DescribeModelInput{
		ModelName: n.ModelName,
	})
	if err != nil {
		return err
	}

	resource.Item = &WrappedSageMakerModel{
		DescribeModelOutput: response,
		ModelArn:            n.ModelArn,
		ModelName:           n.ModelName,
	}
	return nil
}

func resolveSagemakerModelTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, col schema.Column) error {
	r := resource.Item.(*WrappedSageMakerModel)
	c := meta.(*client.Client)
	svc := c.Services().Sagemaker

	config := &sagemaker.ListTagsInput{
		ResourceArn: r.ModelArn,
	}

	paginator := sagemaker.NewListTagsPaginator(svc, config)
	var tags []types.Tag
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		tags = append(tags, page.Tags...)
	}

	return resource.Set(col.Name, client.TagsToMap(tags))
}
