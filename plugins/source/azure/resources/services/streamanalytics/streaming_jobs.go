// Auto generated code - DO NOT EDIT.

package streamanalytics

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func StreamingJobs() *schema.Table {
	return &schema.Table{
		Name:      "azure_streamanalytics_streaming_jobs",
		Resolver:  fetchStreamAnalyticsStreamingJobs,
		Multiplex: client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "sku",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Sku"),
			},
			{
				Name:     "job_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("JobID"),
			},
			{
				Name:     "provisioning_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ProvisioningState"),
			},
			{
				Name:     "job_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("JobState"),
			},
			{
				Name:     "job_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("JobType"),
			},
			{
				Name:     "output_start_mode",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OutputStartMode"),
			},
			{
				Name:     "output_start_time",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("OutputStartTime"),
			},
			{
				Name:     "last_output_event_time",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LastOutputEventTime"),
			},
			{
				Name:     "events_out_of_order_policy",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EventsOutOfOrderPolicy"),
			},
			{
				Name:     "output_error_policy",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OutputErrorPolicy"),
			},
			{
				Name:     "events_out_of_order_max_delay_in_seconds",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("EventsOutOfOrderMaxDelayInSeconds"),
			},
			{
				Name:     "events_late_arrival_max_delay_in_seconds",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("EventsLateArrivalMaxDelayInSeconds"),
			},
			{
				Name:     "data_locale",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DataLocale"),
			},
			{
				Name:     "compatibility_level",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CompatibilityLevel"),
			},
			{
				Name:     "created_date",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CreatedDate"),
			},
			{
				Name:     "inputs",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Inputs"),
			},
			{
				Name:     "transformation",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Transformation"),
			},
			{
				Name:     "outputs",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Outputs"),
			},
			{
				Name:     "functions",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Functions"),
			},
			{
				Name:     "etag",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Etag"),
			},
			{
				Name:     "job_storage_account",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("JobStorageAccount"),
			},
			{
				Name:     "content_storage_policy",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ContentStoragePolicy"),
			},
			{
				Name:     "cluster",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Cluster"),
			},
			{
				Name:     "identity",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Identity"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
			{
				Name:     "location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Location"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
		},
	}
}

func fetchStreamAnalyticsStreamingJobs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().StreamAnalytics.StreamingJobs

	response, err := svc.List(ctx, "")

	if err != nil {
		return err
	}

	for response.NotDone() {
		res <- response.Values()
		if err := response.NextWithContext(ctx); err != nil {
			return err
		}
	}

	return nil
}
