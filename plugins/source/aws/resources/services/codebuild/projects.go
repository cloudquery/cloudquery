package codebuild

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/codebuild"
	"github.com/aws/aws-sdk-go-v2/service/codebuild/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Projects() *schema.Table {
	tableName := "aws_codebuild_projects"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/codebuild/latest/APIReference/API_Project.html`,
		Resolver:    fetchCodebuildProjects,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "codebuild"),
		Transform:   transformers.TransformWithStruct(&types.Project{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}

func fetchCodebuildProjects(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Codebuild
	config := codebuild.ListProjectsInput{}
	paginator := codebuild.NewListProjectsPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		if len(page.Projects) == 0 {
			continue
		}
		projectsOutput, err := svc.BatchGetProjects(ctx, &codebuild.BatchGetProjectsInput{Names: page.Projects})
		if err != nil {
			return err
		}
		res <- projectsOutput.Projects
	}
	return nil
}
