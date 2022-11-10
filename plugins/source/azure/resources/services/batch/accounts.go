// Auto generated code - DO NOT EDIT.

package batch

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Accounts() *schema.Table {
	return &schema.Table{
		Name:        "azure_batch_accounts",
		Description: `https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/batch/mgmt/2021-06-01/batch#Account`,
		Resolver:    fetchBatchAccounts,
		Multiplex:   client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "account_endpoint",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountEndpoint"),
			},
			{
				Name:     "provisioning_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ProvisioningState"),
			},
			{
				Name:     "pool_allocation_mode",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PoolAllocationMode"),
			},
			{
				Name:     "key_vault_reference",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("KeyVaultReference"),
			},
			{
				Name:     "public_network_access",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PublicNetworkAccess"),
			},
			{
				Name:     "private_endpoint_connections",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PrivateEndpointConnections"),
			},
			{
				Name:     "auto_storage",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AutoStorage"),
			},
			{
				Name:     "encryption",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Encryption"),
			},
			{
				Name:     "dedicated_core_quota",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("DedicatedCoreQuota"),
			},
			{
				Name:     "low_priority_core_quota",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("LowPriorityCoreQuota"),
			},
			{
				Name:     "dedicated_core_quota_per_vm_family",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DedicatedCoreQuotaPerVMFamily"),
			},
			{
				Name:     "dedicated_core_quota_per_vm_family_enforced",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("DedicatedCoreQuotaPerVMFamilyEnforced"),
			},
			{
				Name:     "pool_quota",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("PoolQuota"),
			},
			{
				Name:     "active_job_and_job_schedule_quota",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ActiveJobAndJobScheduleQuota"),
			},
			{
				Name:     "allowed_authentication_modes",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("AllowedAuthenticationModes"),
			},
			{
				Name:     "identity",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Identity"),
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
			{
				Name:     "location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Location"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
		},
	}
}

func fetchBatchAccounts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Batch.Accounts

	response, err := svc.List(ctx)

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
