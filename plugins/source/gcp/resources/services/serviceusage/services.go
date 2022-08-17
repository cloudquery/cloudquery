package serviceusage

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
	serviceusage "google.golang.org/api/serviceusage/v1"
)

//go:generate cq-gen --resource services --config gen.hcl --output .
func Services() *schema.Table {
	return &schema.Table{
		Name:        "gcp_serviceusage_services",
		Description: "A service that is available for use by the consumer",
		Resolver:    fetchServiceusageServices,
		Multiplex:   client.ProjectMultiplex,

		Options: schema.TableCreationOptions{PrimaryKeys: []string{"name"}},
		Columns: []schema.Column{
			{
				Name:        "project_id",
				Description: "GCP Project Id of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveProject,
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Config.Name"),
			},
			{
				Name:        "authentication",
				Description: "Auth configuration",
				Type:        schema.TypeJSON,
				Resolver:    resolveServicesAuthentication,
			},
			{
				Name:        "documentation",
				Description: "Additional API documentation",
				Type:        schema.TypeJSON,
				Resolver:    resolveServicesDocumentation,
			},
			{
				Name:        "title",
				Description: "The product title for this service",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Config.Title"),
			},
			{
				Name:        "usage_producer_notification_channel",
				Description: "The full resource name of a channel used for sending notifications to the service producer",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Config.Usage.ProducerNotificationChannel"),
			},
			{
				Name:        "usage_requirements",
				Description: "Requirements that must be satisfied before a consumer project can use the service",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("Config.Usage.Requirements"),
			},
			{
				Name:        "parent",
				Description: "The resource name of the consumer",
				Type:        schema.TypeString,
			},
			{
				Name:        "state",
				Description: "\"STATE_UNSPECIFIED\" - The default value, which indicates that the enabled state of the service is unspecified or not meaningful Currently, all consumers other than projects (such as folders and organizations) are always in this state   \"DISABLED\" - The service cannot be used by this consumer",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:          "gcp_serviceusage_service_apis",
				Description:   "Api is a light-weight descriptor for an API Interface Interfaces are also described as \"protocol buffer services\" in some contexts, such as by the \"service\" keyword in a proto file, but they are different from API Services, which represent a concrete implementation of an interface as opposed to simply a description of methods and bindings",
				Resolver:      fetchServiceusageServiceApis,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "service_cq_id",
						Description: "Unique CloudQuery ID of gcp_serviceusage_services table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "methods",
						Description: "The methods of this interface, in unspecified order",
						Type:        schema.TypeJSON,
						Resolver:    resolveServiceApisMethods,
					},
					{
						Name:        "mixins",
						Description: "Included interfaces",
						Type:        schema.TypeJSON,
						Resolver:    resolveServiceApisMixins,
					},
					{
						Name:        "name",
						Description: "The fully qualified name of this interface, including package name followed by the interface's simple name",
						Type:        schema.TypeString,
					},
					{
						Name:        "options",
						Description: "Any metadata attached to the interface",
						Type:        schema.TypeJSON,
						Resolver:    resolveServiceApisOptions,
					},
					{
						Name:        "source_context_file_name",
						Description: "The path-qualified name of the proto file that contained the associated protobuf element",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SourceContext.FileName"),
					},
					{
						Name:        "syntax",
						Description: "\"SYNTAX_PROTO2\" - Syntax `proto2`   \"SYNTAX_PROTO3\" - Syntax `proto3`",
						Type:        schema.TypeString,
					},
					{
						Name:        "version",
						Description: "A version string for this interface",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:          "gcp_serviceusage_service_endpoints",
				Description:   "`Endpoint` describes a network address of a service that serves a set of APIs",
				Resolver:      fetchServiceusageServiceEndpoints,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "service_cq_id",
						Description: "Unique CloudQuery ID of gcp_serviceusage_services table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "allow_cors",
						Description: "Allowing CORS (https://enwikipediaorg/wiki/Cross-origin_resource_sharing), aka cross-domain traffic, would allow the backends served from this endpoint to receive and respond to HTTP OPTIONS requests",
						Type:        schema.TypeBool,
					},
					{
						Name:        "name",
						Description: "The canonical name of this endpoint",
						Type:        schema.TypeString,
					},
					{
						Name:        "target",
						Description: "The specification of an Internet routable address of API frontend that will handle requests to this API Endpoint (https://cloudgooglecom/apis/design/glossary)",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "gcp_serviceusage_service_monitored_resources",
				Description: "An object that describes the schema of a MonitoredResource object using a type name and a set of labels",
				Resolver:    fetchServiceusageServiceMonitoredResources,
				Columns: []schema.Column{
					{
						Name:        "service_cq_id",
						Description: "Unique CloudQuery ID of gcp_serviceusage_services table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "description",
						Description: "Optional",
						Type:        schema.TypeString,
					},
					{
						Name:        "display_name",
						Description: "Optional",
						Type:        schema.TypeString,
					},
					{
						Name:        "labels",
						Description: "Required",
						Type:        schema.TypeJSON,
						Resolver:    resolveServiceMonitoredResourcesLabels,
					},
					{
						Name:        "launch_stage",
						Description: "Optional",
						Type:        schema.TypeString,
					},
					{
						Name:        "name",
						Description: "Optional",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "Required",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:          "gcp_serviceusage_service_monitoring_consumer_destinations",
				Description:   "Configuration of a specific monitoring destination (the producer project or the consumer project)",
				Resolver:      fetchServiceusageServiceMonitoringConsumerDestinations,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "service_cq_id",
						Description: "Unique CloudQuery ID of gcp_serviceusage_services table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "metrics",
						Description: "Types of the metrics to report to this monitoring destination",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "monitored_resource",
						Description: "The monitored resource type",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:          "gcp_serviceusage_service_monitoring_producer_destinations",
				Description:   "Configuration of a specific monitoring destination (the producer project or the consumer project)",
				Resolver:      fetchServiceusageServiceMonitoringProducerDestinations,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "service_cq_id",
						Description: "Unique CloudQuery ID of gcp_serviceusage_services table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "metrics",
						Description: "Types of the metrics to report to this monitoring destination",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "monitored_resource",
						Description: "The monitored resource type",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:          "gcp_serviceusage_service_quota_limits",
				Description:   "`QuotaLimit` defines a specific limit that applies over a specified duration for a limit type",
				Resolver:      fetchServiceusageServiceQuotaLimits,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "service_cq_id",
						Description: "Unique CloudQuery ID of gcp_serviceusage_services table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "default_limit",
						Description: "Default number of tokens that can be consumed during the specified duration",
						Type:        schema.TypeInt,
						Resolver:    resolveServiceQuotaLimitsDefaultLimit,
					},
					{
						Name:        "description",
						Description: "Optional",
						Type:        schema.TypeString,
					},
					{
						Name:        "display_name",
						Description: "User-visible display name for this limit",
						Type:        schema.TypeString,
					},
					{
						Name:        "duration",
						Description: "Duration of this limit in textual notation",
						Type:        schema.TypeString,
					},
					{
						Name:        "free_tier",
						Description: "Free tier value displayed in the Developers Console for this limit",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "max_limit",
						Description: "Maximum number of tokens that can be consumed during the specified duration",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "metric",
						Description: "The name of the metric this quota limit applies to",
						Type:        schema.TypeString,
					},
					{
						Name:        "name",
						Description: "Name of the quota limit",
						Type:        schema.TypeString,
					},
					{
						Name:        "unit",
						Description: "Specify the unit of the quota limit",
						Type:        schema.TypeString,
					},
					{
						Name:        "values",
						Description: "Tiered limit values",
						Type:        schema.TypeJSON,
					},
				},
			},
			{
				Name:          "gcp_serviceusage_service_quota_metric_rules",
				Description:   "Bind API methods to metrics",
				Resolver:      fetchServiceusageServiceQuotaMetricRules,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "service_cq_id",
						Description: "Unique CloudQuery ID of gcp_serviceusage_services table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "metric_costs",
						Description: "Metrics to update when the selected methods are called, and the associated cost applied to each metric",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "selector",
						Description: "Selects the methods to which this rule applies",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:          "gcp_serviceusage_service_usage_rules",
				Description:   "Usage configuration rules for the service",
				Resolver:      fetchServiceusageServiceUsageRules,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "service_cq_id",
						Description: "Unique CloudQuery ID of gcp_serviceusage_services table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "allow_unregistered_calls",
						Description: "If true, the selected method allows unregistered calls, eg",
						Type:        schema.TypeBool,
					},
					{
						Name:        "selector",
						Description: "Selects the methods to which this rule applies",
						Type:        schema.TypeString,
					},
					{
						Name:        "skip_service_control",
						Description: "If true, the selected method should skip service control and the control plane features, such as quota and billing, will not be available",
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

func fetchServiceusageServices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		output, err := c.Services.ServiceUsage.Services.List(fmt.Sprintf("projects/%s", c.ProjectId)).PageToken(nextPageToken).Do()
		if err != nil {
			return errors.WithStack(err)
		}

		res <- output.Services

		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}

func resolveServicesAuthentication(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(*serviceusage.GoogleApiServiceusageV1Service)
	if p.Config == nil {
		return nil
	}
	j, err := json.Marshal(p.Config.Authentication)
	if err != nil {
		return errors.WithStack(err)
	}
	return errors.WithStack(resource.Set(c.Name, j))
}
func resolveServicesDocumentation(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(*serviceusage.GoogleApiServiceusageV1Service)
	if p.Config == nil {
		return nil
	}
	j, err := json.Marshal(p.Config.Documentation)
	if err != nil {
		return errors.WithStack(err)
	}
	return errors.WithStack(resource.Set(c.Name, j))
}
func fetchServiceusageServiceApis(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(*serviceusage.GoogleApiServiceusageV1Service)
	if p.Config == nil {
		return nil
	}
	res <- p.Config.Apis
	return nil
}
func resolveServiceApisMethods(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(*serviceusage.Api)
	j, err := json.Marshal(p.Methods)
	if err != nil {
		return errors.WithStack(err)
	}
	return errors.WithStack(resource.Set(c.Name, j))
}
func resolveServiceApisMixins(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(*serviceusage.Api)
	j, err := json.Marshal(p.Mixins)
	if err != nil {
		return errors.WithStack(err)
	}
	return errors.WithStack(resource.Set(c.Name, j))
}
func resolveServiceApisOptions(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(*serviceusage.Api)
	j, err := json.Marshal(p.Options)
	if err != nil {
		return errors.WithStack(err)
	}
	return errors.WithStack(resource.Set(c.Name, j))
}
func fetchServiceusageServiceEndpoints(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(*serviceusage.GoogleApiServiceusageV1Service)
	if p.Config == nil {
		return nil
	}
	res <- p.Config.Endpoints
	return nil
}
func fetchServiceusageServiceMonitoredResources(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(*serviceusage.GoogleApiServiceusageV1Service)
	if p.Config == nil {
		return nil
	}
	res <- p.Config.MonitoredResources
	return nil
}
func resolveServiceMonitoredResourcesLabels(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(*serviceusage.MonitoredResourceDescriptor)
	j, err := json.Marshal(p.Labels)
	if err != nil {
		return errors.WithStack(err)
	}
	return errors.WithStack(resource.Set(c.Name, j))
}
func fetchServiceusageServiceMonitoringConsumerDestinations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(*serviceusage.GoogleApiServiceusageV1Service)
	if p.Config == nil || p.Config.Monitoring == nil {
		return nil
	}
	res <- p.Config.Monitoring.ConsumerDestinations
	return nil
}
func fetchServiceusageServiceMonitoringProducerDestinations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(*serviceusage.GoogleApiServiceusageV1Service)
	if p.Config == nil || p.Config.Monitoring == nil {
		return nil
	}
	res <- p.Config.Monitoring.ProducerDestinations
	return nil
}
func fetchServiceusageServiceQuotaLimits(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(*serviceusage.GoogleApiServiceusageV1Service)
	if p.Config == nil || p.Config.Quota == nil {
		return nil
	}
	res <- p.Config.Quota.Limits
	return nil
}
func resolveServiceQuotaLimitsDefaultLimit(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(*serviceusage.QuotaLimit)
	return errors.WithStack(resource.Set(c.Name, int32(p.DefaultLimit)))
}
func fetchServiceusageServiceQuotaMetricRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(*serviceusage.GoogleApiServiceusageV1Service)
	if p.Config == nil || p.Config.Quota == nil {
		return nil
	}
	res <- p.Config.Quota.MetricRules
	return nil
}
func fetchServiceusageServiceUsageRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(*serviceusage.GoogleApiServiceusageV1Service)
	if p.Config == nil || p.Config.Usage == nil {
		return nil
	}
	res <- p.Config.Usage.Rules
	return nil
}
