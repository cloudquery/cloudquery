// Auto generated code - DO NOT EDIT.

package datalake

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/datalake/store/mgmt/account"
)

func StoreAccounts() *schema.Table {
	return &schema.Table{
		Name:                "azure_datalake_store_accounts",
		Description:         `https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/datalake/store/mgmt/2016-11-01/account#DataLakeStoreAccount`,
		Resolver:            fetchDataLakeStoreAccounts,
		PreResourceResolver: getDataLakeStoreAccount,
		Multiplex:           client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "identity",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Identity"),
			},
			{
				Name:     "default_group",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DefaultGroup"),
			},
			{
				Name:     "encryption_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EncryptionConfig"),
			},
			{
				Name:     "encryption_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EncryptionState"),
			},
			{
				Name:     "encryption_provisioning_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EncryptionProvisioningState"),
			},
			{
				Name:     "firewall_rules",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("FirewallRules"),
			},
			{
				Name:     "virtual_network_rules",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("VirtualNetworkRules"),
			},
			{
				Name:     "firewall_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FirewallState"),
			},
			{
				Name:     "firewall_allow_azure_ips",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FirewallAllowAzureIps"),
			},
			{
				Name:     "trusted_id_providers",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("TrustedIDProviders"),
			},
			{
				Name:     "trusted_id_provider_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TrustedIDProviderState"),
			},
			{
				Name:     "new_tier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("NewTier"),
			},
			{
				Name:     "current_tier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CurrentTier"),
			},
			{
				Name:     "account_id",
				Type:     schema.TypeUUID,
				Resolver: schema.PathResolver("AccountID"),
			},
			{
				Name:     "provisioning_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ProvisioningState"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "creation_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreationTime"),
			},
			{
				Name:     "last_modified_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastModifiedTime"),
			},
			{
				Name:     "endpoint",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Endpoint"),
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

func fetchDataLakeStoreAccounts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().DataLake.StoreAccounts

	response, err := svc.List(ctx, "", nil, nil, "", "", nil)

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

func getDataLakeStoreAccount(ctx context.Context, meta schema.ClientMeta, r *schema.Resource) error {
	svc := meta.(*client.Client).Services().DataLake.StoreAccounts

	account := r.Item.(account.DataLakeStoreAccountBasic)
	resourceDetails, err := client.ParseResourceID(*account.ID)
	if err != nil {
		return err
	}
	item, err := svc.Get(ctx, resourceDetails.ResourceGroup, *account.Name)
	if err != nil {
		return err
	}
	r.SetItem(item)
	return nil
}
