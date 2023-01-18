package lightsail

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/lightsail/models"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func DatabaseLogEvents() *schema.Table {
	return &schema.Table{
		Name:        "aws_lightsail_database_log_events",
		Description: `https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_GetRelationalDatabaseLogEvents.html`,
		Resolver:    fetchLightsailDatabaseLogEvents,
		Transform:   transformers.TransformWithStruct(&models.LogEventWrapper{}, transformers.WithUnwrapAllEmbeddedStructs()),
		Multiplex:   client.ServiceAccountRegionMultiplexer("lightsail"),
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
				Name:     "database_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}
