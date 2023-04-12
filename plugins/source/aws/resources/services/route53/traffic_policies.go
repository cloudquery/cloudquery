package route53

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func TrafficPolicies() *schema.Table {
	tableName := "aws_route53_traffic_policies"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/Route53/latest/APIReference/API_TrafficPolicySummary.html`,
		Resolver:    fetchRoute53TrafficPolicies,
		Transform:   transformers.TransformWithStruct(&types.TrafficPolicySummary{}),
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
			trafficPolicyVersions(),
		},
	}
}

func fetchRoute53TrafficPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config route53.ListTrafficPoliciesInput
	c := meta.(*client.Client)
	svc := c.Services().Route53

	for {
		response, err := svc.ListTrafficPolicies(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.TrafficPolicySummaries

		if aws.ToString(response.TrafficPolicyIdMarker) == "" {
			break
		}
		config.TrafficPolicyIdMarker = response.TrafficPolicyIdMarker
	}
	return nil
}
func fetchRoute53TrafficPolicyVersions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(types.TrafficPolicySummary)
	config := route53.ListTrafficPolicyVersionsInput{Id: r.Id}
	svc := meta.(*client.Client).Services().Route53
	// no paginator available
	for {
		response, err := svc.ListTrafficPolicyVersions(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.TrafficPolicies
		if aws.ToString(response.TrafficPolicyVersionMarker) == "" {
			break
		}
		config.TrafficPolicyVersionMarker = response.TrafficPolicyVersionMarker
	}
	return nil
}
func resolveTrafficPolicyArn() schema.ColumnResolver {
	return client.ResolveARNGlobal(client.Route53Service, func(resource *schema.Resource) ([]string, error) {
		return []string{"trafficpolicy", *resource.Item.(types.TrafficPolicySummary).Id}, nil
	})
}
