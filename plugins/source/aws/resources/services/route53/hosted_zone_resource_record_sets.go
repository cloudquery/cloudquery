package route53

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/route53/models"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
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
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}

func fetchRoute53HostedZoneResourceRecordSets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(*models.Route53HostedZoneWrapper)
	cl := meta.(*client.Client)
	svc := cl.Services().Route53
	config := route53.ListResourceRecordSetsInput{HostedZoneId: r.Id}
	for {
		response, err := svc.ListResourceRecordSets(ctx, &config, func(options *route53.Options) {}, func(options *route53.Options) {
			options.Region = cl.Region
		})
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
