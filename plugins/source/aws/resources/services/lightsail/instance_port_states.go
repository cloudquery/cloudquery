package lightsail

import (
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func InstancePortStates() *schema.Table {
	return &schema.Table{
		Name:        "aws_lightsail_instance_port_states",
		Description: `https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_InstancePortState.html`,
		Resolver:    fetchLightsailInstancePortStates,
		Transform:   transformers.TransformWithStruct(&types.InstancePortState{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("lightsail"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "instance_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}
