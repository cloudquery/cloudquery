package route53

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/route53/models"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func HostedZones() *schema.Table {
	return &schema.Table{
		Name:        "aws_route53_hosted_zones",
		Description: `https://docs.aws.amazon.com/Route53/latest/APIReference/API_HostedZone.html`,
		Resolver:    fetchRoute53HostedZones,
		Transform: transformers.TransformWithStruct(
			&models.Route53HostedZoneWrapper{},
			transformers.WithUnwrapStructFields("HostedZone"),
			transformers.WithNameTransformer(client.CreateReplaceTransformer(map[string]string{"vp_cs": "vpcs"})),
		),
		Multiplex: client.ServiceAccountRegionMultiplexer("route53"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveRoute53HostedZoneArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},

		Relations: []*schema.Table{
			HostedZoneQueryLoggingConfigs(),
			HostedZoneResourceRecordSets(),
			HostedZoneTrafficPolicyInstances(),
		},
	}
}
