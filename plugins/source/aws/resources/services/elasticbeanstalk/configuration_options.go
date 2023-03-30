package elasticbeanstalk

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/elasticbeanstalk/models"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ConfigurationOptions() *schema.Table {
	tableName := "aws_elasticbeanstalk_configuration_options"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/elasticbeanstalk/latest/api/API_ConfigurationOptionDescription.html`,
		Resolver:    fetchElasticbeanstalkConfigurationOptions,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "elasticbeanstalk"),
		Transform:   client.TransformWithStruct(&models.ConfigurationOptionDescriptionWrapper{}),
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
