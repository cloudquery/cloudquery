package resources

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Route53HealthChecks() *schema.Table {
	return &schema.Table{
		Name:         "aws_route53_health_checks",
		Resolver:     fetchRoute53HealthChecks,
		Multiplex:    client.AccountMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountFilter,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "cloud_watch_alarm_configuration_dimensions",
				Type:     schema.TypeJSON,
				Resolver: resolveRoute53healthCheckCloudWatchAlarmConfigurationDimensions,
			},
			{
				Name: "tags",
				Type: schema.TypeJSON,
			},
			{
				Name: "caller_reference",
				Type: schema.TypeString,
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HealthCheckConfig.Type"),
			},
			{
				Name:     "alarm_identifier_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HealthCheckConfig.AlarmIdentifier.Name"),
			},
			{
				Name:     "alarm_identifier_region",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HealthCheckConfig.AlarmIdentifier.Region"),
			},
			{
				Name:     "child_health_checks",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("HealthCheckConfig.ChildHealthChecks"),
			},
			{
				Name:     "disabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("HealthCheckConfig.Disabled"),
			},
			{
				Name:     "enable_sni",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("HealthCheckConfig.EnableSNI"),
			},
			{
				Name:     "failure_threshold",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("HealthCheckConfig.FailureThreshold"),
			},
			{
				Name:     "fully_qualified_domain_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HealthCheckConfig.FullyQualifiedDomainName"),
			},
			{
				Name:     "health_threshold",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("HealthCheckConfig.HealthThreshold"),
			},
			{
				Name:     "ip_address",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HealthCheckConfig.IPAddress"),
			},
			{
				Name:     "insufficient_data_health_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HealthCheckConfig.InsufficientDataHealthStatus"),
			},
			{
				Name:     "inverted",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("HealthCheckConfig.Inverted"),
			},
			{
				Name:     "measure_latency",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("HealthCheckConfig.MeasureLatency"),
			},
			{
				Name:     "port",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("HealthCheckConfig.Port"),
			},
			{
				Name:     "regions",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("HealthCheckConfig.Regions"),
			},
			{
				Name:     "request_interval",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("HealthCheckConfig.RequestInterval"),
			},
			{
				Name:     "resource_path",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HealthCheckConfig.ResourcePath"),
			},
			{
				Name:     "search_string",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HealthCheckConfig.SearchString"),
			},
			{
				Name: "health_check_version",
				Type: schema.TypeBigInt,
			},
			{
				Name:     "resource_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
			},
			{
				Name:     "cloud_watch_alarm_config_comparison_operator",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CloudWatchAlarmConfiguration.ComparisonOperator"),
			},
			{
				Name:     "cloud_watch_alarm_config_evaluation_periods",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("CloudWatchAlarmConfiguration.EvaluationPeriods"),
			},
			{
				Name:     "cloud_watch_alarm_config_metric_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CloudWatchAlarmConfiguration.MetricName"),
			},
			{
				Name:     "cloud_watch_alarm_config_namespace",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CloudWatchAlarmConfiguration.Namespace"),
			},
			{
				Name:     "cloud_watch_alarm_config_period",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("CloudWatchAlarmConfiguration.Period"),
			},
			{
				Name:     "cloud_watch_alarm_config_statistic",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CloudWatchAlarmConfiguration.Statistic"),
			},
			{
				Name:     "cloud_watch_alarm_config_threshold",
				Type:     schema.TypeFloat,
				Resolver: schema.PathResolver("CloudWatchAlarmConfiguration.Threshold"),
			},
			{
				Name:     "linked_service_description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LinkedService.Description"),
			},
			{
				Name:     "linked_service_service_principal",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LinkedService.ServicePrincipal"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchRoute53HealthChecks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
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
			return err
		}
		for _, h := range healthChecks {
			tags := getRoute53tagsByResourceID(*h.Id, tagsResponse.ResourceTagSets)
			if len(tags) == 0 {
				continue
			}
			wrapper := Route53HealthCheckWrapper{
				HealthCheck: h,
				Tags:        make(map[string]interface{}, len(tags)),
			}

			for _, t := range tags {
				wrapper.Tags[*t.Key] = t.Value
			}
			res <- wrapper
		}
		return nil
	}

	for {
		response, err := svc.ListHealthChecks(ctx, &config)
		if err != nil {
			return err
		}

		for i := 0; i < len(response.HealthChecks); i += 10 {
			end := i + 10

			if end > len(response.HealthChecks) {
				end = len(response.HealthChecks)
			}
			zones := response.HealthChecks[i:end]
			err := processHealthChecksBundle(zones)
			if err != nil {
				return err
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
	r, ok := resource.Item.(Route53HealthCheckWrapper)
	if !ok {
		return fmt.Errorf("not route53 healch check")
	}

	if r.CloudWatchAlarmConfiguration == nil {
		return nil
	}
	tags := map[string]*string{}
	for _, t := range r.CloudWatchAlarmConfiguration.Dimensions {
		tags[*t.Name] = t.Value
	}
	return resource.Set(c.Name, tags)
}
