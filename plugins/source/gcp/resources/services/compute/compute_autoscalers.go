package compute

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
	"google.golang.org/api/compute/v1"
)

func ComputeAutoscalers() *schema.Table {
	return &schema.Table{
		Name:        "gcp_compute_autoscalers",
		Description: "Represents an Autoscaler resource.",
		Resolver:    fetchComputeAutoscalers,

		Multiplex: client.ProjectMultiplex,

		Options: schema.TableCreationOptions{PrimaryKeys: []string{"project_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "project_id",
				Description: "GCP Project Id of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveProject,
			},
			{
				Name:        "cool_down_period_sec",
				Description: "The number of seconds that the autoscaler waits before it starts collecting information from a new instance This prevents the autoscaler from collecting information when the instance is initializing, during which the collected usage would not be reliable The default time autoscaler waits is 60 seconds  Virtual machine initialization times might vary because of numerous factors We recommend that you test how long an instance may take to initialize To do this, create an instance and time the startup process",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("AutoscalingPolicy.CoolDownPeriodSec"),
			},
			{
				Name:        "cpu_utilization_predictive_method",
				Description: "Indicates whether predictive autoscaling based on CPU metric is enabled Valid values are:  * NONE (default) No predictive method is used The autoscaler scales the group to meet current demand based on real-time metrics * OPTIMIZE_AVAILABILITY Predictive autoscaling improves availability by monitoring daily and weekly load patterns and scaling out ahead of anticipated demand",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AutoscalingPolicy.CpuUtilization.PredictiveMethod"),
			},
			{
				Name:        "cpu_utilization_utilization_target",
				Description: "The target CPU utilization that the autoscaler maintains Must be a float value in the range (0, 1] If not specified, the default is 06  If the CPU level is below the target utilization, the autoscaler scales in the number of instances until it reaches the minimum number of instances you specified or until the average CPU of your instances reaches the target utilization  If the average CPU is above the target utilization, the autoscaler scales out until it reaches the maximum number of instances you specified or until the average utilization reaches the target utilization",
				Type:        schema.TypeFloat,
				Resolver:    schema.PathResolver("AutoscalingPolicy.CpuUtilization.UtilizationTarget"),
			},
			{
				Name:        "load_balancing_utilization_utilization_target",
				Description: "Fraction of backend capacity utilization (set in HTTP(S) load balancing configuration) that the autoscaler maintains Must be a positive float value If not defined, the default is 08",
				Type:        schema.TypeFloat,
				Resolver:    schema.PathResolver("AutoscalingPolicy.LoadBalancingUtilization.UtilizationTarget"),
			},
			{
				Name:        "max_num_replicas",
				Description: "The maximum number of instances that the autoscaler can scale out to This is required when creating or updating an autoscaler The maximum number of replicas must not be lower than minimal number of replicas",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("AutoscalingPolicy.MaxNumReplicas"),
			},
			{
				Name:        "min_num_replicas",
				Description: "The minimum number of replicas that the autoscaler can scale in to This cannot be less than 0 If not provided, autoscaler chooses a default value depending on maximum number of instances allowed",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("AutoscalingPolicy.MinNumReplicas"),
			},
			{
				Name:        "mode",
				Description: "Defines operating mode for this policy",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AutoscalingPolicy.Mode"),
			},
			{
				Name:        "scale_in_control_max_scaled_in_replicas_calculated",
				Description: "Absolute value of VM instances calculated based on the specific mode ",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("AutoscalingPolicy.ScaleInControl.MaxScaledInReplicas.Calculated"),
			},
			{
				Name:        "scale_in_control_max_scaled_in_replicas_fixed",
				Description: "Specifies a fixed number of VM instances This must be a positive integer",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("AutoscalingPolicy.ScaleInControl.MaxScaledInReplicas.Fixed"),
			},
			{
				Name:        "scale_in_control_max_scaled_in_replicas_percent",
				Description: "Specifies a percentage of instances between 0 to 100%, inclusive For example, specify 80 for 80%",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("AutoscalingPolicy.ScaleInControl.MaxScaledInReplicas.Percent"),
			},
			{
				Name:        "scale_in_control_time_window_sec",
				Description: "How far back autoscaling looks when computing recommendations to include directives regarding slower scale in, as described above",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("AutoscalingPolicy.ScaleInControl.TimeWindowSec"),
			},
			{
				Name:          "scaling_schedules",
				Description:   "Scaling schedules defined for an autoscaler Multiple schedules can be set on an autoscaler, and they can overlap During overlapping periods the greatest min_required_replicas of all scaling schedules is applied Up to 128 scaling schedules are allowed",
				Type:          schema.TypeJSON,
				IgnoreInTests: true,
				Resolver:      schema.PathResolver("AutoscalingPolicy.ScalingSchedules"),
			},
			{
				Name:        "creation_timestamp",
				Description: "Creation timestamp in RFC3339 text format",
				Type:        schema.TypeString,
			},
			{
				Name:        "description",
				Description: "An optional description of this resource Provide this property when you create the resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "id",
				Description: "The unique identifier for the resource This identifier is defined by the server",
				Type:        schema.TypeString,
				Resolver:    client.ResolveResourceId,
			},
			{
				Name:        "kind",
				Description: "Type of the resource Always compute#autoscaler for autoscalers",
				Type:        schema.TypeString,
			},
			{
				Name:        "name",
				Description: "Name of the resource Provided by the client when the resource is created The name must be 1-63 characters long, and comply with RFC1035 Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash",
				Type:        schema.TypeString,
			},
			{
				Name:        "recommended_size",
				Description: "Target recommended MIG size (number of instances) computed by autoscaler Autoscaler calculates the recommended MIG size even when the autoscaling policy mode is different from ON This field is empty when autoscaler is not connected to an existing managed instance group or autoscaler did not generate its prediction",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "region",
				Description: "URL of the region where the instance group resides (for autoscalers living in regional scope)",
				Type:        schema.TypeString,
			},
			{
				Name:          "scaling_schedule_status",
				Description:   "Status information of existing scaling schedules",
				Type:          schema.TypeJSON,
				IgnoreInTests: true,
			},
			{
				Name:        "self_link",
				Description: "Server-defined URL for the resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "status",
				Description: "The status of the autoscaler configuration Current set of possible values: - PENDING: Autoscaler backend hasn't read new/updated configuration  - DELETING: Configuration is being deleted - ACTIVE: Configuration is acknowledged to be effective Some warnings might be present in the statusDetails field - ERROR: Configuration has errors Actionable for users Details are present in the statusDetails field  New values might be added in the future",
				Type:        schema.TypeString,
			},
			{
				Name:        "status_details",
				Description: "Human-readable details about the current state of the autoscaler Read the documentation for Commonly returned status messages for examples of status messages you might encounter",
				Type:        schema.TypeJSON,
				Resolver:    resolveComputeAutoscalerStatusDetails,
			},
			{
				Name:        "target",
				Description: "URL of the managed instance group that this autoscaler will scale This field is required when creating an autoscaler",
				Type:        schema.TypeString,
			},
			{
				Name:        "zone",
				Description: "URL of the zone where the instance group resides (for autoscalers living in zonal scope)",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "gcp_compute_autoscaler_custom_metric_utilizations",
				Description: "Custom utilization metric policy",
				Resolver:    fetchComputeAutoscalerCustomMetricUtilizations,
				Columns: []schema.Column{
					{
						Name:        "autoscaler_cq_id",
						Description: "Unique ID of gcp_compute_autoscalers table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "autoscaler_id",
						Type:     schema.TypeString,
						Resolver: schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "filter",
						Description: "A filter string, compatible with a Stackdriver Monitoring filter string",
						Type:        schema.TypeString,
					},
					{
						Name:        "metric",
						Description: "The identifier (type) of the Stackdriver Monitoring metric The metric cannot have negative values  The metric must have a value type of INT64 or DOUBLE",
						Type:        schema.TypeString,
					},
					{
						Name:        "single_instance_assignment",
						Description: "per-group metric value that represents the total amount of work to be done or resource usage",
						Type:        schema.TypeFloat,
					},
					{
						Name:        "utilization_target",
						Description: "The target value of the metric that autoscaler maintains This must be a positive value A utilization metric scales number of virtual machines handling requests to increase or decrease proportionally to the metric  For example, a good metric to use as a utilization_target is https://wwwgoogleapiscom/compute/v1/instance/network/received_bytes_count The autoscaler works to keep this value constant for each of the instances",
						Type:        schema.TypeFloat,
					},
					{
						Name:        "utilization_target_type",
						Description: "Defines how target utilization value is expressed for a Stackdriver Monitoring metric Either GAUGE, DELTA_PER_SECOND, or DELTA_PER_MINUTE",
						Type:        schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchComputeAutoscalers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	nextPageToken := ""
	c := meta.(*client.Client)
	for {
		output, err := c.Services.Compute.Autoscalers.AggregatedList(c.ProjectId).PageToken(nextPageToken).Do()
		if err != nil {
			return errors.WithStack(err)
		}

		var autoscalers []*compute.Autoscaler
		for _, items := range output.Items {
			autoscalers = append(autoscalers, items.Autoscalers...)
		}
		res <- autoscalers

		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
func resolveComputeAutoscalerStatusDetails(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	autoscaler := resource.Item.(*compute.Autoscaler)
	res := map[string]string{}
	for _, v := range autoscaler.StatusDetails {
		res[v.Type] = v.Message
	}
	return errors.WithStack(resource.Set("status_details", res))
}
func fetchComputeAutoscalerCustomMetricUtilizations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	autoscaler := parent.Item.(*compute.Autoscaler)
	if autoscaler.AutoscalingPolicy != nil {
		res <- autoscaler.AutoscalingPolicy.CustomMetricUtilizations
	}
	return nil
}
