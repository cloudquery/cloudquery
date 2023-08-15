package codebuild

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/codebuild"
	"github.com/aws/aws-sdk-go-v2/service/codebuild/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func builds() *schema.Table {
	return &schema.Table{
		Name:        "aws_codebuild_builds",
		Description: `https://docs.aws.amazon.com/codebuild/latest/APIReference/API_Build.html`,
		Resolver:    fetchBuildsForProject,
		Transform:   transformers.TransformWithStruct(&types.Build{}, transformers.WithPrimaryKeys("Arn")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
		},
	}
}

func fetchBuildsForProject(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Codebuild
	project := parent.Item.(types.Project)
	config := codebuild.ListBuildsForProjectInput{
		ProjectName: project.Name,
	}
	paginator := codebuild.NewListBuildsForProjectPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *codebuild.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		if len(page.Ids) == 0 {
			continue
		}
		buildsOutput, err := svc.BatchGetBuilds(ctx, &codebuild.BatchGetBuildsInput{Ids: page.Ids}, func(options *codebuild.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- buildsOutput.Builds
	}
	return nil
}
