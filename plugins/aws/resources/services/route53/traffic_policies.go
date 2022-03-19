package route53

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Route53TrafficPolicies() *schema.Table {
	return &schema.Table{
		Name:          "aws_route53_traffic_policies",
		Description:   "A complex type that contains information about the latest version of one traffic policy that is associated with the current AWS account.",
		Resolver:      fetchRoute53TrafficPolicies,
		Multiplex:     client.AccountMultiplex,
		IgnoreError:   client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter:  client.DeleteAccountFilter,
		Options:       schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "id"}},
		IgnoreInTests: true,
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "id",
				Description: "The ID that Amazon Route 53 assigned to the traffic policy when you created it.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Id"),
			},
			{
				Name:        "latest_version",
				Description: "The version number of the latest version of the traffic policy.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "name",
				Description: "The name that you specified for the traffic policy when you created it.",
				Type:        schema.TypeString,
			},
			{
				Name:        "traffic_policy_count",
				Description: "The number of traffic policies that are associated with the current AWS account.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "type",
				Description: "The DNS type of the resource record sets that Amazon Route 53 creates when you use a traffic policy to create a traffic policy instance.",
				Type:        schema.TypeString,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the resource.",
				Type:        schema.TypeString,
				Resolver: client.ResolveARNGlobal(client.Route53Service, func(resource *schema.Resource) ([]string, error) {
					return []string{"trafficpolicy", *resource.Item.(types.TrafficPolicySummary).Id}, nil
				}),
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_route53_traffic_policy_versions",
				Description: "A complex type that contains settings for a traffic policy.",
				Resolver:    fetchRoute53TrafficPolicyVersions,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"traffic_policy_cq_id", "id", "version"}},
				Columns: []schema.Column{
					{
						Name:        "traffic_policy_cq_id",
						Description: "Unique CloudQuery ID of aws_route53_traffic_policies table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "document",
						Description: "The definition of a traffic policy in JSON format.",
						Type:        schema.TypeJSON,
						Resolver:    resolveRoute53trafficPolicyVersionDocument,
					},
					{
						Name:        "id",
						Description: "The ID that Amazon Route 53 assigned to a traffic policy when you created it.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Id"),
					},
					{
						Name:        "name",
						Description: "The name that you specified when you created the traffic policy.",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "The DNS type of the resource record sets that Amazon Route 53 creates when you use a traffic policy to create a traffic policy instance.",
						Type:        schema.TypeString,
					},
					{
						Name:        "version",
						Description: "The version number that Amazon Route 53 assigns to a traffic policy.",
						Type:        schema.TypeInt,
					},
					{
						Name:        "comment",
						Description: "The comment that you specify in the CreateTrafficPolicy request, if any.",
						Type:        schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchRoute53TrafficPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config route53.ListTrafficPoliciesInput
	c := meta.(*client.Client)
	svc := c.Services().Route53

	for {
		response, err := svc.ListTrafficPolicies(ctx, &config)
		if err != nil {
			return diag.WrapError(err)
		}
		res <- response.TrafficPolicySummaries

		if aws.ToString(response.TrafficPolicyIdMarker) == "" {
			break
		}
		config.TrafficPolicyIdMarker = response.TrafficPolicyIdMarker
	}
	return nil
}
func fetchRoute53TrafficPolicyVersions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.TrafficPolicySummary)
	config := route53.ListTrafficPolicyVersionsInput{Id: r.Id}
	svc := meta.(*client.Client).Services().Route53
	for {
		response, err := svc.ListTrafficPolicyVersions(ctx, &config)
		if err != nil {
			return diag.WrapError(err)
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
	r := resource.Item.(types.TrafficPolicy)
	var value interface{}
	err := json.Unmarshal([]byte(*r.Document), &value)
	if err != nil {
		return diag.WrapError(err)
	}
	return resource.Set(c.Name, value)
}
