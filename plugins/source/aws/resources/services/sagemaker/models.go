package sagemaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Models() *schema.Table {
	tableName := "aws_sagemaker_models"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DescribeModel.html`,
		Resolver:            fetchSagemakerModels,
		PreResourceResolver: getModel,
		Transform:           transformers.TransformWithStruct(&WrappedSageMakerModel{}, transformers.WithUnwrapStructFields("DescribeModelOutput")),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "api.sagemaker"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ModelArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:        "tags",
				Type:        schema.TypeJSON,
				Resolver:    resolveSagemakerModelTags,
				Description: `The tags associated with the model.`,
			},
		},
	}
}

type WrappedSageMakerModel struct {
	*sagemaker.DescribeModelOutput
	ModelArn  *string
	ModelName *string
}

func fetchSagemakerModels(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Sagemaker
	config := sagemaker.ListModelsInput{}
	paginator := sagemaker.NewListModelsPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.Models
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
