package sagemaker

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func NotebookInstances() *schema.Table {
	tableName := "aws_sagemaker_notebook_instances"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DescribeNotebookInstance.html`,
		Resolver:            fetchSagemakerNotebookInstances,
		PreResourceResolver: getNotebookInstance,
		Transform:           client.TransformWithStruct(&WrappedSageMakerNotebookInstance{}, transformers.WithUnwrapStructFields("DescribeNotebookInstanceOutput")),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "api.sagemaker"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("NotebookInstanceArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:        "tags",
				Type:        schema.TypeJSON,
				Resolver:    resolveSagemakerNotebookInstanceTags,
				Description: `The tags associated with the notebook instance.`,
			},
		},
	}
}
