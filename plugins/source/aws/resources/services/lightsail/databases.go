package lightsail

import (
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Databases() *schema.Table {
	tableName := "aws_lightsail_databases"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_RelationalDatabase.html`,
		Resolver:    fetchLightsailDatabases,
		Transform:   client.TransformWithStruct(&types.RelationalDatabase{}, transformers.WithPrimaryKeys("Arn")),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "lightsail"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
		},

		Relations: []*schema.Table{
			DatabaseParameters(),
			DatabaseEvents(),
			DatabaseLogEvents(),
		},
	}
}
