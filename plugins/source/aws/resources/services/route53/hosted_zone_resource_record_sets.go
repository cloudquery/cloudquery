package route53

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/route53/models"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func hostedZoneResourceRecordSets() *schema.Table {
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

func fetchRoute53HostedZoneResourceRecordSets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(*models.Route53HostedZoneWrapper)
	svc := meta.(*client.Client).Services().Route53
	config := route53.ListResourceRecordSetsInput{HostedZoneId: r.Id}
	for {
		response, err := svc.ListResourceRecordSets(ctx, &config, func(options *route53.Options) {})
		if err != nil {
			return err
		}

		res <- response.ResourceRecordSets
		if !response.IsTruncated {
			break
		}

		config.StartRecordIdentifier = response.NextRecordIdentifier
		config.StartRecordType = response.NextRecordType
		config.StartRecordName = response.NextRecordName
	}

	return nil
}
