package elbv1

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	elbv1 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/elbv1/models"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchElbv1LoadBalancers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Elasticloadbalancing
	processLoadBalancers := func(loadBalancers []types.LoadBalancerDescription) error {
		tagsCfg := &elbv1.DescribeTagsInput{LoadBalancerNames: make([]string, 0, len(loadBalancers))}
		for _, lb := range loadBalancers {
			tagsCfg.LoadBalancerNames = append(tagsCfg.LoadBalancerNames, *lb.LoadBalancerName)
		}
		tagsResponse, err := svc.DescribeTags(ctx, tagsCfg)
		if err != nil {
			return err
		}
		for _, lb := range loadBalancers {
			loadBalancerAttributes, err := svc.DescribeLoadBalancerAttributes(ctx, &elbv1.DescribeLoadBalancerAttributesInput{LoadBalancerName: lb.LoadBalancerName})
			if err != nil {
				if c.IsNotFoundError(err) {
					continue
				}
				return err
			}

			wrapper := models.ELBv1LoadBalancerWrapper{
				LoadBalancerDescription: lb,
				Tags:                    client.TagsToMap(getTagsByLoadBalancerName(*lb.LoadBalancerName, tagsResponse.TagDescriptions)),
				Attributes:              loadBalancerAttributes.LoadBalancerAttributes,
			}

			res <- wrapper
		}
		return nil
	}

	var config elbv1.DescribeLoadBalancersInput
	for {
		response, err := svc.DescribeLoadBalancers(ctx, &config)
		if err != nil {
			return err
		}

		for i := 0; i < len(response.LoadBalancerDescriptions); i += 20 {
			end := i + 20

			if end > len(response.LoadBalancerDescriptions) {
				end = len(response.LoadBalancerDescriptions)
			}
			loadBalancers := response.LoadBalancerDescriptions[i:end]
			if err := processLoadBalancers(loadBalancers); err != nil {
				return err
			}
		}

		if aws.ToString(response.NextMarker) == "" {
			break
		}
		config.Marker = response.NextMarker
	}

	return nil
}

func fetchElbv1LoadBalancerPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
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

	response := make(map[string]interface{}, len(r.PolicyAttributeDescriptions))
	for _, a := range r.PolicyAttributeDescriptions {
		response[*a.AttributeName] = a.AttributeValue
	}
	return resource.Set(c.Name, response)
}

func getTagsByLoadBalancerName(id string, tagsResponse []types.TagDescription) []types.Tag {
	for _, t := range tagsResponse {
		if id == *t.LoadBalancerName {
			return t.Tags
		}
	}
	return nil
}

func resolveLoadBalancerARN() schema.ColumnResolver {
	return client.ResolveARN(client.ElasticLoadBalancingService, func(resource *schema.Resource) ([]string, error) {
		return []string{"loadbalancer", *resource.Item.(models.ELBv1LoadBalancerWrapper).LoadBalancerName}, nil
	})
}
