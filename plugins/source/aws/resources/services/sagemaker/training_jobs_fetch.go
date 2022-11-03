package sagemaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchSagemakerTrainingJobs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Sagemaker
	config := sagemaker.ListTrainingJobsInput{}

	for {
		response, err := svc.ListTrainingJobs(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.TrainingJobSummaries
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
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

func resolveSagemakerTrainingJobTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	r := resource.Item.(*sagemaker.DescribeTrainingJobOutput)
	if r == nil {
		return nil
	}

	c := meta.(*client.Client)
	svc := c.Services().Sagemaker
	config := sagemaker.ListTagsInput{
		ResourceArn: r.TrainingJobArn,
	}
	response, err := svc.ListTags(ctx, &config)
	if err != nil {
		return err
	}

	return resource.Set("tags", client.TagsToMap(response.Tags))
}
