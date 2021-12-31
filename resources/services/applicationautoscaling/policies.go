package applicationautoscaling

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling"
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func ApplicationautoscalingPolicies() *schema.Table {
	return &schema.Table{
		Name:         "aws_applicationautoscaling_policies",
		Description:  "Information about a scaling policy to use with Application Auto Scaling",
		Resolver:     fetchApplicationautoscalingPolicies,
		Multiplex:    client.ServiceAccountRegionNamespaceMultiplexer("application-autoscaling"),
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "namespace",
				Description: "The AWS Service Namespace of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSNamespace,
			},
			{
				Name:        "creation_time",
				Description: "The Unix timestamp for when the scaling policy was created.  This member is required.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the scaling policy.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PolicyARN"),
			},
			{
				Name:        "name",
				Description: "The name of the scaling policy.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PolicyName"),
			},
			{
				Name:        "type",
				Description: "The scaling policy type.  This member is required.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("PolicyType"),
			},
			{
				Name:        "resource_id",
				Description: "The identifier of the resource associated with the scaling policy",
				Type:        schema.TypeString,
			},
			{
				Name:        "scalable_dimension",
				Description: "The scalable dimension",
				Type:        schema.TypeString,
			},
			{
				Name:        "service_namespace",
				Description: "The namespace of the Amazon Web Services service that provides the resource, or a custom-resource.  This member is required.",
				Type:        schema.TypeString,
			},
			{
				Name:        "alarms",
				Description: "The CloudWatch alarms associated with the scaling policy.",
				Type:        schema.TypeJSON,
				Resolver:    resolveApplicationautoscalingPolicyAlarms,
			},
			{
				Name:        "step_scaling_policy_configuration",
				Description: "A step scaling policy.",
				Type:        schema.TypeJSON,
				Resolver:    resolveApplicationautoscalingPolicyStepScalingPolicyConfiguration,
			},
			{
				Name:        "target_tracking_scaling_policy_configuration",
				Description: "A target tracking scaling policy.",
				Type:        schema.TypeJSON,
				Resolver:    resolveApplicationautoscalingPolicyTargetTrackingScalingPolicyConfiguration,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchApplicationautoscalingPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().ApplicationAutoscaling

	config := applicationautoscaling.DescribeScalingPoliciesInput{
		ServiceNamespace: types.ServiceNamespace(c.AutoscalingNamespace),
	}
	for {
		output, err := svc.DescribeScalingPolicies(ctx, &config, func(o *applicationautoscaling.Options) {
			o.Region = c.Region
		})
		if err != nil {
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
func resolveApplicationautoscalingPolicyAlarms(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.ScalingPolicy)
	if r.Alarms == nil {
		return nil
	}
	b, _ := json.Marshal(r.Alarms)
	return resource.Set(c.Name, b)
}
func resolveApplicationautoscalingPolicyStepScalingPolicyConfiguration(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.ScalingPolicy)
	if r.StepScalingPolicyConfiguration == nil {
		return nil
	}
	b, _ := json.Marshal(r.StepScalingPolicyConfiguration)
	return resource.Set(c.Name, b)
}
func resolveApplicationautoscalingPolicyTargetTrackingScalingPolicyConfiguration(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.ScalingPolicy)
	if r.TargetTrackingScalingPolicyConfiguration == nil {
		return nil
	}
	b, _ := json.Marshal(r.TargetTrackingScalingPolicyConfiguration)
	return resource.Set(c.Name, b)
}
