package sagemaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Apps() *schema.Table {
	tableName := "aws_sagemaker_apps"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DescribeApp.html`,
		Resolver:            fetchSagemakerApps,
		PreResourceResolver: getApp,
		Transform:           transformers.TransformWithStruct(&sagemaker.DescribeAppOutput{}, transformers.WithSkipFields("ResultMetadata")),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "api.sagemaker"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AppArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:        "tags",
				Type:        schema.TypeJSON,
				Resolver:    resolveSagemakerAppTags,
				Description: `The tags associated with the app.`,
			},
		},
	}
}
func fetchSagemakerApps(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Sagemaker
	paginator := sagemaker.NewListAppsPaginator(svc, &sagemaker.ListAppsInput{})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.Apps
	}
	return nil
}

func getApp(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Sagemaker
	n := resource.Item.(types.AppDetails)

	response, err := svc.DescribeApp(ctx, &sagemaker.DescribeAppInput{
		AppName:   n.AppName,
		AppType:   n.AppType,
		DomainId:  n.DomainId,
		SpaceName: n.SpaceName,
	})
	if err != nil {
		return err
	}

	resource.Item = response
	return nil
}

func resolveSagemakerAppTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	r := resource.Item.(*sagemaker.DescribeAppOutput)
	c := meta.(*client.Client)
	svc := c.Services().Sagemaker

	response, err := svc.ListTags(ctx, &sagemaker.ListTagsInput{
		ResourceArn: r.AppArn,
	})
	if err != nil {
		return err
	}

	return resource.Set("tags", client.TagsToMap(response.Tags))
}
