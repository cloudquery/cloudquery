package sagemaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

type WrappedSageMakerNotebookInstance struct {
	*sagemaker.DescribeNotebookInstanceOutput
	NotebookInstanceArn  string
	NotebookInstanceName string
}

func fetchSagemakerNotebookInstances(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().SageMaker
	config := sagemaker.ListNotebookInstancesInput{}
	for {
		response, err := svc.ListNotebookInstances(ctx, &config)
		if err != nil {
			return err
		}

		// get more details about the notebook instance
		for _, n := range response.NotebookInstances {
			config := sagemaker.DescribeNotebookInstanceInput{
				NotebookInstanceName: n.NotebookInstanceName,
			}
			response, err := svc.DescribeNotebookInstance(ctx, &config, func(options *sagemaker.Options) {
				options.Region = c.Region
			})
			if err != nil {
				return err
			}

			notebook := WrappedSageMakerNotebookInstance{
				DescribeNotebookInstanceOutput: response,
				NotebookInstanceArn:            *n.NotebookInstanceArn,
				NotebookInstanceName:           *n.NotebookInstanceName,
			}

			res <- notebook
		}

		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}

func resolveSagemakerNotebookInstanceTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	r := resource.Item.(WrappedSageMakerNotebookInstance)
	c := meta.(*client.Client)
	svc := c.Services().SageMaker
	config := sagemaker.ListTagsInput{
		ResourceArn: &r.NotebookInstanceArn,
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
