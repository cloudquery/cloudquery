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

func repositoryImageScanFindings() *schema.Table {
	return &schema.Table{
		Name:        "aws_ecr_repository_image_scan_findings",
		Description: `https://docs.aws.amazon.com/AmazonECR/latest/APIReference/API_ImageScanFindings.html`,
		Resolver:    fetchEcrRepositoryImageScanFindings,
		Transform: transformers.TransformWithStruct(&ecr.DescribeImageScanFindingsOutput{},
			transformers.WithPrimaryKeyComponents("RegistryId"),
			transformers.WithSkipFields("NextToken", "ResultMetadata"),
		),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "repository_arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("repository_arn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:                "image_digest",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("image_digest"),
				PrimaryKeyComponent: true,
			},
		},
	}
}
func fetchEcrRepositoryImageScanFindings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceEcr).Ecr
	image := parent.Item.(types.ImageDetail)
	config := ecr.DescribeImageScanFindingsInput{
		RepositoryName: image.RepositoryName,
		RegistryId:     image.RegistryId,
		ImageId:        &types.ImageIdentifier{ImageDigest: image.ImageDigest},
		MaxResults:     aws.Int32(1000),
	}

	var result *ecr.DescribeImageScanFindingsOutput
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
		if result == nil {
			result = output
			continue
		}
		appendDescribeImageScanFindingsOutput(result, output)
	}

	if result != nil {
		res <- result
	}
	return nil
}

// appendDescribeImageScanFindingsOutput will panic if dst is nil
func appendDescribeImageScanFindingsOutput(dst, src *ecr.DescribeImageScanFindingsOutput) {
	if src == nil {
		return
	}
	dst.ImageScanStatus = src.ImageScanStatus
	dst.ImageScanFindings = mergeImageScanFindings(dst.ImageScanFindings, src.ImageScanFindings)
}

func mergeImageScanFindings(dst, src *types.ImageScanFindings) *types.ImageScanFindings {
	if dst == nil {
		return src
	}

	dst.Findings = append(dst.Findings, src.Findings...)
	dst.EnhancedFindings = append(dst.EnhancedFindings, src.EnhancedFindings...)
	dst.VulnerabilitySourceUpdatedAt = src.VulnerabilitySourceUpdatedAt
	dst.ImageScanCompletedAt = src.ImageScanCompletedAt

	if dst.FindingSeverityCounts == nil {
		dst.FindingSeverityCounts = src.FindingSeverityCounts
	} else {
		for k, v := range src.FindingSeverityCounts {
			dst.FindingSeverityCounts[k] += v
		}
	}

	return dst
}
