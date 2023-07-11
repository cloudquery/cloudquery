package route53

import (
	"context"
	"fmt"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/route53/models"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func hostedZoneTrafficPolicyInstances() *schema.Table {
	tableName := "aws_route53_hosted_zone_traffic_policy_instances"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/Route53/latest/APIReference/API_TrafficPolicyInstance.html`,
		Resolver:    fetchRoute53HostedZoneTrafficPolicyInstances,
		Transform:   transformers.TransformWithStruct(&types.TrafficPolicyInstance{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "route53"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			{
				Name:        "arn",
				Type:        arrow.BinaryTypes.String,
				Resolver:    resolveRoute53HostedZoneTrafficPolicyInstancesArn,
				Description: `Amazon Resource Name (ARN) of the route53 hosted zone traffic policy instance.`,
				PrimaryKey:  true,
			},
			{
				Name:     "hosted_zone_arn",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}

func fetchRoute53HostedZoneTrafficPolicyInstances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(*models.Route53HostedZoneWrapper)
	config := route53.ListTrafficPolicyInstancesByHostedZoneInput{HostedZoneId: r.Id}
	cl := meta.(*client.Client)
	svc := cl.Services().Route53
	// No paginator available
	for {
		response, err := svc.ListTrafficPolicyInstancesByHostedZone(ctx, &config, func(options *route53.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- response.TrafficPolicyInstances
		if aws.ToString(response.TrafficPolicyInstanceNameMarker) == "" {
			break
		}
		config.TrafficPolicyInstanceNameMarker = response.TrafficPolicyInstanceNameMarker
	}
	return nil
}

func resolveRoute53HostedZoneTrafficPolicyInstancesArn(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	tp := resource.Item.(types.TrafficPolicyInstance)
	return resource.Set(c.Name, arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.Route53Service),
		Region:    "",
		AccountID: "",
		Resource:  fmt.Sprintf("trafficpolicyinstance/%s", aws.ToString(tp.Id)),
	}.String())
}
