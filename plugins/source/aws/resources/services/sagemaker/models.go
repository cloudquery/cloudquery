package sagemaker

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
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
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ModelArn"),
				PrimaryKey: true,
			},
			{
				Name:        "tags",
				Type:        sdkTypes.ExtensionTypes.JSON,
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
	cl := meta.(*client.Client)
	svc := cl.Services().Sagemaker
	config := sagemaker.ListModelsInput{}
	paginator := sagemaker.NewListModelsPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(o *sagemaker.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Models
	}
	return nil
}

func getModel(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Sagemaker
	n := resource.Item.(types.ModelSummary)

	response, err := svc.DescribeModel(ctx, &sagemaker.DescribeModelInput{
		ModelName: n.ModelName,
	}, func(o *sagemaker.Options) {
		o.Region = cl.Region
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
	cl := meta.(*client.Client)
	svc := cl.Services().Sagemaker

	config := &sagemaker.ListTagsInput{
		ResourceArn: r.ModelArn,
	}

	paginator := sagemaker.NewListTagsPaginator(svc, config)
	var tags []types.Tag
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(o *sagemaker.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		tags = append(tags, page.Tags...)
	}

	return resource.Set(col.Name, client.TagsToMap(tags))
}
