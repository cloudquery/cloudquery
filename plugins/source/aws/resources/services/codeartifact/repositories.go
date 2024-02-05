package codeartifact

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/codeartifact"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"
)

func Repositories() *schema.Table {
	tableName := "aws_codeartifact_repositories"
	return &schema.Table{
		Name: tableName,
		Description: `https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_RepositoryDescription.html
The 'request_account_id' and 'request_region' columns are added to show the account and region of where the request was made from.`,
		Resolver:            fetchRepositories,
		PreResourceResolver: getRepository,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "codeartifact"),
		Transform:           transformers.TransformWithStruct(&types.RepositoryDescription{}, transformers.WithPrimaryKeyComponents("Arn")),
		Columns: []schema.Column{
			client.RequestAccountIDColumn(true),
			client.RequestRegionColumn(true),
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveCodeartifactTags("Arn"),
			},
		},
		Relations: []*schema.Table{},
	}
}

func fetchRepositories(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceCodeartifact).Codeartifact
	paginator := codeartifact.NewListRepositoriesPaginator(svc, nil)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *codeartifact.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Repositories
	}
	return nil
}

func getRepository(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	repository := resource.Item.(types.RepositorySummary)
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceCodeartifact).Codeartifact
	repoOut, err := svc.DescribeRepository(ctx, &codeartifact.DescribeRepositoryInput{
		Repository:  repository.Name,
		Domain:      repository.DomainName,
		DomainOwner: repository.DomainOwner,
	}, func(options *codeartifact.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	resource.Item = repoOut.Repository
	return nil
}
