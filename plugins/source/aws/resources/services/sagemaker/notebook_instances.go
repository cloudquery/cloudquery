package sagemaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
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
		Transform:           transformers.TransformWithStruct(&WrappedSageMakerNotebookInstance{}, transformers.WithUnwrapStructFields("DescribeNotebookInstanceOutput")),
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

type WrappedSageMakerNotebookInstance struct {
	*sagemaker.DescribeNotebookInstanceOutput
	NotebookInstanceArn  string
	NotebookInstanceName string
}

func fetchSagemakerNotebookInstances(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Sagemaker
	config := sagemaker.ListNotebookInstancesInput{}
	paginator := sagemaker.NewListNotebookInstancesPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}

		res <- page.NotebookInstances
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

func resolveSagemakerNotebookInstanceTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, col schema.Column) error {
	r := resource.Item.(*WrappedSageMakerNotebookInstance)
	c := meta.(*client.Client)
	svc := c.Services().Sagemaker
	config := sagemaker.ListTagsInput{
		ResourceArn: &r.NotebookInstanceArn,
	}
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
