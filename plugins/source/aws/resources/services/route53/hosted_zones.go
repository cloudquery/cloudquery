package route53

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/route53/models"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func HostedZones() *schema.Table {
	return &schema.Table{
		Name:     "aws_route53_hosted_zones",
		Resolver: fetchRoute53HostedZones,
		Transform: transformers.TransformWithStruct(
			&models.Route53HostedZoneWrapper{},
			transformers.WithUnwrapStructFields("HostedZone"),
			transformers.WithNameTransformer(client.CreateReplaceTransformer(map[string]string{"vp_cs": "vpcs"})),
		),
		Multiplex: client.AccountMultiplex,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
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
