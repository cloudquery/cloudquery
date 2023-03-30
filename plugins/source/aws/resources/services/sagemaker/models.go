package sagemaker

import (
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
		Transform:           client.TransformWithStruct(&WrappedSageMakerModel{}, transformers.WithUnwrapStructFields("DescribeModelOutput")),
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
