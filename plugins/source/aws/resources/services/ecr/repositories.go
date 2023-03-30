package ecr

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Repositories() *schema.Table {
	tableName := "aws_ecr_repositories"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AmazonECR/latest/APIReference/API_Repository.html`,
		Resolver:    fetchEcrRepositories,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "api.ecr"),
		Transform:   client.TransformWithStruct(&types.Repository{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RepositoryArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveRepositoryTags,
			},
			{
				Name:     "policy_text",
				Type:     schema.TypeJSON,
				Resolver: resolveRepositoryPolicy,
			},
		},

		Relations: []*schema.Table{
			repositoryImages(),
		},
	}
}
func fetchEcrRepositories(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	svc := meta.(*client.Client).Services().Ecr
	paginator := ecr.NewDescribeRepositoriesPaginator(svc, &ecr.DescribeRepositoriesInput{
		MaxResults: aws.Int32(1000),
	})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)

		if err != nil {
			return err
		}
		res <- page.Repositories
	}
	return nil
}

func resolveRepositoryTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	output, err := meta.(*client.Client).Services().Ecr.ListTagsForResource(ctx, &ecr.ListTagsForResourceInput{
		ResourceArn: resource.Item.(types.Repository).RepositoryArn,
	})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, client.TagsToMap(output.Tags))
}

func resolveRepositoryPolicy(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	svc := meta.(*client.Client).Services().Ecr
	repo := resource.Item.(types.Repository)
	output, err := svc.GetRepositoryPolicy(ctx, &ecr.GetRepositoryPolicyInput{
		RepositoryName: repo.RepositoryName,
		RegistryId:     repo.RegistryId,
	})
	if err != nil {
		if client.IsAWSError(err, "RepositoryPolicyNotFoundException") {
			return nil
		}
		return err
	}
	return resource.Set(c.Name, output.PolicyText)
}
