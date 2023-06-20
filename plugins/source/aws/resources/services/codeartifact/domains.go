package codeartifact

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v3/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/codebuild"
	"github.com/aws/aws-sdk-go-v2/service/codebuild/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
)

func Domains() *schema.Table {
	tableName := "aws_codeartifact_domains"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_DomainDescription.html`,
		Resolver:            fetchDomains,
		PreResourceResolver: getDomain,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "codebuild"),
		Transform:           transformers.TransformWithStruct(&types.Project{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("Arn"),
				PrimaryKey: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}

func fetchDomains(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Codeartifact
	config := codebuild.ListProjectsInput{}
	paginator := codebuild.NewListProjectsPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *codebuild.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		if len(page.Projects) == 0 {
			continue
		}
		projectsOutput, err := svc.BatchGetProjects(ctx, &codebuild.BatchGetProjectsInput{Names: page.Projects}, func(options *codebuild.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- projectsOutput.Projects
	}
	return nil
}

func getDomain(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource) error {
	return nil
}
