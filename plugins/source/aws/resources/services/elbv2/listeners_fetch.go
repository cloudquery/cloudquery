package elbv2

import (
	"context"
	"regexp"

	"github.com/aws/aws-sdk-go-v2/aws"
	elbv2 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

var (
	notSupportedGatewayLB = regexp.MustCompile("This operation does not support Gateway Load Balancer Listeners")
)

func fetchElbv2Listeners(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	lb := parent.Item.(types.LoadBalancer)
	config := elbv2.DescribeListenersInput{
		LoadBalancerArn: lb.LoadBalancerArn,
	}
	c := meta.(*client.Client)
	svc := c.Services().Elasticloadbalancingv2
	for {
		response, err := svc.DescribeListeners(ctx, &config)
		if err != nil {
			if c.IsNotFoundError(err) {
				return nil
			}
			return err
		}
		res <- response.Listeners
		if aws.ToString(response.NextMarker) == "" {
			break
		}
		config.Marker = response.NextMarker
	}
	return nil
}
func resolveElbv2listenerTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	region := meta.(*client.Client).Region
	svc := meta.(*client.Client).Services().Elasticloadbalancingv2
	listener := resource.Item.(types.Listener)
	tagsOutput, err := svc.DescribeTags(ctx, &elbv2.DescribeTagsInput{
		ResourceArns: []string{
			*listener.ListenerArn,
		},
	}, func(o *elbv2.Options) {
		o.Region = region
	})
	if err != nil {
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
func fetchElbv2ListenerCertificates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	region := c.Region
	svc := c.Services().Elasticloadbalancingv2
	listener := parent.Item.(types.Listener)
	config := elbv2.DescribeListenerCertificatesInput{ListenerArn: listener.ListenerArn}
	for {
		response, err := svc.DescribeListenerCertificates(ctx, &config, func(options *elbv2.Options) {
			options.Region = region
		})
		if err != nil {
			if client.IsErrorRegex(err, "ValidationError", notSupportedGatewayLB) {
				c.Logger().Debug().Msg("ELBv2: DescribeListenerCertificates not supported for Gateway Load Balancers")
				return nil
			}
			return err
		}
		res <- response.Certificates
		if aws.ToString(response.NextMarker) == "" {
			break
		}
		config.Marker = response.NextMarker
	}
	return nil
}
