package elasticbeanstalk

import (
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ApplicationVersions() *schema.Table {
	return &schema.Table{
		Name:        "aws_elasticbeanstalk_application_versions",
		Description: `https://docs.aws.amazon.com/elasticbeanstalk/latest/api/API_ApplicationVersionDescription.html`,
		Resolver:    fetchElasticbeanstalkApplicationVersions,
		Multiplex:   client.ServiceAccountRegionMultiplexer("elasticbeanstalk"),
		Transform:   transformers.TransformWithStruct(&types.ApplicationVersionDescription{}),
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
				Resolver: schema.PathResolver("ApplicationVersionArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
