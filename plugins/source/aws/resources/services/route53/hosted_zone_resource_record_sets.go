package route53

import (
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func HostedZoneResourceRecordSets() *schema.Table {
	tableName := "aws_route53_hosted_zone_resource_record_sets"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/Route53/latest/APIReference/API_ResourceRecordSet.html`,
		Resolver:    fetchRoute53HostedZoneResourceRecordSets,
		Transform:   transformers.TransformWithStruct(&types.ResourceRecordSet{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "route53"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:     "hosted_zone_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}
