package monitor

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2019-11-01-preview/insights"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func MonitorActivityLogAlerts() *schema.Table {
	return &schema.Table{
		Name:         "azure_monitor_activity_log_alerts",
		Description:  "ActivityLogAlertResource an activity log alert resource",
		Resolver:     fetchMonitorActivityLogAlerts,
		Multiplex:    client.SubscriptionMultiplex,
		DeleteFilter: client.DeleteSubscriptionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"subscription_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "subscription_id",
				Description: "Azure subscription id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAzureSubscription,
			},
			{
				Name:        "scopes",
				Description: "A list of resourceIds that will be used as prefixes The alert will only apply to activityLogs with resourceIds that fall under one of these prefixes This list must include at least one item",
				Type:        schema.TypeStringArray,
				Resolver:    resolveMonitorActivityLogAlertScopes,
			},
			{
				Name:        "enabled",
				Description: "Indicates whether this activity log alert is enabled If an activity log alert is not enabled, then none of its actions will be activated",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("ActivityLogAlert.Enabled"),
			},
			{
				Name:        "description",
				Description: "A description of this activity log alert",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ActivityLogAlert.Description"),
			},
			{
				Name:        "id",
				Description: "Azure resource Id",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "Azure resource name",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "Azure resource type",
				Type:        schema.TypeString,
			},
			{
				Name:        "location",
				Description: "Resource location",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "Resource tags",
				Type:        schema.TypeJSON,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "azure_monitor_activity_log_alert_conditions",
				Description: "ActivityLogAlertLeafCondition an Activity Log alert condition that is met by comparing an activity log field and value",
				Resolver:    fetchMonitorActivityLogAlertConditions,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"activity_log_alert_cq_id", "field"}},
				Columns: []schema.Column{
					{
						Name:        "activity_log_alert_cq_id",
						Description: "Unique ID of azure_monitor_activity_log_alerts table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "activity_log_alert_id",
						Description: "ID of azure_monitor_activity_log_alerts table (FK)",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "field",
						Description: "The name of the field that this condition will examine The possible values for this field are (case-insensitive): 'resourceId', 'category', 'caller', 'level', 'operationName', 'resourceGroup', 'resourceProvider', 'status', 'subStatus', 'resourceType', or anything beginning with 'properties'",
						Type:        schema.TypeString,
						Resolver:    resolveMonitorActivityLogAlertConditionField,
					},
					{
						Name:        "equals",
						Description: "The field value will be compared to this value (case-insensitive) to determine if the condition is met",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "azure_monitor_activity_log_alert_action_groups",
				Description: "ActivityLogAlertActionGroup a pointer to an Azure Action Group",
				Resolver:    fetchMonitorActivityLogAlertActionGroups,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"activity_log_alert_cq_id", "action_group_id"}},
				Columns: []schema.Column{
					{
						Name:        "activity_log_alert_cq_id",
						Description: "Unique ID of azure_monitor_activity_log_alerts table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "activity_log_alert_id",
						Description: "ID of azure_monitor_activity_log_alerts table (FK)",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "action_group_id",
						Description: "The resourceId of the action group This cannot be null or empty",
						Type:        schema.TypeString,
						Resolver:    resolveMonitorActivityLogAlertActionGroupActionGroupID,
					},
					{
						Name:        "webhook_properties",
						Description: "the dictionary of custom properties to include with the post operation These data are appended to the webhook payload",
						Type:        schema.TypeJSON,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchMonitorActivityLogAlerts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	svc := meta.(*client.Client).Services().Monitor.ActivityLogAlerts
	response, err := svc.ListBySubscriptionID(ctx)
	if err != nil {
		return err
	}
	if response.Value == nil {
		return nil
	}
	res <- *response.Value
	return nil
}
func resolveMonitorActivityLogAlertScopes(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r, ok := resource.Item.(insights.ActivityLogAlertResource)
	if !ok {
		return fmt.Errorf("expected to have insights.ActivityLogAlertResource but got %T", resource.Item)
	}
	if r.Scopes == nil {
		return nil
	}
	return resource.Set(c.Name, *r.Scopes)
}
func fetchMonitorActivityLogAlertConditions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(insights.ActivityLogAlertResource)
	if !ok {
		return fmt.Errorf("expected to have insights.ActivityLogAlertResource but got %T", parent.Item)
	}
	if p.Condition == nil {
		return nil
	}
	res <- *p.Condition.AllOf
	return nil
}
func resolveMonitorActivityLogAlertConditionField(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r, ok := resource.Item.(insights.ActivityLogAlertLeafCondition)
	if !ok {
		return fmt.Errorf("expected to have insights.ActivityLogAlertLeafCondition but got %T", resource.Item)
	}
	if r.Field == nil {
		return nil
	}
	return resource.Set(c.Name, *r.Field)
}
func fetchMonitorActivityLogAlertActionGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(insights.ActivityLogAlertResource)
	if !ok {
		return fmt.Errorf("expected to have insights.ActivityLogAlertResource but got %T", parent.Item)
	}
	if p.Actions == nil {
		return nil
	}
	res <- *p.Actions.ActionGroups
	return nil
}
func resolveMonitorActivityLogAlertActionGroupActionGroupID(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r, ok := resource.Item.(insights.ActivityLogAlertActionGroup)
	if !ok {
		return fmt.Errorf("expected to have insights.ActivityLogAlertActionGroup but got %T", resource.Item)
	}
	if r.ActionGroupID == nil {
		return nil
	}
	return resource.Set(c.Name, *r.ActionGroupID)
}
