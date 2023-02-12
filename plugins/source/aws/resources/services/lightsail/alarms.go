package lightsail

import (
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Alarms() *schema.Table {
	return &schema.Table{
		Name:        "aws_lightsail_alarms",
		Description: `https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_Alarm.html`,
		Resolver:    fetchLightsailAlarms,
		Transform:   transformers.TransformWithStruct(&types.Alarm{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("lightsail"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name: "arn",
				Type: schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
