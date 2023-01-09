package codebuild

import (
	"github.com/aws/aws-sdk-go-v2/service/codebuild/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Projects() *schema.Table {
	return &schema.Table{
		Name:        "aws_codebuild_projects",
		Description: `https://docs.aws.amazon.com/codebuild/latest/APIReference/API_Project.html`,
		Resolver:    fetchCodebuildProjects,
		Multiplex:   client.ServiceAccountRegionMultiplexer("codebuild"),
		Transform:   transformers.TransformWithStruct(&types.Project{}),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
