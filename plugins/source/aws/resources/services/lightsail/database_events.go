package lightsail

import (
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func DatabaseEvents() *schema.Table {
	return &schema.Table{
		Name:        "aws_lightsail_database_events",
		Description: `https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_RelationalDatabaseEvent.html`,
		Resolver:    fetchLightsailDatabaseEvents,
		Transform:   transformers.TransformWithStruct(&types.RelationalDatabaseEvent{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("lightsail"),
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Type:        schema.TypeString,
				RetainOrder: true,
				Resolver:    client.ResolveAWSAccount,
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
