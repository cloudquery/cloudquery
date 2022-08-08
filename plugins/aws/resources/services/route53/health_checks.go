package route53

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

type Route53HealthCheckWrapper struct {
	types.HealthCheck
	Tags map[string]string
}

func Route53HealthChecks() *schema.Table {
	return &schema.Table{
		Name:          "aws_route53_health_checks",
		Description:   "A complex type that contains information about one health check that is associated with the current AWS account.",
		Resolver:      fetchRoute53HealthChecks,
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
				Name:        "cloud_watch_alarm_configuration_dimensions",
				Description: "the metric that the CloudWatch alarm is associated with, a complex type that contains information about the dimensions for the metric.",
				Type:        schema.TypeJSON,
				Resolver:    resolveRoute53healthCheckCloudWatchAlarmConfigurationDimensions,
			},
			{
				Name:        "tags",
				Description: "The tags associated with the health check.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "caller_reference",
				Description: "A unique string that you specified when you created the health check.",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "The type of health check that you want to create, which indicates how Amazon Route 53 determines whether an endpoint is healthy.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("HealthCheckConfig.Type"),
			},
			{
				Name:        "alarm_identifier_name",
				Description: "The name of the CloudWatch alarm that you want Amazon Route 53 health checkers to use to determine whether this health check is healthy.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("HealthCheckConfig.AlarmIdentifier.Name"),
			},
			{
				Name:        "alarm_identifier_region",
				Description: "For the CloudWatch alarm that you want Route 53 health checkers to use to determine whether this health check is healthy, the region that the alarm was created in.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("HealthCheckConfig.AlarmIdentifier.Region"),
			},
			{
				Name:        "child_health_checks",
				Description: "(CALCULATED Health Checks Only) A complex type that contains one ChildHealthCheck element for each health check that you want to associate with a CALCULATED health check.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("HealthCheckConfig.ChildHealthChecks"),
			},
			{
				Name:        "disabled",
				Description: "Stops Route 53 from performing health checks.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("HealthCheckConfig.Disabled"),
			},
			{
				Name:        "enable_sni",
				Description: "Specify whether you want Amazon Route 53 to send the value of FullyQualifiedDomainName to the endpoint in the client_hello message during TLS negotiation.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("HealthCheckConfig.EnableSNI"),
			},
			{
				Name:        "failure_threshold",
				Description: "The number of consecutive health checks that an endpoint must pass or fail for Amazon Route 53 to change the current status of the endpoint from unhealthy to healthy or vice versa.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("HealthCheckConfig.FailureThreshold"),
			},
			{
				Name:        "fully_qualified_domain_name",
				Description: "Amazon Route 53 behavior depends on whether you specify a value for IPAddress.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("HealthCheckConfig.FullyQualifiedDomainName"),
			},
			{
				Name:        "health_threshold",
				Description: "The number of child health checks that are associated with a CALCULATED health check that Amazon Route 53 must consider healthy for the CALCULATED health check to be considered healthy.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("HealthCheckConfig.HealthThreshold"),
			},
			{
				Name:        "ip_address",
				Description: "The IPv4 or IPv6 IP address of the endpoint that you want Amazon Route 53 to perform health checks on.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("HealthCheckConfig.IPAddress"),
			},
			{
				Name:        "insufficient_data_health_status",
				Description: "When CloudWatch has insufficient data about the metric to determine the alarm state, the status that you want Amazon Route 53 to assign to the health check.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("HealthCheckConfig.InsufficientDataHealthStatus"),
			},
			{
				Name:        "inverted",
				Description: "Specify whether you want Amazon Route 53 to invert the status of a health check, for example, to consider a health check unhealthy when it otherwise would be considered healthy.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("HealthCheckConfig.Inverted"),
			},
			{
				Name:        "measure_latency",
				Description: "Specify whether you want Amazon Route 53 to measure the latency between health checkers in multiple AWS regions and your endpoint, and to display CloudWatch latency graphs on the Health Checks page in the Route 53 console.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("HealthCheckConfig.MeasureLatency"),
			},
			{
				Name:        "port",
				Description: "The port on the endpoint that you want Amazon Route 53 to perform health checks on.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("HealthCheckConfig.Port"),
			},
			{
				Name:        "regions",
				Description: "A complex type that contains one Region element for each region from which you want Amazon Route 53 health checkers to check the specified endpoint.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("HealthCheckConfig.Regions"),
			},
			{
				Name:        "request_interval",
				Description: "The number of seconds between the time that Amazon Route 53 gets a response from your endpoint and the time that it sends the next health check request.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("HealthCheckConfig.RequestInterval"),
			},
			{
				Name:        "resource_path",
				Description: "The path, if any, that you want Amazon Route 53 to request when performing health checks.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("HealthCheckConfig.ResourcePath"),
			},
			{
				Name:        "search_string",
				Description: "If the value of Type is HTTP_STR_MATCH or HTTPS_STR_MATCH, the string that you want Amazon Route 53 to search for in the response body from the specified resource.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("HealthCheckConfig.SearchString"),
			},
			{
				Name:        "health_check_version",
				Description: "The version of the health check.",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "id",
				Description: "The identifier that Amazon Route 53 assigned to the health check when you created it.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Id"),
			},
			{
				Name:        "cloud_watch_alarm_config_comparison_operator",
				Description: "For the metric that the CloudWatch alarm is associated with, the arithmetic operation that is used for the comparison.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("CloudWatchAlarmConfiguration.ComparisonOperator"),
			},
			{
				Name:        "cloud_watch_alarm_config_evaluation_periods",
				Description: "For the metric that the CloudWatch alarm is associated with, the number of periods that the metric is compared to the threshold.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("CloudWatchAlarmConfiguration.EvaluationPeriods"),
			},
			{
				Name:        "cloud_watch_alarm_config_metric_name",
				Description: "The name of the CloudWatch metric that the alarm is associated with.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("CloudWatchAlarmConfiguration.MetricName"),
			},
			{
				Name:        "cloud_watch_alarm_config_namespace",
				Description: "The namespace of the metric that the alarm is associated with.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("CloudWatchAlarmConfiguration.Namespace"),
			},
			{
				Name:        "cloud_watch_alarm_config_period",
				Description: "For the metric that the CloudWatch alarm is associated with, the duration of one evaluation period in seconds.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("CloudWatchAlarmConfiguration.Period"),
			},
			{
				Name:        "cloud_watch_alarm_config_statistic",
				Description: "For the metric that the CloudWatch alarm is associated with, the statistic that is applied to the metric.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("CloudWatchAlarmConfiguration.Statistic"),
			},
			{
				Name:        "cloud_watch_alarm_config_threshold",
				Description: "For the metric that the CloudWatch alarm is associated with, the value the metric is compared with.",
				Type:        schema.TypeFloat,
				Resolver:    schema.PathResolver("CloudWatchAlarmConfiguration.Threshold"),
			},
			{
				Name:        "linked_service_description",
				Description: "If the health check or hosted zone was created by another service, an optional description that can be provided by the other service.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LinkedService.Description"),
			},
			{
				Name:        "linked_service_service_principal",
				Description: "If the health check or hosted zone was created by another service, the service that created the resource.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LinkedService.ServicePrincipal"),
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the resource.",
				Type:        schema.TypeString,
				Resolver: client.ResolveARNGlobal(client.Route53Service, func(resource *schema.Resource) ([]string, error) {
					return []string{"healthcheck", *resource.Item.(Route53HealthCheckWrapper).Id}, nil
				}),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchRoute53HealthChecks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config route53.ListHealthChecksInput
	c := meta.(*client.Client)
	svc := c.Services().Route53

	processHealthChecksBundle := func(healthChecks []types.HealthCheck) error {
		tagsCfg := &route53.ListTagsForResourcesInput{ResourceType: types.TagResourceTypeHealthcheck, ResourceIds: make([]string, 0, len(healthChecks))}
		for _, h := range healthChecks {
			tagsCfg.ResourceIds = append(tagsCfg.ResourceIds, *h.Id)
		}
		tagsResponse, err := svc.ListTagsForResources(ctx, tagsCfg)
		if err != nil {
			return diag.WrapError(err)
		}
		for _, h := range healthChecks {
			wrapper := Route53HealthCheckWrapper{
				HealthCheck: h,
				Tags:        client.TagsToMap(getRoute53tagsByResourceID(*h.Id, tagsResponse.ResourceTagSets)),
			}
			res <- wrapper
		}
		return nil
	}

	for {
		response, err := svc.ListHealthChecks(ctx, &config)
		if err != nil {
			return diag.WrapError(err)
		}

		for i := 0; i < len(response.HealthChecks); i += 10 {
			end := i + 10

			if end > len(response.HealthChecks) {
				end = len(response.HealthChecks)
			}
			zones := response.HealthChecks[i:end]
			err := processHealthChecksBundle(zones)
			if err != nil {
				return diag.WrapError(err)
			}
		}

		if aws.ToString(response.Marker) == "" {
			break
		}
		config.Marker = response.Marker
	}
	return nil
}
func resolveRoute53healthCheckCloudWatchAlarmConfigurationDimensions(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(Route53HealthCheckWrapper)

	if r.CloudWatchAlarmConfiguration == nil {
		return nil
	}
	tags := map[string]*string{}
	for _, t := range r.CloudWatchAlarmConfiguration.Dimensions {
		tags[*t.Name] = t.Value
	}
	return diag.WrapError(resource.Set(c.Name, tags))
}
