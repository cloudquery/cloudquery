package route53

import (
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func TrafficPolicies() *schema.Table {
	tableName := "aws_route53_traffic_policies"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/Route53/latest/APIReference/API_TrafficPolicySummary.html`,
		Resolver:    fetchRoute53TrafficPolicies,
		Transform:   client.TransformWithStruct(&types.TrafficPolicySummary{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "route53"),
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
