package servicebus

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/preview/servicebus/mgmt/2021-06-01-preview/servicebus"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource namespaces --config gen.hcl --output .
func Namespaces() *schema.Table {
	return &schema.Table{
		Name:         "azure_servicebus_namespaces",
		Description:  "SBNamespace description of a namespace resource",
		Resolver:     fetchServicebusNamespaces,
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
				Name:        "sku_name",
				Description: "Name of this SKU",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Sku.Name"),
			},
			{
				Name:        "sku_tier",
				Description: "The billing tier of this particular SKU",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Sku.Tier"),
			},
			{
				Name:          "sku_capacity",
				Description:   "The specified messaging units for the tier",
				Type:          schema.TypeInt,
				Resolver:      schema.PathResolver("Sku.Capacity"),
				IgnoreInTests: true,
			},
			{
				Name:          "identity_principal_id",
				Description:   "ObjectId from the KeyVault",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("Identity.PrincipalID"),
				IgnoreInTests: true,
			},
			{
				Name:          "identity_tenant_id",
				Description:   "TenantId from the KeyVault",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("Identity.TenantID"),
				IgnoreInTests: true,
			},
			{
				Name:        "identity_type",
				Description: "Type of managed service identity",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Identity.Type"),
			},
			{
				Name:          "user_assigned_identities",
				Description:   "Properties for User Assigned Identities",
				Type:          schema.TypeJSON,
				Resolver:      schema.PathResolver("Identity.UserAssignedIdentities"),
				IgnoreInTests: true,
			},
			{
				Name:          "system_data",
				Description:   "The system meta data relating to this resource",
				Type:          schema.TypeJSON,
				IgnoreInTests: true,
			},
			{
				Name:        "provisioning_state",
				Description: "Provisioning state of the namespace",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SBNamespaceProperties.ProvisioningState"),
			},
			{
				Name:        "status",
				Description: "Status of the namespace",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SBNamespaceProperties.Status"),
			},
			{
				Name:     "created_at_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("SBNamespaceProperties.CreatedAt.Time"),
			},
			{
				Name:     "updated_at_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("SBNamespaceProperties.UpdatedAt.Time"),
			},
			{
				Name:        "service_bus_endpoint",
				Description: "Endpoint you can use to perform Service Bus operations",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SBNamespaceProperties.ServiceBusEndpoint"),
			},
			{
				Name:        "metric_id",
				Description: "Identifier for Azure Insights metrics",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SBNamespaceProperties.MetricID"),
			},
			{
				Name:        "zone_redundant",
				Description: "Enabling this property creates a Premium Service Bus Namespace in regions supported availability zones",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("SBNamespaceProperties.ZoneRedundant"),
			},
			{
				Name:          "key_vault_properties",
				Description:   "Properties of KeyVault",
				Type:          schema.TypeJSON,
				Resolver:      schema.PathResolver("SBNamespaceProperties.Encryption.KeyVaultProperties"),
				IgnoreInTests: true,
			},
			{
				Name:        "key_source",
				Description: "Enumerates the possible value of keySource for Encryption",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SBNamespaceProperties.Encryption.KeySource"),
			},
			{
				Name:          "require_infrastructure_encryption",
				Description:   "Enable Infrastructure Encryption (Double Encryption)",
				Type:          schema.TypeBool,
				Resolver:      schema.PathResolver("SBNamespaceProperties.Encryption.RequireInfrastructureEncryption"),
				IgnoreInTests: true,
			},
			{
				Name:        "disable_local_auth",
				Description: "This property disables SAS authentication for the Service Bus namespace",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("SBNamespaceProperties.DisableLocalAuth"),
			},
			{
				Name:        "location",
				Description: "The Geo-location where the resource lives",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "Resource tags",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "id",
				Description: "Resource Id",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "Resource name",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "Resource type",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:          "azure_servicebus_namespace_private_endpoint_connections",
				Description:   "List of private endpoint connections",
				Resolver:      fetchServicebusNamespacePrivateEndpointConnections,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "namespace_cq_id",
						Description: "Unique CloudQuery ID of azure_servicebus_namespaces table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "private_endpoint_id",
						Description: "The ARM identifier for Private Endpoint",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("PrivateEndpointConnectionProperties.PrivateEndpoint.ID"),
					},
					{
						Name:        "status",
						Description: "Status of the connection",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("PrivateEndpointConnectionProperties.PrivateLinkServiceConnectionState.Status"),
					},
					{
						Name:        "status_description",
						Description: "Description of the connection state",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("PrivateEndpointConnectionProperties.PrivateLinkServiceConnectionState.Description"),
					},
					{
						Name:        "provisioning_state",
						Description: "Provisioning state of the Private Endpoint Connection",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("PrivateEndpointConnectionProperties.ProvisioningState"),
					},
					{
						Name:        "system_data",
						Description: "The system meta data relating to this resource",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "id",
						Description: "Resource Id",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
					{
						Name:        "name",
						Description: "Resource name",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "Resource type",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "azure_servicebus_namespace_topics",
				Description: "Description of servicebus namespace topic resource",
				Resolver:    fetchServicebusNamespaceTopics,
				Columns: []schema.Column{
					{
						Name:        "namespace_cq_id",
						Description: "Unique CloudQuery ID of azure_servicebus_namespaces table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "size_in_bytes",
						Description: "Size of the topic, in bytes",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("SBTopicProperties.SizeInBytes"),
					},
					{
						Name:     "created_at_time",
						Type:     schema.TypeTimestamp,
						Resolver: schema.PathResolver("SBTopicProperties.CreatedAt.Time"),
					},
					{
						Name:     "updated_at_time",
						Type:     schema.TypeTimestamp,
						Resolver: schema.PathResolver("SBTopicProperties.UpdatedAt.Time"),
					},
					{
						Name:     "accessed_at_time",
						Type:     schema.TypeTimestamp,
						Resolver: schema.PathResolver("SBTopicProperties.AccessedAt.Time"),
					},
					{
						Name:        "subscription_count",
						Description: "Number of subscriptions",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("SBTopicProperties.SubscriptionCount"),
					},
					{
						Name:        "count_details_active_message_count",
						Description: "Number of active messages in the queue, topic, or subscription",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("SBTopicProperties.CountDetails.ActiveMessageCount"),
					},
					{
						Name:        "count_details_dead_letter_message_count",
						Description: "Number of messages that are dead lettered",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("SBTopicProperties.CountDetails.DeadLetterMessageCount"),
					},
					{
						Name:        "count_details_scheduled_message_count",
						Description: "Number of scheduled messages",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("SBTopicProperties.CountDetails.ScheduledMessageCount"),
					},
					{
						Name:        "count_details_transfer_message_count",
						Description: "Number of messages transferred to another queue, topic, or subscription",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("SBTopicProperties.CountDetails.TransferMessageCount"),
					},
					{
						Name:        "count_details_transfer_dead_letter_message_count",
						Description: "Number of messages transferred into dead letters",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("SBTopicProperties.CountDetails.TransferDeadLetterMessageCount"),
					},
					{
						Name:        "default_message_time_to_live",
						Description: "ISO 8601 Default message timespan to live value",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SBTopicProperties.DefaultMessageTimeToLive"),
					},
					{
						Name:        "max_size_in_megabytes",
						Description: "Maximum size of the topic in megabytes, which is the size of the memory allocated for the topic",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("SBTopicProperties.MaxSizeInMegabytes"),
					},
					{
						Name:        "max_message_size_in_kilobytes",
						Description: "Maximum size (in KB) of the message payload that can be accepted by the topic",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("SBTopicProperties.MaxMessageSizeInKilobytes"),
					},
					{
						Name:        "requires_duplicate_detection",
						Description: "Value indicating if this topic requires duplicate detection",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("SBTopicProperties.RequiresDuplicateDetection"),
					},
					{
						Name:        "duplicate_detection_history_time_window",
						Description: "ISO8601 timespan structure that defines the duration of the duplicate detection history",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SBTopicProperties.DuplicateDetectionHistoryTimeWindow"),
					},
					{
						Name:        "enable_batched_operations",
						Description: "Value that indicates whether server-side batched operations are enabled",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("SBTopicProperties.EnableBatchedOperations"),
					},
					{
						Name:        "status",
						Description: "Enumerates the possible values for the status of a messaging entity",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SBTopicProperties.Status"),
					},
					{
						Name:        "support_ordering",
						Description: "Value that indicates whether the topic supports ordering",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("SBTopicProperties.SupportOrdering"),
					},
					{
						Name:        "auto_delete_on_idle",
						Description: "ISO 8601 timespan idle interval after which the topic is automatically deleted",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SBTopicProperties.AutoDeleteOnIdle"),
					},
					{
						Name:        "enable_partitioning",
						Description: "Value that indicates whether the topic to be partitioned across multiple message brokers is enabled",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("SBTopicProperties.EnablePartitioning"),
					},
					{
						Name:        "enable_express",
						Description: "Value that indicates whether Express Entities are enabled",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("SBTopicProperties.EnableExpress"),
					},
					{
						Name:          "system_data",
						Description:   "The system meta data relating to this resource",
						Type:          schema.TypeJSON,
						IgnoreInTests: true,
					},
					{
						Name:        "id",
						Description: "Resource Id",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
					{
						Name:        "name",
						Description: "Resource name",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "Resource type",
						Type:        schema.TypeString,
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "azure_servicebus_namespace_topic_authorization_rules",
						Description: "Description of servicebus namespace topic authorization rules",
						Resolver:    fetchServicebusNamespaceTopicAuthorizationRules,
						Columns: []schema.Column{
							{
								Name:        "namespace_topic_cq_id",
								Description: "Unique CloudQuery ID of azure_servicebus_namespace_topics table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:     "access_keys",
								Type:     schema.TypeJSON,
								Resolver: ResolveServicebusNamespaceTopicAuthorizationRuleAccessKeys,
							},
							{
								Name:        "rights",
								Description: "The rights associated with the rule",
								Type:        schema.TypeStringArray,
								Resolver:    schema.PathResolver("SBAuthorizationRuleProperties.Rights"),
							},
							{
								Name:          "system_data",
								Description:   "The system meta data relating to this resource",
								Type:          schema.TypeJSON,
								IgnoreInTests: true,
							},
							{
								Name:        "id",
								Description: "Resource Id",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ID"),
							},
							{
								Name:        "name",
								Description: "Resource name",
								Type:        schema.TypeString,
							},
							{
								Name:        "type",
								Description: "Resource type",
								Type:        schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchServicebusNamespaces(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Servicebus.Namespaces
	r, err := svc.List(ctx)
	if err != nil {
		return diag.WrapError(err)
	}
	for r.NotDone() {
		res <- r.Values()
		if err := r.NextWithContext(ctx); err != nil {
			return diag.WrapError(err)
		}
	}
	return nil
}
func fetchServicebusNamespacePrivateEndpointConnections(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	n := parent.Item.(servicebus.SBNamespace)
	if n.SBNamespaceProperties == nil || n.SBNamespaceProperties.PrivateEndpointConnections == nil {
		return nil
	}
	res <- *n.SBNamespaceProperties.PrivateEndpointConnections
	return nil
}
func fetchServicebusNamespaceTopics(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Servicebus.Topics
	n := parent.Item.(servicebus.SBNamespace)
	details, err := client.ParseResourceID(*n.ID)
	if err != nil {
		return diag.WrapError(err)
	}
	r, err := svc.ListByNamespace(ctx, details.ResourceGroup, details.ResourceName, nil, nil)
	if err != nil {
		return diag.WrapError(err)
	}
	for r.NotDone() {
		res <- r.Values()
		if err := r.NextWithContext(ctx); err != nil {
			return diag.WrapError(err)
		}
	}
	return nil
}
func fetchServicebusNamespaceTopicAuthorizationRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Servicebus.Topics
	n := parent.Parent.Item.(servicebus.SBNamespace)
	t := parent.Item.(servicebus.SBTopic)
	details, err := client.ParseResourceID(*t.ID)
	if err != nil {
		return diag.WrapError(err)
	}
	r, err := svc.ListAuthorizationRules(ctx, details.ResourceGroup, *n.Name, details.ResourceName)
	if err != nil {
		return diag.WrapError(err)
	}
	for r.NotDone() {
		res <- r.Values()
		if err := r.NextWithContext(ctx); err != nil {
			return diag.WrapError(err)
		}
	}
	return nil
}
func ResolveServicebusNamespaceTopicAuthorizationRuleAccessKeys(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	svc := meta.(*client.Client).Services().Servicebus.Topics
	n := resource.Parent.Parent.Item.(servicebus.SBNamespace)
	t := resource.Parent.Item.(servicebus.SBTopic)
	a := resource.Item.(servicebus.SBAuthorizationRule)
	details, err := client.ParseResourceID(*a.ID)
	if err != nil {
		return diag.WrapError(err)
	}
	r, err := svc.ListKeys(ctx, details.ResourceGroup, *n.Name, *t.Name, details.ResourceName)
	if err != nil {
		return diag.WrapError(err)
	}
	j, err := r.MarshalJSON()
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, j))
}
