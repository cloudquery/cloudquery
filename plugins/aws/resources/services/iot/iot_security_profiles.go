package iot

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func IotSecurityProfiles() *schema.Table {
	return &schema.Table{
		Name:         "aws_iot_security_profiles",
		Resolver:     fetchIotSecurityProfiles,
		Multiplex:    client.ServiceAccountRegionMultiplexer("iot"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
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
				Name:        "targets",
				Description: "Targets associated with the security profile",
				Type:        schema.TypeStringArray,
				Resolver:    ResolveIotSecurityProfileTargets,
			},
			{
				Name:        "tags",
				Description: "Tags of the resource",
				Type:        schema.TypeJSON,
				Resolver:    ResolveIotSecurityProfileTags,
			},
			{
				Name:        "additional_metrics_to_retain",
				Description: "Please use DescribeSecurityProfileResponse$additionalMetricsToRetainV2 instead. A list of metrics whose data is retained (stored)",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "additional_metrics_to_retain_v2",
				Description: "A list of metrics whose data is retained (stored)",
				Type:        schema.TypeJSON,
				Resolver:    resolveIotSecurityProfilesAdditionalMetricsToRetainV2,
			},
			{
				Name:        "alert_targets",
				Description: "Where the alerts are sent",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "creation_date",
				Description: "The time the security profile was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "last_modified_date",
				Description: "The time the security profile was last modified.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "arn",
				Description: "The ARN of the security profile.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SecurityProfileArn"),
			},
			{
				Name:        "description",
				Description: "A description of the security profile (associated with the security profile when it was created or updated).",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SecurityProfileDescription"),
			},
			{
				Name:        "name",
				Description: "The name of the security profile.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SecurityProfileName"),
			},
			{
				Name:        "version",
				Description: "The version of the security profile",
				Type:        schema.TypeBigInt,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_iot_security_profile_behaviors",
				Description: "A Device Defender security profile behavior.",
				Resolver:    fetchIotSecurityProfileBehaviors,
				Columns: []schema.Column{
					{
						Name:        "security_profile_cq_id",
						Description: "Unique CloudQuery ID of aws_iot_security_profiles table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "name",
						Description: "The name you've given to the behavior. ",
						Type:        schema.TypeString,
					},
					{
						Name:        "criteria_comparison_operator",
						Description: "The operator that relates the thing measured (metric) to the criteria (containing a value or statisticalThreshold)",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Criteria.ComparisonOperator"),
					},
					{
						Name:        "criteria_consecutive_datapoints_to_alarm",
						Description: "If a device is in violation of the behavior for the specified number of consecutive datapoints, an alarm occurs",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("Criteria.ConsecutiveDatapointsToAlarm"),
					},
					{
						Name:        "criteria_consecutive_datapoints_to_clear",
						Description: "If an alarm has occurred and the offending device is no longer in violation of the behavior for the specified number of consecutive datapoints, the alarm is cleared",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("Criteria.ConsecutiveDatapointsToClear"),
					},
					{
						Name:        "criteria_duration_seconds",
						Description: "Use this to specify the time duration over which the behavior is evaluated, for those criteria that have a time dimension (for example, NUM_MESSAGES_SENT)",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("Criteria.DurationSeconds"),
					},
					{
						Name:        "criteria_ml_detection_config_confidence_level",
						Description: "The sensitivity of anomalous behavior evaluation",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Criteria.MlDetectionConfig.ConfidenceLevel"),
					},
					{
						Name:        "criteria_statistical_threshold_statistic",
						Description: "The percentile that resolves to a threshold value by which compliance with a behavior is determined",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Criteria.StatisticalThreshold.Statistic"),
					},
					{
						Name:        "criteria_value",
						Description: "The value to be compared with the metric.",
						Type:        schema.TypeJSON,
						Resolver:    resolveIotSecurityProfileBehaviorsCriteriaValue,
					},
					{
						Name:        "metric",
						Description: "What is measured by the behavior.",
						Type:        schema.TypeString,
					},
					{
						Name:        "metric_dimension_dimension_name",
						Description: "A unique identifier for the dimension. ",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("MetricDimension.DimensionName"),
					},
					{
						Name:        "metric_dimension_operator",
						Description: "Defines how the dimensionValues of a dimension are interpreted",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("MetricDimension.Operator"),
					},
					{
						Name:        "suppress_alerts",
						Description: "Suppresses alerts.",
						Type:        schema.TypeBool,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchIotSecurityProfiles(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().IOT
	input := iot.ListSecurityProfilesInput{
		MaxResults: aws.Int32(250),
	}

	for {
		response, err := svc.ListSecurityProfiles(ctx, &input, func(options *iot.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}

		for _, s := range response.SecurityProfileIdentifiers {
			profile, err := svc.DescribeSecurityProfile(ctx, &iot.DescribeSecurityProfileInput{
				SecurityProfileName: s.Name,
			}, func(options *iot.Options) {
				options.Region = cl.Region
			})
			if err != nil {
				return diag.WrapError(err)
			}
			res <- profile
		}

		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}
func ResolveIotSecurityProfileTargets(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i := resource.Item.(*iot.DescribeSecurityProfileOutput)
	cl := meta.(*client.Client)
	svc := cl.Services().IOT
	input := iot.ListTargetsForSecurityProfileInput{
		SecurityProfileName: i.SecurityProfileName,
		MaxResults:          aws.Int32(250),
	}

	var targets []string
	for {
		response, err := svc.ListTargetsForSecurityProfile(ctx, &input, func(options *iot.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}

		for _, t := range response.SecurityProfileTargets {
			targets = append(targets, *t.Arn)
		}

		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return resource.Set(c.Name, targets)
}
func ResolveIotSecurityProfileTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i := resource.Item.(*iot.DescribeSecurityProfileOutput)
	cl := meta.(*client.Client)
	svc := cl.Services().IOT
	input := iot.ListTagsForResourceInput{
		ResourceArn: i.SecurityProfileArn,
	}
	tags := make(map[string]string)

	for {
		response, err := svc.ListTagsForResource(ctx, &input, func(options *iot.Options) {
			options.Region = cl.Region
		})

		if err != nil {
			return diag.WrapError(err)
		}

		client.TagsIntoMap(response.Tags, tags)

		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return diag.WrapError(resource.Set(c.Name, tags))
}
func resolveIotSecurityProfilesAdditionalMetricsToRetainV2(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i := resource.Item.(*iot.DescribeSecurityProfileOutput)

	if i.AdditionalMetricsToRetainV2 == nil {
		return nil
	}

	b, err := json.Marshal(i.AdditionalMetricsToRetainV2)
	if err != nil {
		return diag.WrapError(err)
	}
	return resource.Set(c.Name, b)
}
func fetchIotSecurityProfileBehaviors(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	i := parent.Item.(*iot.DescribeSecurityProfileOutput)
	if i.Behaviors == nil {
		return nil
	}

	res <- i.Behaviors
	return nil
}
func resolveIotSecurityProfileBehaviorsCriteriaValue(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i := resource.Item.(types.Behavior)
	if i.Criteria == nil || i.Criteria.Value == nil {
		return nil
	}

	data, err := json.Marshal(i.Criteria.Value)
	if err != nil {
		return diag.WrapError(err)
	}

	return resource.Set(c.Name, data)
}
