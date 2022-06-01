package streamanalytics

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/streamanalytics/mgmt/2020-03-01/streamanalytics"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func StreamanalyticsJobs() *schema.Table {
	return &schema.Table{
		Name:         "azure_streamanalytics_jobs",
		Description:  "StreamingJob a streaming job object, containing all information associated with the named streaming job.",
		Resolver:     fetchStreamanalyticsJobs,
		Multiplex:    client.SubscriptionMultiplex,
		DeleteFilter: client.DeleteSubscriptionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"subscription_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "subscription_id",
				Description: "Azure subscription id.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAzureSubscription,
			},
			{
				Name:        "sku_name",
				Description: "The name of the SKU.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("StreamingJobProperties.Sku.Name"),
			},
			{
				Name:        "job_id",
				Description: "A GUID uniquely identifying the streaming job.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("StreamingJobProperties.JobID"),
			},
			{
				Name:        "provisioning_state",
				Description: "Describes the provisioning status of the streaming job.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("StreamingJobProperties.ProvisioningState"),
			},
			{
				Name:        "job_state",
				Description: "Describes the state of the streaming job.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("StreamingJobProperties.JobState"),
			},
			{
				Name:        "job_type",
				Description: "Describes the type of the job.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("StreamingJobProperties.JobType"),
			},
			{
				Name:        "output_start_mode",
				Description: "This property should only be utilized when it is desired that the job be started immediately upon creation.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("StreamingJobProperties.OutputStartMode"),
			},
			{
				Name:        "output_start_time",
				Description: "Value is either an ISO-8601 formatted time stamp that indicates the starting point of the output event stream, or null to indicate that the output event stream will start whenever the streaming job is started.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("StreamingJobProperties.OutputStartTime.Time"),
			},
			{
				Name:        "last_output_event_time",
				Description: "Value is either an ISO-8601 formatted timestamp indicating the last output event time of the streaming job or null indicating that output has not yet been produced.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("StreamingJobProperties.LastOutputEventTime.Time"),
			},
			{
				Name:        "events_out_of_order_policy",
				Description: "Indicates the policy to apply to events that arrive out of order in the input event stream.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("StreamingJobProperties.EventsOutOfOrderPolicy"),
			},
			{
				Name:        "output_error_policy",
				Description: "Indicates the policy to apply to events that arrive at the output and cannot be written to the external storage due to being malformed.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("StreamingJobProperties.OutputErrorPolicy"),
			},
			{
				Name:        "events_out_of_order_max_delay",
				Description: "The maximum tolerable delay in seconds where out-of-order events can be adjusted to be back in order.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("StreamingJobProperties.EventsOutOfOrderMaxDelayInSeconds"),
			},
			{
				Name:        "events_late_arrival_max_delay",
				Description: "The maximum tolerable delay in seconds where events arriving late could be included.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("StreamingJobProperties.EventsLateArrivalMaxDelayInSeconds"),
			},
			{
				Name:        "data_locale",
				Description: "The data locale of the stream analytics job.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("StreamingJobProperties.DataLocale"),
			},
			{
				Name:        "compatibility_level",
				Description: "Controls certain runtime behaviors of the streaming job.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("StreamingJobProperties.CompatibilityLevel"),
			},
			{
				Name:        "created_date",
				Description: "Value is an ISO-8601 formatted UTC timestamp indicating when the streaming job was created.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("StreamingJobProperties.CreatedDate.Time"),
			},
			{
				Name:          "transformation_properties_streaming_units",
				Description:   "Specifies the number of streaming units that the streaming job uses.",
				Type:          schema.TypeInt,
				Resolver:      resolveStreamAnalyticsJobsTransformationPropertiesStreamingUnits,
				IgnoreInTests: true,
			},
			{
				Name:          "transformation_properties_valid_streaming_units",
				Description:   "Specifies the valid streaming units a streaming job can scale to.",
				Type:          schema.TypeIntArray,
				Resolver:      resolveStreamAnalyticsJobsTransformationPropertiesValidStreamingUnits,
				IgnoreInTests: true,
			},
			{
				Name:          "transformation_properties_query",
				Description:   "Specifies the query that will be run in the streaming job.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("StreamingJobProperties.Transformation.TransformationProperties.Query"),
				IgnoreInTests: true,
			},
			{
				Name:          "transformation_properties_etag",
				Description:   "The current entity tag for the transformation",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("StreamingJobProperties.Transformation.TransformationProperties.Etag"),
				IgnoreInTests: true,
			},
			{
				Name:          "transformation_id",
				Description:   "Transformation resource Id.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("StreamingJobProperties.Transformation.ID"),
				IgnoreInTests: true,
			},
			{
				Name:          "transformation_name",
				Description:   "Transformation resource name.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("StreamingJobProperties.Transformation.Name"),
				IgnoreInTests: true,
			},
			{
				Name:          "transformation_type",
				Description:   "Transformation resource type.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("StreamingJobProperties.Transformation.Type"),
				IgnoreInTests: true,
			},
			{
				Name:        "etag",
				Description: "The current entity tag for the streaming job.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("StreamingJobProperties.Etag"),
			},
			{
				Name:        "job_storage_account_authentication_mode",
				Description: "Authentication Mode.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("StreamingJobProperties.JobStorageAccount.AuthenticationMode"),
			},
			{
				Name:          "job_storage_account_name",
				Description:   "The name of the Azure Storage account.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("StreamingJobProperties.JobStorageAccount.AccountName"),
				IgnoreInTests: true,
			},
			{
				Name:          "job_storage_account_key",
				Description:   "The account key for the Azure Storage account.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("StreamingJobProperties.JobStorageAccount.AccountKey"),
				IgnoreInTests: true,
			},
			{
				Name:        "content_storage_policy",
				Description: "Valid values are JobStorageAccount and SystemAccount",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("StreamingJobProperties.ContentStoragePolicy"),
			},
			{
				Name:          "cluster_id",
				Description:   "The resource id of cluster.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("StreamingJobProperties.Cluster.ID"),
				IgnoreInTests: true,
			},
			{
				Name:          "identity_tenant_id",
				Description:   "The identity tenantId.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("Identity.TenantID"),
				IgnoreInTests: true,
			},
			{
				Name:          "identity_principal_id",
				Description:   "The identity principal ID.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("Identity.PrincipalID"),
				IgnoreInTests: true,
			},
			{
				Name:          "identity_type",
				Description:   "The identity type",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("Identity.Type"),
				IgnoreInTests: true,
			},
			{
				Name:        "tags",
				Description: "Resource tags.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "location",
				Description: "The geo-location where the resource lives.",
				Type:        schema.TypeString,
			},
			{
				Name:        "id",
				Description: "Fully qualified resource Id for the resource.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "The name of the resource.",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "The type of the resource",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchStreamanalyticsJobs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().StreamAnalytics.Jobs
	result, err := svc.List(ctx, "")
	if err != nil {
		return diag.WrapError(err)
	}
	for result.NotDone() {
		res <- result.Values()
		if err := result.NextWithContext(ctx); err != nil {
			return diag.WrapError(err)
		}
	}
	return nil
}

