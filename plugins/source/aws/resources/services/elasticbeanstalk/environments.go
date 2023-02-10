package elasticbeanstalk

import (
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Environments() *schema.Table {
	return &schema.Table{
		Name:        "aws_elasticbeanstalk_environments",
		Description: `https://docs.aws.amazon.com/elasticbeanstalk/latest/APIReference/API_EnvironmentDescription.html`,
		Resolver:    fetchElasticbeanstalkEnvironments,
		Multiplex:   client.ServiceAccountRegionMultiplexer("elasticbeanstalk"),
		Transform:   transformers.TransformWithStruct(&types.EnvironmentDescription{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EnvironmentArn"),
			},
			client.DefaultRegionColumn(false),
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveElasticbeanstalkEnvironmentTags,
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EnvironmentId"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "listeners",
				Type:     schema.TypeJSON,
				Resolver: resolveElasticbeanstalkEnvironmentListeners,
			},
		},

		Relations: []*schema.Table{
			ConfigurationSettings(),
			ConfigurationOptions(),
		},
	}
}
