package lightsail

import (
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func DatabaseSnapshots() *schema.Table {
	return &schema.Table{
		Name:        "aws_lightsail_database_snapshots",
		Description: `https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_RelationalDatabaseSnapshot.html`,
		Resolver:    fetchLightsailDatabaseSnapshots,
		Transform:   transformers.TransformWithStruct(&types.RelationalDatabaseSnapshot{}),
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
				Name: "arn",
				Type: schema.TypeString,
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
