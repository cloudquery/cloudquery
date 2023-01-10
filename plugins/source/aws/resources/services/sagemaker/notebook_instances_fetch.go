package sagemaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

type WrappedSageMakerNotebookInstance struct {
	*sagemaker.DescribeNotebookInstanceOutput
	NotebookInstanceArn  string
	NotebookInstanceName string
}

func fetchSagemakerNotebookInstances(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Sagemaker
	config := sagemaker.ListNotebookInstancesInput{}
	for {
		response, err := svc.ListNotebookInstances(ctx, &config)
		if err != nil {
			return err
		}

		res <- response.NotebookInstances

		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}

func getNotebookInstance(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Sagemaker
	n := resource.Item.(types.NotebookInstanceSummary)

	// get more details about the notebook instance
	response, err := svc.DescribeNotebookInstance(ctx, &sagemaker.DescribeNotebookInstanceInput{
		NotebookInstanceName: n.NotebookInstanceName,
	})
	if err != nil {
		return err
	}

	resource.Item = &WrappedSageMakerNotebookInstance{
		DescribeNotebookInstanceOutput: response,
		NotebookInstanceArn:            *n.NotebookInstanceArn,
		NotebookInstanceName:           *n.NotebookInstanceName,
	}
	return nil
}

func resolveSagemakerNotebookInstanceTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	r := resource.Item.(*WrappedSageMakerNotebookInstance)
	c := meta.(*client.Client)
	svc := c.Services().Sagemaker
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
