package ecr

import (
	"context"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
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
		Transform: transformers.TransformWithStruct(&types.ImageDetail{},
			transformers.WithPrimaryKeyComponents("ImageDigest", "RegistryId"),
		),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "repository_arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("arn"),
				PrimaryKeyComponent: true,
			},
		},

		Relations: schema.Tables{repositoryImageScanFindings()},
	}
}
func fetchEcrRepositoryImages(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceEcr).Ecr
	repository := parent.Item.(types.Repository)
	config := ecr.DescribeImagesInput{
		RepositoryName: repository.RepositoryName,
		RegistryId:     repository.RegistryId,
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
