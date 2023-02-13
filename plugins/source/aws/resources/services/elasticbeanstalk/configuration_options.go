package elasticbeanstalk

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/elasticbeanstalk/models"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ConfigurationOptions() *schema.Table {
	return &schema.Table{
		Name:        "aws_elasticbeanstalk_configuration_options",
		Description: `https://docs.aws.amazon.com/elasticbeanstalk/latest/api/API_ConfigurationOptionDescription.html`,
		Resolver:    fetchElasticbeanstalkConfigurationOptions,
		Multiplex:   client.ServiceAccountRegionMultiplexer("elasticbeanstalk"),
		Transform:   transformers.TransformWithStruct(&models.ConfigurationOptionDescriptionWrapper{}, transformers.WithUnwrapAllEmbeddedStructs()),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "environment_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
			},
		},
	}
}
