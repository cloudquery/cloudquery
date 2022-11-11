package elbv2

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	elbv2 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	wafv2types "github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchElbv2LoadBalancers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config elbv2.DescribeLoadBalancersInput
	c := meta.(*client.Client)
	svc := c.Services().Elasticloadbalancingv2
	for {
		response, err := svc.DescribeLoadBalancers(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.LoadBalancers
		if aws.ToString(response.NextMarker) == "" {
			break
		}
		config.Marker = response.NextMarker
	}
	return nil
}
func resolveElbv2loadBalancerWebACLArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(types.LoadBalancer)
	// only application load balancer can have web acl arn
	if p.Type != types.LoadBalancerTypeEnumApplication {
		return nil
	}
	cl := meta.(*client.Client).Services().Wafv2
	input := wafv2.GetWebACLForResourceInput{ResourceArn: p.LoadBalancerArn}
	response, err := cl.GetWebACLForResource(ctx, &input, func(options *wafv2.Options) {})
	if err != nil {
		var exc *wafv2types.WAFNonexistentItemException
		if errors.As(err, &exc) {
			if exc.ErrorCode() == "WAFNonexistentItemException" {
				return nil
			}
		}

		return err
	}
	if response.WebACL == nil {
		return nil
	}

	return resource.Set(c.Name, response.WebACL.ARN)
}
func resolveElbv2loadBalancerTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	region := cl.Region
	svc := cl.Services().Elasticloadbalancingv2
	loadBalancer := resource.Item.(types.LoadBalancer)
	tagsOutput, err := svc.DescribeTags(ctx, &elbv2.DescribeTagsInput{
		ResourceArns: []string{
			*loadBalancer.LoadBalancerArn,
		},
	}, func(o *elbv2.Options) {
		o.Region = region
	})
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	if len(tagsOutput.TagDescriptions) == 0 {
		return nil
	}
	tags := make(map[string]*string)
	for _, td := range tagsOutput.TagDescriptions {
		for _, s := range td.Tags {
			tags[*s.Key] = s.Value
		}
	}

	return resource.Set(c.Name, tags)
}

func fetchElbv2LoadBalancerAttributes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	lb := parent.Item.(types.LoadBalancer)
	c := meta.(*client.Client)
	svc := c.Services().Elasticloadbalancingv2
	result, err := svc.DescribeLoadBalancerAttributes(ctx, &elbv2.DescribeLoadBalancerAttributesInput{LoadBalancerArn: lb.LoadBalancerArn})
	if err != nil {
		return err
	}
	res <- result.Attributes
	return nil
}
