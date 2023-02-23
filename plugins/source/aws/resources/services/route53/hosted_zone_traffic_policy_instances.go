package route53

import (
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func HostedZoneTrafficPolicyInstances() *schema.Table {
	return &schema.Table{
		Name:        "aws_route53_hosted_zone_traffic_policy_instances",
		Description: `https://docs.aws.amazon.com/Route53/latest/APIReference/API_TrafficPolicyInstance.html`,
		Resolver:    fetchRoute53HostedZoneTrafficPolicyInstances,
		Transform:   transformers.TransformWithStruct(&types.TrafficPolicyInstance{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("route53"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			{
				Name:        "arn",
				Type:        schema.TypeString,
				Resolver:    resolveRoute53HostedZoneTrafficPolicyInstancesArn,
				Description: `Amazon Resource Name (ARN) of the route53 hosted zone traffic policy instance.`,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "hosted_zone_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}
