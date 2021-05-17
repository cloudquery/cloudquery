package resources

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Route53TrafficPolicies() *schema.Table {
	return &schema.Table{
		Name:         "aws_route53_traffic_policies",
		Resolver:     fetchRoute53TrafficPolicies,
		Multiplex:    client.AccountMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountFilter,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "resource_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
			},
			{
				Name: "latest_version",
				Type: schema.TypeInt,
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name: "traffic_policy_count",
				Type: schema.TypeInt,
			},
			{
				Name: "type",
				Type: schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "aws_route53_traffic_policy_versions",
				Resolver: fetchRoute53TrafficPolicyVersions,
				Columns: []schema.Column{
					{
						Name:     "traffic_policy_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "document",
						Type:     schema.TypeJSON,
						Resolver: resolveRoute53trafficPolicyVersionDocument,
					},
					{
						Name:     "version_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Id"),
					},
					{
						Name: "name",
						Type: schema.TypeString,
					},
					{
						Name: "type",
						Type: schema.TypeString,
					},
					{
						Name: "version",
						Type: schema.TypeInt,
					},
					{
						Name: "comment",
						Type: schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchRoute53TrafficPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
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
func fetchRoute53TrafficPolicyVersions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r, ok := parent.Item.(types.TrafficPolicySummary)
	if !ok {
		return fmt.Errorf("not route53 traffic policy")
	}
	config := route53.ListTrafficPolicyVersionsInput{Id: r.Id}
	svc := meta.(*client.Client).Services().Route53
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
func resolveRoute53trafficPolicyVersionDocument(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r, ok := resource.Item.(types.TrafficPolicy)
	if !ok {
		return fmt.Errorf("not route53 traffic policy")
	}
	var value interface{}
	err := json.Unmarshal([]byte(*r.Document), &value)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, value)
}
