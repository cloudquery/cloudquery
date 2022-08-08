package datalake

import (
	"context"
	"fmt"
	"net"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/datalake/analytics/mgmt/account"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func AnalyticsAccounts() *schema.Table {
	return &schema.Table{
		Name:         "azure_datalake_analytics_accounts",
		Description:  "Data Lake Analytics account",
		Resolver:     fetchDatalakeAnalyticsAccounts,
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
				Name:        "default_data_lake_store_account",
				Description: "The default Data Lake Store account associated with this account",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DataLakeAnalyticsAccountProperties.DefaultDataLakeStoreAccount"),
			},
			{
				Name:        "firewall_state",
				Description: "The current state of the IP address firewall for this account",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DataLakeAnalyticsAccountProperties.FirewallState"),
			},
			{
				Name:        "firewall_allow_azure_ips",
				Description: "The current state of allowing or disallowing IPs originating within Azure through the firewall",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DataLakeAnalyticsAccountProperties.FirewallAllowAzureIps"),
			},
			{
				Name:        "new_tier",
				Description: "The commitment tier for the next month",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DataLakeAnalyticsAccountProperties.NewTier"),
			},
			{
				Name:        "current_tier",
				Description: "The commitment tier in use for the current month",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DataLakeAnalyticsAccountProperties.CurrentTier"),
			},
			{
				Name:        "max_job_count",
				Description: "The maximum supported jobs running under the account at the same time",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("DataLakeAnalyticsAccountProperties.MaxJobCount"),
			},
			{
				Name:        "system_max_job_count",
				Description: "The system defined maximum supported jobs running under the account at the same time, which restricts the maximum number of running jobs the user can set for the account",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("DataLakeAnalyticsAccountProperties.SystemMaxJobCount"),
			},
			{
				Name:        "max_degree_of_parallelism",
				Description: "The maximum supported degree of parallelism for this account",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("DataLakeAnalyticsAccountProperties.MaxDegreeOfParallelism"),
			},
			{
				Name:        "system_max_degree_of_parallelism",
				Description: "The system defined maximum supported degree of parallelism for this account, which restricts the maximum value of parallelism the user can set for the account",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("DataLakeAnalyticsAccountProperties.SystemMaxDegreeOfParallelism"),
			},
			{
				Name:        "max_degree_of_parallelism_per_job",
				Description: "The maximum supported degree of parallelism per job for this account",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("DataLakeAnalyticsAccountProperties.MaxDegreeOfParallelismPerJob"),
			},
			{
				Name:        "min_priority_per_job",
				Description: "The minimum supported priority per job for this account",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("DataLakeAnalyticsAccountProperties.MinPriorityPerJob"),
			},
			{
				Name:        "query_store_retention",
				Description: "The number of days that job metadata is retained",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("DataLakeAnalyticsAccountProperties.QueryStoreRetention"),
			},
			{
				Name:        "account_id",
				Description: "The unique identifier associated with this Data Lake Analytics account",
				Type:        schema.TypeUUID,
				Resolver:    schema.PathResolver("DataLakeAnalyticsAccountProperties.AccountID"),
			},
			{
				Name:        "provisioning_state",
				Description: "The provisioning status of the Data Lake Analytics account",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DataLakeAnalyticsAccountProperties.ProvisioningState"),
			},
			{
				Name:        "state",
				Description: "The state of the Data Lake Analytics account",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DataLakeAnalyticsAccountProperties.State"),
			},
			{
				Name:     "creation_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("DataLakeAnalyticsAccountProperties.CreationTime.Time"),
			},
			{
				Name:     "last_modified_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("DataLakeAnalyticsAccountProperties.LastModifiedTime.Time"),
			},
			{
				Name:        "endpoint",
				Description: "The full CName endpoint for this account",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DataLakeAnalyticsAccountProperties.Endpoint"),
			},
			{
				Name:        "id",
				Description: "The resource identifier",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "The resource name",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "The resource type",
				Type:        schema.TypeString,
			},
			{
				Name:        "location",
				Description: "The resource location",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "The resource tags",
				Type:        schema.TypeJSON,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "azure_datalake_analytics_account_data_lake_store_accounts",
				Description: "DataLakeStoreAccountInformation data Lake Store account information",
				Resolver:    fetchDatalakeAnalyticsAccountDataLakeStoreAccounts,
				Columns: []schema.Column{
					{
						Name:        "analytics_account_cq_id",
						Description: "Unique CloudQuery ID of azure_datalake_analytics_accounts table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "suffix",
						Description: "The optional suffix for the Data Lake Store account",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DataLakeStoreAccountInformationProperties.Suffix"),
					},
					{
						Name:        "id",
						Description: "The resource identifier",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
						// This looks like a deprecated field, always returns nil
						// we might want to delete it in the future
						IgnoreInTests: true,
					},
					{
						Name:        "name",
						Description: "The resource name",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "The resource type",
						Type:        schema.TypeString,
						// This looks like a deprecated field, always returns nil
						// we might want to delete it in the future
						IgnoreInTests: true,
					},
				},
			},
			{
				Name:          "azure_datalake_analytics_account_storage_accounts",
				Description:   "StorageAccountInformation azure Storage account information",
				Resolver:      fetchDatalakeAnalyticsAccountStorageAccounts,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "analytics_account_cq_id",
						Description: "Unique CloudQuery ID of azure_datalake_analytics_accounts table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "suffix",
						Description: "The optional suffix for the storage account",
						Type:        schema.TypeString,
					},
					{
						Name:        "id",
						Description: "The resource identifier",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
					{
						Name:        "name",
						Description: "The resource name",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "The resource type",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "azure_datalake_analytics_account_compute_policies",
				Description: "ComputePolicy data Lake Analytics compute policy information",
				Resolver:    fetchDatalakeAnalyticsAccountComputePolicies,
				// Not sure if it's possible to create those via terraform
				// https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/data_lake_analytics_account
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "analytics_account_cq_id",
						Description: "Unique CloudQuery ID of azure_datalake_analytics_accounts table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "object_id",
						Description: "The AAD object identifier for the entity to create a policy for",
						Type:        schema.TypeUUID,
						Resolver:    schema.PathResolver("ComputePolicyProperties.ObjectID"),
					},
					{
						Name:        "object_type",
						Description: "The type of AAD object the object identifier refers to",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ComputePolicyProperties.ObjectType"),
					},
					{
						Name:        "max_degree_of_parallelism_per_job",
						Description: "The maximum degree of parallelism per job this user can use to submit jobs",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("ComputePolicyProperties.MaxDegreeOfParallelismPerJob"),
					},
					{
						Name:        "min_priority_per_job",
						Description: "The minimum priority per job this user can use to submit jobs",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("ComputePolicyProperties.MinPriorityPerJob"),
					},
					{
						Name:        "id",
						Description: "The resource identifier",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
						// This looks like a deprecated field, always returns nil
						// we might want to delete it in the future
						IgnoreInTests: true,
					},
					{
						Name:        "name",
						Description: "The resource name",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "The resource type",
						Type:        schema.TypeString,
						// This looks like a deprecated field, always returns nil
						// we might want to delete it in the future
						IgnoreInTests: true,
					},
				},
			},
			{
				Name:        "azure_datalake_analytics_account_firewall_rules",
				Description: "FirewallRule data Lake Analytics firewall rule information",
				Resolver:    fetchDatalakeAnalyticsAccountFirewallRules,
				Columns: []schema.Column{
					{
						Name:        "analytics_account_cq_id",
						Description: "Unique CloudQuery ID of azure_datalake_analytics_accounts table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "start_ip_address",
						Description: "The start IP address for the firewall rule",
						Type:        schema.TypeInet,
						Resolver:    resolveAnalyticsAccountFirewallRulesStartIpAddress,
					},
					{
						Name:        "end_ip_address",
						Description: "The end IP address for the firewall rule",
						Type:        schema.TypeInet,
						Resolver:    resolveAnalyticsAccountFirewallRulesEndIpAddress,
					},
					{
						Name:        "id",
						Description: "The resource identifier",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
						// This column is deprecated and empty
						IgnoreInTests: true,
					},
					{
						Name:        "name",
						Description: "The resource name",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "The resource type",
						Type:        schema.TypeString,
						// This column is deprecated and empty
						IgnoreInTests: true,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchDatalakeAnalyticsAccounts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().DataLake.DataLakeAnalyticsAccounts
	result, err := svc.List(ctx, "", nil, nil, "", "", nil)
	if err != nil {
		return diag.WrapError(err)
	}
	for result.NotDone() {
		accounts := result.Values()
		for _, a := range accounts {
			resourceDetails, err := client.ParseResourceID(*a.ID)
			if err != nil {
				return diag.WrapError(err)
			}
			result, err := svc.Get(ctx, resourceDetails.ResourceGroup, *a.Name)
			if err != nil {
				return diag.WrapError(err)
			}
			res <- result
		}

		if err := result.NextWithContext(ctx); err != nil {
			return diag.WrapError(err)
		}
	}
	return nil
}
func fetchDatalakeAnalyticsAccountDataLakeStoreAccounts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(account.DataLakeAnalyticsAccount)
	if p.DataLakeStoreAccounts != nil {
		res <- *p.DataLakeStoreAccounts
	}

	return nil
}
func fetchDatalakeAnalyticsAccountStorageAccounts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(account.DataLakeAnalyticsAccount)
	if p.StorageAccounts != nil {
		res <- *p.StorageAccounts
	}

	return nil
}
func fetchDatalakeAnalyticsAccountComputePolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(account.DataLakeAnalyticsAccount)
	if p.ComputePolicies != nil {
		res <- *p.ComputePolicies
	}

	return nil
}
func fetchDatalakeAnalyticsAccountFirewallRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(account.DataLakeAnalyticsAccount)
	if p.FirewallRules != nil {
		res <- *p.FirewallRules
	}

	return nil
}
func resolveAnalyticsAccountFirewallRulesStartIpAddress(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(account.FirewallRule)
	i := net.ParseIP(*p.StartIPAddress)
	if i == nil {
		return diag.WrapError(fmt.Errorf("wrong format of IP: %s", *p.StartIPAddress))
	}
	return diag.WrapError(resource.Set(c.Name, i))
}
func resolveAnalyticsAccountFirewallRulesEndIpAddress(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(account.FirewallRule)
	i := net.ParseIP(*p.EndIPAddress)
	if i == nil {
		return diag.WrapError(fmt.Errorf("wrong format of IP: %s", *p.EndIPAddress))
	}
	return diag.WrapError(resource.Set(c.Name, i))
}
