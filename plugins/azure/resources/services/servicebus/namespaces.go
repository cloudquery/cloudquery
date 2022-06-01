package servicebus

import (
	"context"
	"encoding/json"

	"github.com/Azure/azure-sdk-for-go/services/preview/servicebus/mgmt/2021-06-01-preview/servicebus"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func ServicebusNamespaces() *schema.Table {
	return &schema.Table{
		Name:         "azure_servicebus_namespaces",
		Description:  "SBNamespace description of a namespace resource.",
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
				Description: "Name of this SKU.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Sku.Name"),
			},
			{
				Name:        "sku_tier",
				Description: "The billing tier of this particular SKU.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Sku.Tier"),
			},
			{
				Name:          "sku_capacity",
				Description:   "The specified messaging units for the tier.",
				Type:          schema.TypeInt,
				Resolver:      schema.PathResolver("Sku.Capacity"),
				IgnoreInTests: true,
			},
			{
				Name:          "identity_principal_id",
				Description:   "ObjectId from the KeyVault.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("Identity.PrincipalID"),
				IgnoreInTests: true,
			},
			{
				Name:          "identity_tenant_id",
				Description:   "TenantId from the KeyVault.",
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
				Description:   "Properties for User Assigned Identities.",
				Type:          schema.TypeJSON,
				Resolver:      schema.PathResolver("Identity.UserAssignedIdentities"),
				IgnoreInTests: true,
			},
			{
				Name:          "system_data",
				Description:   "The system meta data relating to this resource.",
				Type:          schema.TypeJSON,
				Resolver:      resolveServicebusNamespacesSystemData,
				IgnoreInTests: true,
			},
			{
				Name:        "location",
				Description: "The Geo-location where the resource lives.",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "Resource tags.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "id",
				Description: "Resource Id.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "Resource name.",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "Resource type.",
				Type:        schema.TypeString,
			},

			{
				Name:        "provisioning_state",
				Description: "Provisioning state of the namespace.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SBNamespaceProperties.ProvisioningState"),
			},
			{
				Name:        "status",
				Description: "Status of the namespace.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SBNamespaceProperties.Status"),
			},
			{
				Name:        "created_at",
				Description: "The time the namespace was created.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("SBNamespaceProperties.CreatedAt.Time"),
			},
			{
				Name:        "updated_at",
				Description: "The time the namespace was updated.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("SBNamespaceProperties.UpdatedAt.Time"),
			},
			{
				Name:        "service_bus_endpoint",
				Description: "Endpoint you can use to perform Service Bus operations.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SBNamespaceProperties.ServiceBusEndpoint"),
			},
			{
				Name:        "metric_id",
				Description: "Identifier for Azure Insights metrics.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SBNamespaceProperties.MetricID"),
			},
			{
				Name:        "zone_redundant",
				Description: "Enabling this property creates a Premium Service Bus Namespace in regions supported availability zones.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("SBNamespaceProperties.ZoneRedundant"),
			},
			{
				Name:          "key_vault_properties",
				Description:   "Properties of KeyVault (BYOK Encryption).",
				Type:          schema.TypeJSON,
				Resolver:      resolveServicebusNamespacesKeyVaultProperties,
				IgnoreInTests: true,
			},
			{
				Name:        "key_source",
				Description: "Enumerates the possible value of keySource for Encryption.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SBNamespaceProperties.Encryption.KeySource"),
			},
			{
				Name:          "require_infrastructure_encryption",
				Description:   "Enable Infrastructure Encryption (Double Encryption).",
				Type:          schema.TypeBool,
				Resolver:      schema.PathResolver("SBNamespaceProperties.Encryption.RequireInfrastructureEncryption"),
				IgnoreInTests: true,
			},
			{
				Name:        "disable_local_auth",
				Description: "This property disables SAS authentication for the Service Bus namespace.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("SBNamespaceProperties.DisableLocalAuth"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:          "azure_servicebus_namespace_private_endpoint_connections",
				Description:   "List of private endpoint connections.",
				Resolver:      fetchServicebusNamespacesPrivateEndpointConnections,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "namespace_cq_id",
						Description: "Unique ID of azure_servicebus_namespaces table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "id",
						Description: "Resource Id.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
					{
						Name:        "name",
						Description: "Resource name.",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "Resource type.",
						Type:        schema.TypeString,
					},
					{
						Name:        "system_data",
						Description: "The system meta data relating to this resource.",
						Type:        schema.TypeJSON,
						Resolver:    resolveServicebusNamespacePrivateEndpointConnectionSystemData,
					},
					{
						Name:        "status",
						Description: "The private link service connection status.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("PrivateEndpointConnectionProperties.PrivateLinkServiceConnectionState.Status"),
					},
					{
						Name:        "status_description",
						Description: "The private link service connection description.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("PrivateEndpointConnectionProperties.PrivateLinkServiceConnectionState.Description"),
					},
					{
						Name:        "provisioning_state",
						Description: "State of the private endpoint connection.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("PrivateEndpointConnectionProperties.ProvisioningState"),
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

func resolveServicebusNamespacesSystemData(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	n := resource.Item.(servicebus.SBNamespace)
	if n.SystemData == nil {
		return nil
	}
	b, err := json.Marshal(n.SystemData)
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, b))
}

func resolveServicebusNamespacesKeyVaultProperties(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	n := resource.Item.(servicebus.SBNamespace)
	if n.SBNamespaceProperties == nil || n.SBNamespaceProperties.Encryption == nil || n.SBNamespaceProperties.Encryption.KeyVaultProperties == nil {
		return nil
	}
	b, err := json.Marshal(n.SBNamespaceProperties.Encryption.KeyVaultProperties)
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, b))
}

func fetchServicebusNamespacesPrivateEndpointConnections(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	n := parent.Item.(servicebus.SBNamespace)
	if n.SBNamespaceProperties == nil || n.SBNamespaceProperties.PrivateEndpointConnections == nil {
		return nil
	}
	res <- *n.SBNamespaceProperties.PrivateEndpointConnections
	return nil
}

func resolveServicebusNamespacePrivateEndpointConnectionSystemData(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	n := resource.Item.(servicebus.PrivateEndpointConnection)
	if n.SystemData == nil {
		return nil
	}
	b, err := json.Marshal(n.SystemData)
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, b))
}
