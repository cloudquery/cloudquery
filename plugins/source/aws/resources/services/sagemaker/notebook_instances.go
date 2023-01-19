package sagemaker

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func NotebookInstances() *schema.Table {
	return &schema.Table{
		Name:                "aws_sagemaker_notebook_instances",
		Description:         `https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DescribeNotebookInstance.html`,
		Resolver:            fetchSagemakerNotebookInstances,
		PreResourceResolver: getNotebookInstance,
		Transform:           transformers.TransformWithStruct(&WrappedSageMakerNotebookInstance{}, transformers.WithUnwrapStructFields("DescribeNotebookInstanceOutput")),
		Multiplex:           client.ServiceAccountRegionMultiplexer("api.sagemaker"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
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
