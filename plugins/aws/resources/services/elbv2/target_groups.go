package elbv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	elbv2 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Elbv2TargetGroups() *schema.Table {
	return &schema.Table{
		Name:          "aws_elbv2_target_groups",
		Description:   "Information about a target group.",
		Resolver:      fetchElbv2TargetGroups,
		Multiplex:     client.ServiceAccountRegionMultiplexer("elasticloadbalancing"),
		IgnoreError:   client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter:  client.DeleteAccountRegionFilter,
		Options:       schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		IgnoreInTests: true,
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
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveElbv2targetGroupTags,
			},
			{
				Name:        "health_check_enabled",
				Description: "Indicates whether health checks are enabled.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "health_check_interval_seconds",
				Description: "The approximate amount of time, in seconds, between health checks of an individual target.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "health_check_path",
				Description: "The destination for health checks on the targets.",
				Type:        schema.TypeString,
			},
			{
				Name:        "health_check_port",
				Description: "The port to use to connect with the target.",
				Type:        schema.TypeString,
			},
			{
				Name:        "health_check_protocol",
				Description: "The protocol to use to connect with the target",
				Type:        schema.TypeString,
			},
			{
				Name:        "health_check_timeout_seconds",
				Description: "The amount of time, in seconds, during which no response means a failed health check.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "healthy_threshold_count",
				Description: "The number of consecutive health checks successes required before considering an unhealthy target healthy.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "load_balancer_arns",
				Description: "The Amazon Resource Names (ARN) of the load balancers that route traffic to this target group.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "matcher_grpc_code",
				Description: "You can specify values between 0 and 99",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Matcher.GrpcCode"),
			},
			{
				Name:        "matcher_http_code",
				Description: "For Application Load Balancers, you can specify values between 200 and 499, and the default value is 200",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Matcher.HttpCode"),
			},
			{
				Name:        "port",
				Description: "The port on which the targets are listening",
				Type:        schema.TypeInt,
			},
			{
				Name:        "protocol",
				Description: "The protocol to use for routing traffic to the targets.",
				Type:        schema.TypeString,
			},
			{
				Name:        "protocol_version",
				Description: "[HTTP/HTTPS protocol] The protocol version",
				Type:        schema.TypeString,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the target group.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TargetGroupArn"),
			},
			{
				Name:        "name",
				Description: "The name of the target group.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TargetGroupName"),
			},
			{
				Name:        "target_type",
				Description: "The type of target that you must specify when registering targets with this target group",
				Type:        schema.TypeString,
			},
			{
				Name:        "unhealthy_threshold_count",
				Description: "The number of consecutive health check failures required before considering the target unhealthy.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "vpc_id",
				Description: "The ID of the VPC for the targets.",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:          "aws_elbv2_target_group_target_health_descriptions",
				Description:   "Information about the health of a target.",
				Resolver:      resolveElbv2TargetGroupTargetHealthDescriptions,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "target_group_cq_id",
						Description: "Unique CloudQuery ID of aws_elbv2_target_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "health_check_port",
						Description: "The port to use to connect with the target.",
						Type:        schema.TypeString,
					},
					{
						Name:        "target_id",
						Description: "The ID of the target.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Target.Id"),
					},
					{
						Name:        "target_availability_zone",
						Description: "An Availability Zone or all.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Target.AvailabilityZone"),
					},
					{
						Name:        "target_port",
						Description: "The port on which the target is listening.",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("Target.Port"),
					},
					{
						Name:        "target_health_description",
						Description: "A description of the target health that provides additional details.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("TargetHealth.Description"),
					},
					{
						Name:        "target_health_reason",
						Description: "The reason code. If the target state is healthy, a reason code is not provided.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("TargetHealth.Reason"),
					},
					{
						Name:        "target_health_state",
						Description: "The state of the target.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("TargetHealth.State"),
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchElbv2TargetGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config elbv2.DescribeTargetGroupsInput
	c := meta.(*client.Client)
	svc := c.Services().ELBv2
	for {
		response, err := svc.DescribeTargetGroups(ctx, &config, func(options *elbv2.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- response.TargetGroups
		if aws.ToString(response.NextMarker) == "" {
			break
		}
		config.Marker = response.NextMarker
	}
	return nil
}
func resolveElbv2targetGroupTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	region := meta.(*client.Client).Region
	svc := meta.(*client.Client).Services().ELBv2
	targetGroup := resource.Item.(types.TargetGroup)
	tagsOutput, err := svc.DescribeTags(ctx, &elbv2.DescribeTagsInput{
		ResourceArns: []string{
			*targetGroup.TargetGroupArn,
		},
	}, func(o *elbv2.Options) {
		o.Region = region
	})
	if err != nil {
		return diag.WrapError(err)
	}
	if len(tagsOutput.TagDescriptions) == 0 {
		return nil
	}
	tags := make(map[string]*string)
	for _, s := range tagsOutput.TagDescriptions[0].Tags {
		tags[*s.Key] = s.Value
	}
	return diag.WrapError(resource.Set(c.Name, tags))
}

func resolveElbv2TargetGroupTargetHealthDescriptions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().ELBv2
	tg := parent.Item.(types.TargetGroup)
	response, err := svc.DescribeTargetHealth(ctx, &elbv2.DescribeTargetHealthInput{
		TargetGroupArn: tg.TargetGroupArn,
	})
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return diag.WrapError(err)
	}
	res <- response.TargetHealthDescriptions
	return nil
}
