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
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("AppArn"),
				PrimaryKey: true,
			},
			{
				Name:        "tags",
				Type:        sdkTypes.ExtensionTypes.JSON,
				Resolver:    resolveSagemakerAppTags,
				Description: `The tags associated with the app.`,
			},
		},
	}
}
func fetchSagemakerApps(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Sagemaker
	paginator := sagemaker.NewListAppsPaginator(svc, &sagemaker.ListAppsInput{})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(o *sagemaker.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Apps
	}
	return nil
}

func getApp(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Sagemaker
	n := resource.Item.(types.AppDetails)
	input := &sagemaker.DescribeAppInput{
		AppName:  n.AppName,
		AppType:  n.AppType,
		DomainId: n.DomainId,
	}
	if n.UserProfileName != nil {
		input.UserProfileName = n.UserProfileName
	}

	if n.SpaceName != nil {
		input.SpaceName = n.SpaceName
	}

	response, err := svc.DescribeApp(ctx, input, func(o *sagemaker.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		return err
	}

	resource.Item = response
	return nil
}

func resolveSagemakerAppTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	r := resource.Item.(*sagemaker.DescribeAppOutput)
	cl := meta.(*client.Client)
	svc := cl.Services().Sagemaker

	response, err := svc.ListTags(ctx, &sagemaker.ListTagsInput{
		ResourceArn: r.AppArn,
	}, func(o *sagemaker.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		return err
	}

	return resource.Set("tags", client.TagsToMap(response.Tags))
}
