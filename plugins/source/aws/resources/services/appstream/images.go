package appstream

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appstream"
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Images() *schema.Table {
	tableName := "aws_appstream_images"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/appstream2/latest/APIReference/API_Image.html`,
		Resolver:    fetchAppstreamImages,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "appstream2"),
		Transform:   transformers.TransformWithStruct(&types.Image{}, transformers.WithPrimaryKeys("Arn")),
		Columns:     []schema.Column{client.DefaultAccountIDColumn(true), client.DefaultRegionColumn(true)},
	}
}

func fetchAppstreamImages(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	input := appstream.DescribeImagesInput{MaxResults: aws.Int32(25)}
	c := meta.(*client.Client)
	svc := c.Services().Appstream
	for {
		response, err := svc.DescribeImages(ctx, &input)
		if err != nil {
			return err
		}
		res <- response.Images

		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}
