package batch

import (
	"context"
	"encoding/json"

	"github.com/Azure/azure-sdk-for-go/services/batch/mgmt/2021-06-01/batch"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func BatchAccounts() *schema.Table {
	return &schema.Table{
		Name:         "azure_batch_accounts",
		Description:  "Account contains information about an Azure Batch account",
		Resolver:     fetchBatchAccounts,
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
				Name:        "account_endpoint",
				Description: "The account endpoint used to interact with the Batch service",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AccountProperties.AccountEndpoint"),
			},
			{
				Name:        "provisioning_state",
				Description: "The provisioned state of the resource",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AccountProperties.ProvisioningState"),
			},
			{
				Name:        "pool_allocation_mode",
				Description: "Possible values include: 'PoolAllocationModeBatchService', 'PoolAllocationModeUserSubscription'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AccountProperties.PoolAllocationMode"),
			},
			{
				Name:          "key_vault_reference_id",
				Description:   "The resource ID of the Azure key vault associated with the Batch account",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("AccountProperties.KeyVaultReference.ID"),
				IgnoreInTests: true,
			},
			{
				Name:          "key_vault_reference_url",
				Description:   "The URL of the Azure key vault associated with the Batch account",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("AccountProperties.KeyVaultReference.URL"),
				IgnoreInTests: true,
			},
			{
				Name:        "public_network_access",
				Description: "If not specified, the default value is 'enabled'. Possible values include: 'PublicNetworkAccessTypeEnabled', 'PublicNetworkAccessTypeDisabled'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AccountProperties.PublicNetworkAccess"),
			},
			{
				Name:        "auto_storage_last_key_sync_time",
				Description: "The UTC time at which storage keys were last synchronized with the Batch account.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("AccountProperties.AutoStorage.LastKeySync.Time"),
			},
			{
				Name:        "auto_storage_storage_account_id",
				Description: "The resource ID of the storage account to be used for auto-storage account",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AccountProperties.AutoStorage.StorageAccountID"),
			},
			{
				Name:        "auto_storage_authentication_mode",
				Description: "The authentication mode which the Batch service will use to manage the auto-storage account",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AccountProperties.AutoStorage.AuthenticationMode"),
			},
			{
				Name:          "auto_storage_node_identity_reference_resource_id",
				Description:   "The ARM resource id of the user assigned identity",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("AccountProperties.AutoStorage.NodeIdentityReference.ResourceID"),
				IgnoreInTests: true,
			},
			{
				Name:        "encryption_key_source",
				Description: "Type of the key source. Possible values include: 'KeySourceMicrosoftBatch', 'KeySourceMicrosoftKeyVault'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AccountProperties.Encryption.KeySource"),
			},
			{
				Name:          "encryption_key_vault_properties_key_identifier",
				Description:   "Full path to the versioned secret",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("AccountProperties.Encryption.KeyVaultProperties.KeyIdentifier"),
				IgnoreInTests: true,
			},
			{
				Name:        "dedicated_core_quota",
				Description: "For accounts with PoolAllocationMode set to UserSubscription, quota is managed on the subscription so this value is not returned",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("AccountProperties.DedicatedCoreQuota"),
			},
			{
				Name:        "low_priority_core_quota",
				Description: "For accounts with PoolAllocationMode set to UserSubscription, quota is managed on the subscription so this value is not returned",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("AccountProperties.LowPriorityCoreQuota"),
			},
			{
				Name:        "dedicated_core_quota_per_vm_family",
				Description: "A list of the dedicated core quota per Virtual Machine family for the Batch account",
				Type:        schema.TypeJSON,
				Resolver:    resolveBatchAccountsDedicatedCoreQuotaPerVmFamily,
			},
			{
				Name:        "dedicated_core_quota_per_vm_family_enforced",
				Description: "Batch is transitioning its core quota system for dedicated cores to be enforced per Virtual Machine family",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("AccountProperties.DedicatedCoreQuotaPerVMFamilyEnforced"),
			},
			{
				Name:        "pool_quota",
				Description: "The pool quota for the Batch account.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("AccountProperties.PoolQuota"),
			},
			{
				Name:        "active_job_and_job_schedule_quota",
				Description: "The active job and job schedule quota for the Batch account.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("AccountProperties.ActiveJobAndJobScheduleQuota"),
			},
			{
				Name:        "allowed_authentication_modes",
				Description: "List of allowed authentication modes for the Batch account that can be used to authenticate with the data plane",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("AccountProperties.AllowedAuthenticationModes"),
			},
			{
				Name:          "identity_principal_id",
				Description:   "The principal id of the Batch account",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("Identity.PrincipalID"),
				IgnoreInTests: true,
			},
			{
				Name:          "identity_tenant_id",
				Description:   "The tenant id associated with the Batch account",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("Identity.TenantID"),
				IgnoreInTests: true,
			},
			{
				Name:        "identity_type",
				Description: "The type of identity used for the Batch account",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Identity.Type"),
			},
			{
				Name:          "identity_user_assigned_identities",
				Description:   "The list of user identities associated with the Batch account",
				Type:          schema.TypeJSON,
				Resolver:      schema.PathResolver("Identity.UserAssignedIdentities"),
				IgnoreInTests: true,
			},
			{
				Name:        "id",
				Description: "The ID of the resource",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "The name of the resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "The type of the resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "location",
				Description: "The location of the resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "The tags of the resource",
				Type:        schema.TypeJSON,
			},
		},
		Relations: []*schema.Table{
			{
				Name:          "azure_batch_account_private_endpoint_connections",
				Description:   "PrivateEndpointConnection contains information about a private link resource",
				Resolver:      fetchBatchAccountPrivateEndpointConnections,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "account_cq_id",
						Description: "Unique CloudQuery ID of azure_batch_accounts table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "provisioning_state",
						Description: "Possible values include: 'PrivateEndpointConnectionProvisioningStateSucceeded', 'PrivateEndpointConnectionProvisioningStateUpdating', 'PrivateEndpointConnectionProvisioningStateFailed'",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("PrivateEndpointConnectionProperties.ProvisioningState"),
					},
					{
						Name:        "private_endpoint_id",
						Description: "The resource id of the private endpoint resource from Microsoft.Network provider.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("PrivateEndpointConnectionProperties.PrivateEndpoint.ID"),
					},
					{
						Name:        "private_link_connection_status",
						Description: "Possible values include: 'PrivateLinkServiceConnectionStatusApproved', 'PrivateLinkServiceConnectionStatusPending', 'PrivateLinkServiceConnectionStatusRejected', 'PrivateLinkServiceConnectionStatusDisconnected'",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("PrivateEndpointConnectionProperties.PrivateLinkServiceConnectionState.Status"),
					},
					{
						Name:        "private_link_connection_description",
						Description: "The description for the private link service connection state.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("PrivateEndpointConnectionProperties.PrivateLinkServiceConnectionState.Description"),
					},
					{
						Name:        "private_link_connection_action_required",
						Description: "A description of any extra actions that may be required.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("PrivateEndpointConnectionProperties.PrivateLinkServiceConnectionState.ActionRequired"),
					},
					{
						Name:        "id",
						Description: "The ID of the resource",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
					{
						Name:        "name",
						Description: "The name of the resource",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "The type of the resource",
						Type:        schema.TypeString,
					},
					{
						Name:        "etag",
						Description: "The ETag of the resource, used for concurrency statements",
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

func fetchBatchAccounts(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Batch.Account
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
func resolveBatchAccountsDedicatedCoreQuotaPerVmFamily(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	account := resource.Item.(batch.Account)
	if account.DedicatedCoreQuotaPerVMFamily == nil {
		return nil
	}
	b, err := json.Marshal(account.DedicatedCoreQuotaPerVMFamily)
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, b))
}
func fetchBatchAccountPrivateEndpointConnections(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	account := parent.Item.(batch.Account)
	if account.PrivateEndpointConnections == nil {
		return nil
	}
	res <- *account.PrivateEndpointConnections
	return nil
}
