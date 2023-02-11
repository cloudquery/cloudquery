package elbv2

import (
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Listeners() *schema.Table {
	return &schema.Table{
		Name:        "aws_elbv2_listeners",
		Description: `https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_Listener.html`,
		Resolver:    fetchElbv2Listeners,
		Multiplex:   client.ServiceAccountRegionMultiplexer("elasticloadbalancing"),
		Transform:   transformers.TransformWithStruct(&types.Listener{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ListenerArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveElbv2listenerTags,
			},
		},

		Relations: []*schema.Table{
			ListenerCertificates(),
		},
	}
}
