package eventbridge

import (
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ApiDestinations() *schema.Table {
	tableName := "aws_eventbridge_api_destinations"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_ApiDestination.html`,
		Resolver:    fetchEventbridgeApiDestinations,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "events"),
		Transform:   transformers.TransformWithStruct(&types.ApiDestination{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ApiDestinationArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
