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

func fetchEcrRepositories(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
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
		return err
	}
	return resource.Set(c.Name, output.PolicyText)
}

func fetchEcrRepositoryImages(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	maxResults := int32(1000)
	p := parent.Item.(types.Repository)
	config := ecr.DescribeImagesInput{
		RepositoryName: p.RepositoryName,
		MaxResults:     &maxResults,
	}
	c := meta.(*client.Client)
	svc := c.Services().Ecr
	for {
		output, err := svc.DescribeImages(ctx, &config)
		if err != nil {
			return err
		}
		res <- output.ImageDetails
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
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
		Resource:  "repository_image/" + *item.RegistryId + "/" + *item.ImageDigest,
	}
	return resource.Set(c.Name, a.String())
}
