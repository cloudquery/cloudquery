package autoscaling

import (
	"context"
	"errors"
	"regexp"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/aws/smithy-go"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/autoscaling/models"
	"github.com/cloudquery/plugin-sdk/schema"
)

var groupNotFoundRegex = regexp.MustCompile(`AutoScalingGroup name not found|Group .* not found`)

func fetchAutoscalingGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Autoscaling
	processGroupsBundle := func(groups []types.AutoScalingGroup) error {
		input := autoscaling.DescribeNotificationConfigurationsInput{
			MaxRecords: aws.Int32(100),
		}
		for _, h := range groups {
			input.AutoScalingGroupNames = append(input.AutoScalingGroupNames, *h.AutoScalingGroupName)
		}
		var configurations []types.NotificationConfiguration
		for {
			output, err := svc.DescribeNotificationConfigurations(ctx, &input, func(o *autoscaling.Options) {
				o.Region = c.Region
			})
			if err != nil {
				return err
			}
			configurations = append(configurations, output.NotificationConfigurations...)
			if aws.ToString(output.NextToken) == "" {
				break
			}
			input.NextToken = output.NextToken
		}
		for _, gr := range groups {
			wrapper := models.AutoScalingGroupWrapper{
				AutoScalingGroup:           gr,
				NotificationConfigurations: getNotificationConfigurationByGroupName(*gr.AutoScalingGroupName, configurations),
			}
			res <- wrapper
		}
		return nil
	}

	config := autoscaling.DescribeAutoScalingGroupsInput{}
	for {
		output, err := svc.DescribeAutoScalingGroups(ctx, &config)
		if err != nil {
			return err
		}
		groups := output.AutoScalingGroups
		for i := 0; i < len(groups); i += 255 {
			end := i + 255

			if end > len(groups) {
				end = len(groups)
			}
			t := groups[i:end]
			err := processGroupsBundle(t)
			if err != nil {
				return err
			}
		}

		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
func resolveAutoscalingGroupLoadBalancers(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(models.AutoScalingGroupWrapper)
	cl := meta.(*client.Client)
	svc := cl.Services().Autoscaling
	config := autoscaling.DescribeLoadBalancersInput{AutoScalingGroupName: p.AutoScalingGroupName}
	j := map[string]interface{}{}
	for {
		output, err := svc.DescribeLoadBalancers(ctx, &config)
		if err != nil {
			if isAutoScalingGroupNotExistsError(err) {
				return nil
			}
			return err
		}
		for _, lb := range output.LoadBalancers {
			j[*lb.LoadBalancerName] = *lb.State
		}

		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return resource.Set(c.Name, j)
}
func resolveAutoscalingGroupLoadBalancerTargetGroups(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(models.AutoScalingGroupWrapper)
	cl := meta.(*client.Client)
	svc := cl.Services().Autoscaling
	config := autoscaling.DescribeLoadBalancerTargetGroupsInput{AutoScalingGroupName: p.AutoScalingGroupName}
	j := map[string]interface{}{}
	for {
		output, err := svc.DescribeLoadBalancerTargetGroups(ctx, &config)
		if err != nil {
			if isAutoScalingGroupNotExistsError(err) {
				return nil
			}
			return err
		}
		for _, lb := range output.LoadBalancerTargetGroups {
			j[*lb.LoadBalancerTargetGroupARN] = *lb.State
		}

		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return resource.Set(c.Name, j)
}
func fetchAutoscalingGroupScalingPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(models.AutoScalingGroupWrapper)
	cl := meta.(*client.Client)
	svc := cl.Services().Autoscaling
	config := autoscaling.DescribePoliciesInput{AutoScalingGroupName: p.AutoScalingGroupName}

	for {
		output, err := svc.DescribePolicies(ctx, &config)
		if err != nil {
			if isAutoScalingGroupNotExistsError(err) {
				return nil
			}
			return err
		}
		res <- output.ScalingPolicies

		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
func fetchAutoscalingGroupLifecycleHooks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(models.AutoScalingGroupWrapper)
	cl := meta.(*client.Client)
	svc := cl.Services().Autoscaling
	config := autoscaling.DescribeLifecycleHooksInput{AutoScalingGroupName: p.AutoScalingGroupName}

	output, err := svc.DescribeLifecycleHooks(ctx, &config)
	if err != nil {
		if isAutoScalingGroupNotExistsError(err) {
			return nil
		}
		return err
	}
	res <- output.LifecycleHooks
	return nil
}

func getNotificationConfigurationByGroupName(name string, set []types.NotificationConfiguration) []types.NotificationConfiguration {
	var response []types.NotificationConfiguration
	for _, s := range set {
		if *s.AutoScalingGroupName == name {
			response = append(response, s)
		}
	}
	return response
}

func isAutoScalingGroupNotExistsError(err error) bool {
	var ae smithy.APIError
	if errors.As(err, &ae) {
		if ae.ErrorCode() == "ValidationError" && groupNotFoundRegex.MatchString(ae.ErrorMessage()) {
			return true
		}
	}
	return false
}
