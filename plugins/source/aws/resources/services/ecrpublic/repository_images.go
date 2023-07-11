package ecrpublic

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/ecrpublic"
	"github.com/aws/aws-sdk-go-v2/service/ecrpublic/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func repositoryImages() *schema.Table {
	return &schema.Table{
		Name:        "aws_ecrpublic_repository_images",
		Description: `https://docs.aws.amazon.com/AmazonECRPublic/latest/APIReference/API_ImageDetail.html`,
		Resolver:    fetchEcrpublicRepositoryImages,
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
	}
}

func fetchEcrpublicRepositoryImages(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	p := parent.Item.(types.Repository)
	config := ecrpublic.DescribeImagesInput{
		RepositoryName: p.RepositoryName,
		MaxResults:     aws.Int32(1000),
	}
	cl := meta.(*client.Client)
	svc := cl.Services().Ecrpublic
	paginator := ecrpublic.NewDescribeImagesPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *ecrpublic.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.ImageDetails
	}
	return nil
}

func resolveImageArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	item := resource.Item.(types.ImageDetail)

	a := arn.ARN{
		Partition: cl.Partition,
		Service:   "ecr-public",
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  "repository_image/" + *item.RegistryId + "/" + *item.ImageDigest,
	}
	return resource.Set(c.Name, a.String())
}
