// Auto generated code - DO NOT EDIT.

package datalake

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/datalake/analytics/mgmt/account"
)

func AnalyticsAccounts() *schema.Table {
	return &schema.Table{
		Name:                "azure_datalake_analytics_accounts",
		Description:         `https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/datalake/analytics/mgmt/2016-11-01/account#DataLakeAnalyticsAccount`,
		Resolver:            fetchDataLakeAnalyticsAccounts,
		PreResourceResolver: getDataLakeAnalyticsAccount,
		Multiplex:           client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "default_data_lake_store_account",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DefaultDataLakeStoreAccount"),
			},
			{
				Name:     "data_lake_store_accounts",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DataLakeStoreAccounts"),
			},
			{
				Name:     "storage_accounts",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("StorageAccounts"),
			},
			{
				Name:     "compute_policies",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ComputePolicies"),
			},
			{
				Name:     "firewall_rules",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("FirewallRules"),
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
				Name:     "max_job_count",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MaxJobCount"),
			},
			{
				Name:     "system_max_job_count",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("SystemMaxJobCount"),
			},
			{
				Name:     "max_degree_of_parallelism",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MaxDegreeOfParallelism"),
			},
			{
				Name:     "system_max_degree_of_parallelism",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("SystemMaxDegreeOfParallelism"),
			},
			{
				Name:     "max_degree_of_parallelism_per_job",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MaxDegreeOfParallelismPerJob"),
			},
			{
				Name:     "min_priority_per_job",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MinPriorityPerJob"),
			},
			{
				Name:     "query_store_retention",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("QueryStoreRetention"),
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

func fetchDataLakeAnalyticsAccounts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().DataLake.AnalyticsAccounts

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

func getDataLakeAnalyticsAccount(ctx context.Context, meta schema.ClientMeta, r *schema.Resource) error {
	svc := meta.(*client.Client).Services().DataLake.AnalyticsAccounts

	account := r.Item.(account.DataLakeAnalyticsAccountBasic)
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
