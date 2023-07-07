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
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("TrainingJobArn"),
				PrimaryKey: true,
			},
			{
				Name:        "tags",
				Type:        sdkTypes.ExtensionTypes.JSON,
				Resolver:    resolveSagemakerTrainingJobTags,
				Description: `The tags associated with the model.`,
			},
		},
	}
}

func fetchSagemakerTrainingJobs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Sagemaker
	config := sagemaker.ListTrainingJobsInput{}
	paginator := sagemaker.NewListTrainingJobsPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(o *sagemaker.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.TrainingJobSummaries
	}
	return nil
}

func getTrainingJob(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Sagemaker
	n := resource.Item.(types.TrainingJobSummary)
	config := sagemaker.DescribeTrainingJobInput{
		TrainingJobName: n.TrainingJobName,
	}
	response, err := svc.DescribeTrainingJob(ctx, &config, func(o *sagemaker.Options) {
		o.Region = cl.Region
	})
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
	cl := meta.(*client.Client)
	svc := cl.Services().Sagemaker
	config := sagemaker.ListTagsInput{ResourceArn: r.TrainingJobArn}
	paginator := sagemaker.NewListTagsPaginator(svc, &config)
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
