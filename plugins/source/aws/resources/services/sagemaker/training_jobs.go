package sagemaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func TrainingJobs() *schema.Table {
	tableName := "aws_sagemaker_training_jobs"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DescribeTrainingJob.html`,
		Resolver:            fetchSagemakerTrainingJobs,
		PreResourceResolver: getTrainingJob,
		Transform:           transformers.TransformWithStruct(&sagemaker.DescribeTrainingJobOutput{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "api.sagemaker"),
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

func fetchSagemakerTrainingJobs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Sagemaker
	config := sagemaker.ListTrainingJobsInput{}
	paginator := sagemaker.NewListTrainingJobsPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.TrainingJobSummaries
	}
	return nil
}

func getTrainingJob(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Sagemaker
	n := resource.Item.(types.TrainingJobSummary)
	config := sagemaker.DescribeTrainingJobInput{
		TrainingJobName: n.TrainingJobName,
	}
	response, err := svc.DescribeTrainingJob(ctx, &config)
	if err != nil {
		return err
	}
	resource.Item = response
	return nil
}

func resolveSagemakerTrainingJobTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, col schema.Column) error {
	r := resource.Item.(*sagemaker.DescribeTrainingJobOutput)
	if r == nil {
		return nil
	}
	svc := meta.(*client.Client).Services().Sagemaker
	config := sagemaker.ListTagsInput{ResourceArn: r.TrainingJobArn}
	paginator := sagemaker.NewListTagsPaginator(svc, &config)
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
