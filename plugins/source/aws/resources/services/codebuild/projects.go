package codebuild

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
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
	for {
		response, err := svc.ListProjects(ctx, &config)
		if err != nil {
			return err
		}
		if len(response.Projects) == 0 {
			break
		}
		projectsOutput, err := svc.BatchGetProjects(ctx, &codebuild.BatchGetProjectsInput{Names: response.Projects})
		if err != nil {
			return err
		}

		res <- projectsOutput.Projects
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
