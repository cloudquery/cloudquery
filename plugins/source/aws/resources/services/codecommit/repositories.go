package codecommit

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codecommit"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Repositories() *schema.Table {
	tableName := "aws_codecommit_repositories"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/codecommit/latest/APIReference/API_RepositoryMetadata.html`,
		Resolver:    fetchRepositories,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "codecommit"),
		Transform:   transformers.TransformWithStruct(&types.RepositoryMetadata{}, transformers.WithPrimaryKeys("Arn")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveCodecommitTags,
			},
		},
	}
}

func fetchRepositories(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Codecommit
	// Note: this API doesn't support limiting the number of results in a single call and the nested BatchRepositories doesn't have a listed limit
	// So we are assuming that the number of repositories is not too large and we can fetch (`BatchGet`) all of their details in a single call
	config := codecommit.ListRepositoriesInput{}
	paginator := codecommit.NewListRepositoriesPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *codecommit.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		if len(page.Repositories) == 0 {
			continue
		}
		repoNames := make([]string, len(page.Repositories))
		for i, repo := range page.Repositories {
			repoNames[i] = *repo.RepositoryName
		}
		repositoryOutput, err := svc.BatchGetRepositories(ctx, &codecommit.BatchGetRepositoriesInput{RepositoryNames: repoNames}, func(options *codecommit.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- repositoryOutput.Repositories
	}
	return nil
}

func resolveCodecommitTags(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Codecommit
	params := codecommit.ListTagsForResourceInput{ResourceArn: r.Item.(types.RepositoryMetadata).Arn}
	tags := make(map[string]string)
	for {
		output, err := svc.ListTagsForResource(ctx, &params, func(options *codecommit.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			if cl.IsNotFoundError(err) {
				return nil
			}
			return err
		}
		for key, value := range output.Tags {
			tags[key] = value
		}

		if aws.ToString(output.NextToken) == "" {
			break
		}
		params.NextToken = output.NextToken
	}
	return r.Set(c.Name, tags)
}
