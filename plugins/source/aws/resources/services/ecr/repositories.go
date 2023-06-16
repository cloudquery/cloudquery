package ecr

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Repositories() *schema.Table {
	tableName := "aws_ecr_repositories"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AmazonECR/latest/APIReference/API_Repository.html`,
		Resolver:    fetchEcrRepositories,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "api.ecr"),
		Transform:   transformers.TransformWithStruct(&types.Repository{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("RepositoryArn"),
				PrimaryKey: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveRepositoryTags,
			},
			{
				Name:     "policy_text",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveRepositoryPolicy,
			},
		},

		Relations: []*schema.Table{
			repositoryImages(),
		},
	}
}
func fetchEcrRepositories(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Ecr
	paginator := ecr.NewDescribeRepositoriesPaginator(svc, &ecr.DescribeRepositoriesInput{
		MaxResults: aws.Int32(1000),
	})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *ecr.Options) {
			options.Region = cl.Region
		})

		if err != nil {
			return err
		}
		res <- page.Repositories
	}
	return nil
}

func resolveRepositoryTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Ecr
	output, err := svc.ListTagsForResource(ctx, &ecr.ListTagsForResourceInput{
		ResourceArn: resource.Item.(types.Repository).RepositoryArn,
	}, func(options *ecr.Options) {
		options.Region = meta.(*client.Client).Region
	})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, client.TagsToMap(output.Tags))
}

func resolveRepositoryPolicy(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Ecr
	repo := resource.Item.(types.Repository)
	output, err := svc.GetRepositoryPolicy(ctx, &ecr.GetRepositoryPolicyInput{
		RepositoryName: repo.RepositoryName,
		RegistryId:     repo.RegistryId,
	}, func(options *ecr.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		if client.IsAWSError(err, "RepositoryPolicyNotFoundException") {
			return nil
		}
		return err
	}
	return resource.Set(c.Name, output.PolicyText)
}
