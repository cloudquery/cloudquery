package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"google.golang.org/api/compute/v1"
)

func ComputeAutoscalers() *schema.Table {
	return &schema.Table{
		Name:         "gcp_compute_autoscalers",
		Resolver:     fetchComputeAutoscalers,
		Multiplex:    client.ProjectMultiplex,
		DeleteFilter: client.DeleteProjectFilter,
		IgnoreError:  client.IgnoreErrorHandler,

		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name:     "cool_down_period_sec",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("AutoscalingPolicy.CoolDownPeriodSec"),
			},
			{
				Name:     "cpu_utilization_predictive_method",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AutoscalingPolicy.CpuUtilization.PredictiveMethod"),
			},
			{
				Name:     "cpu_utilization_utilization_target",
				Type:     schema.TypeFloat,
				Resolver: schema.PathResolver("AutoscalingPolicy.CpuUtilization.UtilizationTarget"),
			},
			{
				Name:     "load_balancing_utilization_utilization_target",
				Type:     schema.TypeFloat,
				Resolver: schema.PathResolver("AutoscalingPolicy.LoadBalancingUtilization.UtilizationTarget"),
			},
			{
				Name:     "max_num_replicas",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("AutoscalingPolicy.MaxNumReplicas"),
			},
			{
				Name:     "min_num_replicas",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("AutoscalingPolicy.MinNumReplicas"),
			},
			{
				Name:     "mode",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AutoscalingPolicy.Mode"),
			},
			{
				Name:     "scale_in_control_max_scaled_in_replicas_calculated ",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("AutoscalingPolicy.ScaleInControl.MaxScaledInReplicas.Calculated"),
			},
			{
				Name:     "scale_in_control_max_scaled_in_replicas_fixed",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("AutoscalingPolicy.ScaleInControl.MaxScaledInReplicas.Fixed"),
			},
			{
				Name:     "scale_in_control_max_scaled_in_replicas_percent",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("AutoscalingPolicy.ScaleInControl.MaxScaledInReplicas.Percent"),
			},
			{
				Name:     "scale_in_control_time_window_sec",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("AutoscalingPolicy.ScaleInControl.TimeWindowSec"),
			},
			{
				Name:     "scaling_schedules",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AutoscalingPolicy.ScalingSchedules"),
			},
			{
				Name: "creation_timestamp",
				Type: schema.TypeString,
			},
			{
				Name: "description",
				Type: schema.TypeString,
			},
			{
				Name:     "resource_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveResourceId,
			},
			{
				Name: "kind",
				Type: schema.TypeString,
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name: "recommended_size",
				Type: schema.TypeBigInt,
			},
			{
				Name: "region",
				Type: schema.TypeString,
			},
			{
				Name: "scaling_schedule_status",
				Type: schema.TypeJSON,
			},
			{
				Name: "self_link",
				Type: schema.TypeString,
			},
			{
				Name: "status",
				Type: schema.TypeString,
			},
			{
				Name:     "status_details",
				Type:     schema.TypeJSON,
				Resolver: resolveComputeAutoscalerStatusDetails,
			},
			{
				Name: "target",
				Type: schema.TypeString,
			},
			{
				Name: "zone",
				Type: schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "gcp_compute_autoscaler_custom_metric_utilizations",
				Resolver: fetchComputeAutoscalerCustomMetricUtilizations,
				Columns: []schema.Column{
					{
						Name:     "autoscaler_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "filter",
						Type: schema.TypeString,
					},
					{
						Name: "metric",
						Type: schema.TypeString,
					},
					{
						Name: "single_instance_assignment",
						Type: schema.TypeFloat,
					},
					{
						Name: "utilization_target",
						Type: schema.TypeFloat,
					},
					{
						Name: "utilization_target_type",
						Type: schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchComputeAutoscalers(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	nextPageToken := ""
	c := meta.(*client.Client)
	for {
		call := c.Services.Compute.Autoscalers.AggregatedList(c.ProjectId).Context(ctx)
		call.PageToken(nextPageToken)
		output, err := call.Do()
		if err != nil {
			return err
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
func resolveComputeAutoscalerStatusDetails(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	autoscaler := resource.Item.(*compute.Autoscaler)
	res := map[string]string{}
	for _, v := range autoscaler.StatusDetails {
		res[v.Type] = v.Message
	}
	resource.Set("status_details", res)
	return nil
}
func fetchComputeAutoscalerCustomMetricUtilizations(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	autoscaler := parent.Item.(*compute.Autoscaler)
	if autoscaler.AutoscalingPolicy != nil {
		res <- autoscaler.AutoscalingPolicy.CustomMetricUtilizations
	}
	return nil
}
