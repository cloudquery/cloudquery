package lightsail

import (
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func StaticIps() *schema.Table {
	return &schema.Table{
		Name:        "aws_lightsail_static_ips",
		Description: `https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_StaticIp.html`,
		Resolver:    fetchLightsailStaticIps,
		Transform:   transformers.TransformWithStruct(&types.StaticIp{}),
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
