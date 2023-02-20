package route53

import (
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func TrafficPolicies() *schema.Table {
	return &schema.Table{
		Name:        "aws_route53_traffic_policies",
		Description: `https://docs.aws.amazon.com/Route53/latest/APIReference/API_TrafficPolicySummary.html`,
		Resolver:    fetchRoute53TrafficPolicies,
		Transform:   transformers.TransformWithStruct(&types.TrafficPolicySummary{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("route53"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveTrafficPolicyArn(),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},

		Relations: []*schema.Table{
			TrafficPolicyVersions(),
		},
	}
}
