package monitor

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func MonitorActivityLogs() *schema.Table {
	return &schema.Table{
		Name:         "azure_monitor_activity_logs",
		Description:  "Azure network watcher",
		Resolver:     fetchMonitorActivityLogs,
		Multiplex:    client.SubscriptionMultiplex,
		DeleteFilter: client.DeleteSubscriptionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"subscription_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "authorization_action",
				Description: "the permissible actions For instance: microsoftsupport/supporttickets/write",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Authorization.Action"),
			},
			{
				Name:          "authorization_role",
				Description:   "the role of the user For instance: Subscription Admin",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("Authorization.Role"),
				IgnoreInTests: true,
			},
			{
				Name:        "authorization_scope",
				Description: "the scope",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Authorization.Scope"),
			},
			{
				Name:        "claims",
				Description: "key value pairs to identify ARM permissions",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "caller",
				Description: "the email address of the user who has performed the operation, the UPN claim or SPN claim based on availability",
				Type:        schema.TypeString,
			},
			{
				Name:        "description",
				Description: "the description of the event",
				Type:        schema.TypeString,
			},
			{
				Name:        "id",
				Description: "the Id of this event as required by ARM for RBAC It contains the EventDataID and a timestamp information",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "event_data_id",
				Description: "the event data Id This is a unique identifier for an event",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("EventDataID"),
			},
			{
				Name:        "correlation_id",
				Description: "the correlation Id, usually a GUID in the string format The correlation Id is shared among the events that belong to the same uber operation",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("CorrelationID"),
			},
			{
				Name:        "event_name_value",
				Description: "the invariant value",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("EventName.Value"),
			},
			{
				Name:        "event_name_localized_value",
				Description: "the locale specific value",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("EventName.LocalizedValue"),
			},
			{
				Name:        "category_value",
				Description: "the invariant value",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Category.Value"),
			},
			{
				Name:        "category_localized_value",
				Description: "the locale specific value",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Category.LocalizedValue"),
			},
			{
				Name:        "http_request_client_request_id",
				Description: "the client request id",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("HTTPRequest.ClientRequestID"),
			},
			{
				Name:        "http_request_client_ip_address",
				Description: "the client Ip Address",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("HTTPRequest.ClientIPAddress"),
			},
			{
				Name:        "http_request_method",
				Description: "the Http request method",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("HTTPRequest.Method"),
			},
			{
				Name:          "http_request_uri",
				Description:   "the Uri",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("HTTPRequest.URI"),
				IgnoreInTests: true,
			},
			{
				Name:        "level",
				Description: "the event level Possible values include: 'EventLevelCritical', 'EventLevelError', 'EventLevelWarning', 'EventLevelInformational', 'EventLevelVerbose'",
				Type:        schema.TypeString,
			},
			{
				Name:        "resource_group_name",
				Description: "the resource group name of the impacted resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "resource_provider_name_value",
				Description: "the invariant value",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ResourceProviderName.Value"),
			},
			{
				Name:        "resource_provider_name_localized_value",
				Description: "the locale specific value",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ResourceProviderName.LocalizedValue"),
			},
			{
				Name:        "resource_id",
				Description: "the resource uri that uniquely identifies the resource that caused this event",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ResourceID"),
			},
			{
				Name:        "resource_type_value",
				Description: "the invariant value",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ResourceType.Value"),
			},
			{
				Name:        "resource_type_localized_value",
				Description: "the locale specific value",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ResourceType.LocalizedValue"),
			},
			{
				Name:        "operation_id",
				Description: "It is usually a GUID shared among the events corresponding to single operation This value should not be confused with EventName",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("OperationID"),
			},
			{
				Name:        "operation_name_value",
				Description: "the invariant value",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("OperationName.Value"),
			},
			{
				Name:        "operation_name_localized_value",
				Description: "the locale specific value",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("OperationName.LocalizedValue"),
			},
			{
				Name:        "properties",
				Description: "the set of <Key, Value> pairs (usually a Dictionary<String, String>) that includes details about the event",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "status_value",
				Description: "the invariant value",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Status.Value"),
			},
			{
				Name:        "status_localized_value",
				Description: "the locale specific value",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Status.LocalizedValue"),
			},
			{
				Name:        "sub_status_value",
				Description: "the invariant value",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SubStatus.Value"),
			},
			{
				Name:        "sub_status_localized_value",
				Description: "the locale specific value",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SubStatus.LocalizedValue"),
			},
			{
				Name:     "event_timestamp_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("EventTimestamp.Time"),
			},
			{
				Name:     "submission_timestamp_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("SubmissionTimestamp.Time"),
			},
			{
				Name:        "subscription_id",
				Description: "the Azure subscription Id usually a GUID",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SubscriptionID"),
			},
			{
				Name:        "tenant_id",
				Description: "the Azure tenant Id",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TenantID"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchMonitorActivityLogs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	// we fetch activity logs records from now to some point in the past. fetchWindow defines how far that point in the past is.
	const fetchWindow = 24 * time.Hour
	svc := meta.(*client.Client).Services().Monitor.ActivityLogs
	now := time.Now().UTC()
	past := now.Add(-fetchWindow)
	filter := fmt.Sprintf("eventTimestamp ge '%s' and eventTimestamp le '%s'", past.Format(time.RFC3339Nano), now.Format(time.RFC3339Nano))
	response, err := svc.List(ctx, filter, "")
	if err != nil {
		return err
	}
	// azure returns same events sometimes so we have to filter out the duplicates
	seen := make(map[string]struct{})
	for response.NotDone() {
		for _, v := range response.Values() {
			if _, ok := seen[*v.ID]; ok {
				continue
			}
			seen[*v.ID] = struct{}{}
			res <- v
		}
		if err := response.NextWithContext(ctx); err != nil {
			return err
		}
	}
	return nil
}
