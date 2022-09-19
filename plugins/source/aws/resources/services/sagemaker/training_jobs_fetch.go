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
	return client.ListAndDetailResolver(ctx, meta, res, listSagemakerTrainingJobs, sagemakerTrainingJobsDetail)
}

func sagemakerTrainingJobsDetail(ctx context.Context, meta schema.ClientMeta, resultsChan chan<- interface{}, errorChan chan<- error, detail interface{}) {
	c := meta.(*client.Client)
	svc := c.Services().SageMaker
	n := detail.(types.TrainingJobSummary)
	config := sagemaker.DescribeTrainingJobInput{
		TrainingJobName: n.TrainingJobName,
	}
	response, err := svc.DescribeTrainingJob(ctx, &config)
	if err != nil {
		errorChan <- err
		return
	}
	resultsChan <- response
}

func listSagemakerTrainingJobs(ctx context.Context, meta schema.ClientMeta, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().SageMaker
	config := sagemaker.ListTrainingJobsInput{}

	for {
		response, err := svc.ListTrainingJobs(ctx, &config)
		if err != nil {
			return err
		}
		for _, d := range response.TrainingJobSummaries {
			res <- d
		}
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}

func resolveSagemakerTrainingJobTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	r := resource.Item.(*sagemaker.DescribeTrainingJobOutput)
	if r == nil {
		return nil
	}

	c := meta.(*client.Client)
	svc := c.Services().SageMaker
	config := sagemaker.ListTagsInput{
		ResourceArn: r.TrainingJobArn,
	}
	response, err := svc.ListTags(ctx, &config)
	if err != nil {
		return err
	}

	tags := make(map[string]*string, len(response.Tags))
	for _, t := range response.Tags {
		tags[*t.Key] = t.Value
	}

	return resource.Set("tags", tags)
}
