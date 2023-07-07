package ecr

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/ecr/models"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func repositoryImageScanFindings() *schema.Table {
	return &schema.Table{
		Name:        "aws_ecr_repository_image_scan_findings",
		Description: `https://docs.aws.amazon.com/AmazonECR/latest/APIReference/API_ImageScanFindings.html`,
		Resolver:    fetchEcrRepositoryImageScanFindings,
		Transform:   transformers.TransformWithStruct(&models.ImageScanWrapper{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
		},
	}
}
func fetchEcrRepositoryImageScanFindings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Ecr
	image := parent.Item.(types.ImageDetail)
	repo := parent.Parent.Item.(types.Repository)
	for _, tag := range image.ImageTags {
		config := ecr.DescribeImageScanFindingsInput{
			RepositoryName: repo.RepositoryName,
			ImageId: &types.ImageIdentifier{
				ImageDigest: image.ImageDigest,
				ImageTag:    aws.String(tag),
			},
			MaxResults: aws.Int32(1000),
		}

		paginator := ecr.NewDescribeImageScanFindingsPaginator(svc, &config)
		for paginator.HasMorePages() {
			output, err := paginator.NextPage(ctx, func(options *ecr.Options) {
				options.Region = cl.Region
			})
			if err != nil {
				if client.IsAWSError(err, "ScanNotFoundException") {
					return nil
				}
				return err
			}
			res <- models.ImageScanWrapper{
				ImageScanFindings: output.ImageScanFindings,
				ImageTag:          aws.String(tag),
				ImageDigest:       image.ImageDigest,
				ImageScanStatus:   output.ImageScanStatus,
				RegistryId:        repo.RegistryId,
				RepositoryName:    repo.RepositoryName,
			}
		}
	}

	return nil
}
