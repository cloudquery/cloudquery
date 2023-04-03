package ecrpublic

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/ecrpublic"
	"github.com/aws/aws-sdk-go-v2/service/ecrpublic/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
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
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveImageArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchEcrpublicRepositoryImages(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	maxResults := int32(1000)
	p := parent.Item.(types.Repository)
	config := ecrpublic.DescribeImagesInput{
		RepositoryName: p.RepositoryName,
		MaxResults:     &maxResults,
	}
	c := meta.(*client.Client)
	svc := c.Services().Ecrpublic
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
		Service:   "ecr-public",
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  "repository_image/" + *item.RegistryId + "/" + *item.ImageDigest,
	}
	return resource.Set(c.Name, a.String())
}