func resolveStreamAnalyticsJobsTransformationPropertiesStreamingUnits(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	j := resource.Item.(streamanalytics.StreamingJob)
	if j.StreamingJobProperties == nil || j.StreamingJobProperties.Transformation == nil || j.StreamingJobProperties.Transformation.TransformationProperties == nil || j.StreamingJobProperties.Transformation.TransformationProperties.StreamingUnits == nil {
		return nil
	}
	return diag.WrapError(resource.Set(c.Name, *j.StreamingJobProperties.Transformation.TransformationProperties.StreamingUnits))
}

func resolveStreamAnalyticsJobsTransformationPropertiesValidStreamingUnits(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	j := resource.Item.(streamanalytics.StreamingJob)
	if j.StreamingJobProperties == nil || j.StreamingJobProperties.Transformation == nil || j.StreamingJobProperties.Transformation.TransformationProperties == nil || j.StreamingJobProperties.Transformation.TransformationProperties.ValidStreamingUnits == nil {
		return nil
	}
	// repack []int32 as []int as SDK currently does not handle []int32
	u := *j.StreamingJobProperties.Transformation.TransformationProperties.ValidStreamingUnits
	items := make([]int, len(u))
	for i, v := range u {
		items[i] = int(v)
	}
	return diag.WrapError(resource.Set(c.Name, items))
}
