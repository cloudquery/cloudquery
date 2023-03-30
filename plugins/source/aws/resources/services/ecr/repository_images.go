package ecr

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func repositoryImages() *schema.Table {
	return &schema.Table{
		Name:        "aws_ecr_repository_images",
		Description: `https://docs.aws.amazon.com/AmazonECR/latest/APIReference/API_ImageDetail.html`,
		Resolver:    fetchEcrRepositoryImages,
		Transform:   client.TransformWithStruct(&types.ImageDetail{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveImageArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},

		Relations: []*schema.Table{
			repositoryImageScanFindings(),
		},
	}
}
func fetchEcrRepositoryImages(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	config := ecr.DescribeImagesInput{
		RepositoryName: parent.Item.(types.Repository).RepositoryName,
		MaxResults:     aws.Int32(1000),
	}
	paginator := ecr.NewDescribeImagesPaginator(meta.(*client.Client).Services().Ecr, &config)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
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
