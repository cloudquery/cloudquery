package codebuild

import (
	"github.com/aws/aws-sdk-go-v2/service/codebuild/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Projects() *schema.Table {
	tableName := "aws_codebuild_projects"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/codebuild/latest/APIReference/API_Project.html`,
		Resolver:    fetchCodebuildProjects,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "codebuild"),
		Transform:   client.TransformWithStruct(&types.Project{}),
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
