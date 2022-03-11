package eventhub

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/preview/eventhub/mgmt/2018-01-01-preview/eventhub"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func EventHubNamespaces() *schema.Table {
	return &schema.Table{
		Name:         "azure_eventhub_namespaces",
		Description:  "Azure EventHub namespace",
		Resolver:     fetchEventhubNamespaces,
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
				Name:        "sku_capacity",
				Description: "The Event Hubs throughput units, value should be 0 to 20 throughput units.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Sku.Capacity"),
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
				Description: "Enumerates the possible value Identity type, which currently supports only 'SystemAssigned'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Identity.Type"),
			},
			{
				Name:        "provisioning_state",
				Description: "Provisioning state of the Namespace.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("EHNamespaceProperties.ProvisioningState"),
			},
			{
				Name:     "created_at_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("EHNamespaceProperties.CreatedAt.Time"),
			},
			{
				Name:     "updated_at_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("EHNamespaceProperties.UpdatedAt.Time"),
			},
			{
				Name:        "service_bus_endpoint",
				Description: "Endpoint you can use to perform Service Bus operations.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("EHNamespaceProperties.ServiceBusEndpoint"),
			},
			{
				Name:          "cluster_arm_id",
				Description:   "Cluster ARM ID of the Namespace.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("EHNamespaceProperties.ClusterArmID"),
				IgnoreInTests: true,
			},
			{
				Name:        "metric_id",
				Description: "Identifier for Azure Insights metrics.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("EHNamespaceProperties.MetricID"),
			},
			{
				Name:        "is_auto_inflate_enabled",
				Description: "Value that indicates whether AutoInflate is enabled for eventhub namespace.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("EHNamespaceProperties.IsAutoInflateEnabled"),
			},
			{
				Name:        "maximum_throughput_units",
				Description: "Upper limit of throughput units when AutoInflate is enabled, value should be within 0 to 20 throughput units",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("EHNamespaceProperties.MaximumThroughputUnits"),
			},
			{
				Name:        "kafka_enabled",
				Description: "Value that indicates whether Kafka is enabled for eventhub namespace.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("EHNamespaceProperties.KafkaEnabled"),
			},
			{
				Name:        "zone_redundant",
				Description: "Enabling this property creates a Standard Event Hubs Namespace in regions supported availability zones.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("EHNamespaceProperties.ZoneRedundant"),
			},
			{
				Name:        "encryption_key_source",
				Description: "Enumerates the possible value of keySource for Encryption",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("EHNamespaceProperties.Encryption.KeySource"),
			},
			{
				Name:        "location",
				Description: "Resource location.",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "Resource tags.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "id",
				Description: "Resource ID.",
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
				Name:        "network_rule_set",
				Description: "Network rule set for a namespace.",
				Type:        schema.TypeJSON,
				Resolver:    resolveNamespaceNetworkRuleSet,
			},
		},
		Relations: []*schema.Table{
			{
				Name:          "azure_eventhub_namespace_encryption_key_vault_properties",
				Description:   "KeyVaultProperties properties to configure keyVault Properties",
				Resolver:      fetchEventhubNamespaceEncryptionKeyVaultProperties,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "namespace_cq_id",
						Description: "Unique CloudQuery ID of azure_eventhub_namespaces table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "key_name",
						Description: "Name of the Key from KeyVault",
						Type:        schema.TypeString,
					},
					{
						Name:        "key_vault_uri",
						Description: "Uri of KeyVault",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("KeyVaultURI"),
					},
					{
						Name:        "key_version",
						Description: "Key Version",
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

func fetchEventhubNamespaces(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().EventHub
	response, err := svc.List(ctx)
	if err != nil {
		return diag.WrapError(err)
	}
	for response.NotDone() {
		res <- response.Values()
		if err := response.NextWithContext(ctx); err != nil {
			return diag.WrapError(err)
		}
	}
	return nil
}
func fetchEventhubNamespaceEncryptionKeyVaultProperties(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	namespace, ok := parent.Item.(eventhub.EHNamespace)
	if !ok {
		return fmt.Errorf("expected to have eventhub.EHNamespace but got %T", parent.Item)
	}
	if namespace.Encryption == nil || namespace.Encryption.KeyVaultProperties == nil {
		return nil
	}
	res <- *namespace.Encryption.KeyVaultProperties
	return nil
}

func resolveNamespaceNetworkRuleSet(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	svc := meta.(*client.Client).Services().EventHub
	namespace := resource.Item.(eventhub.EHNamespace)
	details, err := client.ParseResourceID(*namespace.ID)
	if err != nil {
		return diag.WrapError(err)
	}
	rs, err := svc.GetNetworkRuleSet(ctx, details.ResourceGroup, *namespace.Name)
	if err != nil {
		return diag.WrapError(err)
	}
	b, err := json.Marshal(rs)
	if err != nil {
		return diag.WrapError(err)
	}
	return resource.Set(c.Name, b)
}
