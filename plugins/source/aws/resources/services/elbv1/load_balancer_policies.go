package elbv1

import (
	"context"

	elbv1 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/elbv1/models"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func loadBalancerPolicies() *schema.Table {
	tableName := "aws_elbv1_load_balancer_policies"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/elasticloadbalancing/2012-06-01/APIReference/API_PolicyDescription.html`,
		Resolver:    fetchElbv1LoadBalancerPolicies,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "elasticloadbalancing"),
		Transform:   transformers.TransformWithStruct(&types.PolicyDescription{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "load_balancer_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "load_balancer_name",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("load_balancer_name"),
			},
			{
				Name:     "policy_attribute_descriptions",
				Type:     schema.TypeJSON,
				Resolver: resolveElbv1loadBalancerPolicyAttributeDescriptions,
			},
		},
	}
}

func fetchElbv1LoadBalancerPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(models.ELBv1LoadBalancerWrapper)
	c := meta.(*client.Client)
	svc := c.Services().Elasticloadbalancing
	response, err := svc.DescribeLoadBalancerPolicies(ctx, &elbv1.DescribeLoadBalancerPoliciesInput{LoadBalancerName: r.LoadBalancerName})
	if err != nil {
		return err
	}
	res <- response.PolicyDescriptions
	return nil
}
func resolveElbv1loadBalancerPolicyAttributeDescriptions(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.PolicyDescription)

	response := make(map[string]any, len(r.PolicyAttributeDescriptions))
	for _, a := range r.PolicyAttributeDescriptions {
		response[*a.AttributeName] = a.AttributeValue
	}
	return resource.Set(c.Name, response)
}
