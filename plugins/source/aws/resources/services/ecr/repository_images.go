package ecr

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func repositoryImages() *schema.Table {
	return &schema.Table{
		Name:        "aws_ecr_repository_images",
		Description: `https://docs.aws.amazon.com/AmazonECR/latest/APIReference/API_ImageDetail.html`,
		Resolver:    fetchEcrRepositoryImages,
		Transform:   transformers.TransformWithStruct(&types.ImageDetail{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   resolveImageArn,
				PrimaryKey: true,
			},
		},

		Relations: []*schema.Table{
			repositoryImageScanFindings(),
		},
	}
}
func fetchEcrRepositoryImages(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Ecr
	config := ecr.DescribeImagesInput{
		RepositoryName: parent.Item.(types.Repository).RepositoryName,
		MaxResults:     aws.Int32(1000),
	}
	paginator := ecr.NewDescribeImagesPaginator(svc, &config)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx, func(options *ecr.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- output.ImageDetails
	}
	return nil
}
func resolveImageArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	item := resource.Item.(types.ImageDetail)

	a := arn.ARN{
		Partition: cl.Partition,
		Service:   "ecr",
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  "repository/" + *item.RepositoryName + "/image/" + *item.ImageDigest,
	}
	return resource.Set(c.Name, a.String())
}
