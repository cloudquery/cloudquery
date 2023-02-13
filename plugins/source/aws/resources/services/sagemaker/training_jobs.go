package sagemaker

import (
	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func TrainingJobs() *schema.Table {
	return &schema.Table{
		Name:                "aws_sagemaker_training_jobs",
		Description:         `https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DescribeTrainingJob.html`,
		Resolver:            fetchSagemakerTrainingJobs,
		PreResourceResolver: getTrainingJob,
		Transform:           transformers.TransformWithStruct(&sagemaker.DescribeTrainingJobOutput{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer("api.sagemaker"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TrainingJobArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:        "tags",
				Type:        schema.TypeJSON,
				Resolver:    resolveSagemakerTrainingJobTags,
				Description: `The tags associated with the model.`,
			},
		},
	}
}
