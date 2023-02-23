package ecr

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/ecr/models"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchEcrRepositories(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	maxResults := int32(1000)
	config := ecr.DescribeRepositoriesInput{
		MaxResults: &maxResults,
	}
	c := meta.(*client.Client)
	svc := c.Services().Ecr
	for {
		output, err := svc.DescribeRepositories(ctx, &config)
		if err != nil {
			return err
		}
		res <- output.Repositories
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}

func resolveRepositoryTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Ecr
	repo := resource.Item.(types.Repository)

	input := ecr.ListTagsForResourceInput{
		ResourceArn: repo.RepositoryArn,
	}
	output, err := svc.ListTagsForResource(ctx, &input)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, client.TagsToMap(output.Tags))
}

func resolveRepositoryPolicy(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Ecr
	repo := resource.Item.(types.Repository)

	input := ecr.GetRepositoryPolicyInput{
		RepositoryName: repo.RepositoryName,
		RegistryId:     repo.RegistryId,
	}
	output, err := svc.GetRepositoryPolicy(ctx, &input)
	if err != nil {
		if client.IsAWSError(err, "RepositoryPolicyNotFoundException") {
			return nil
		}
		return err
	}
	return resource.Set(c.Name, output.PolicyText)
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

func fetchEcrRepositoryImageScanFindings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
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

		paginator := ecr.NewDescribeImageScanFindingsPaginator(meta.(*client.Client).Services().Ecr, &config)
		for paginator.HasMorePages() {
			output, err := paginator.NextPage(ctx)
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
