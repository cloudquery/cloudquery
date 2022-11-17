// Code generated by codegen; DO NOT EDIT.

package sql

import (
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func databaseBlobAuditingPolicies() *schema.Table {
	return &schema.Table{
		Name:        "azure_sql_database_blob_auditing_policies",
		Description: `https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql#DatabaseBlobAuditingPolicy`,
		Resolver:    fetchDatabaseBlobAuditingPolicies,
		Columns: []schema.Column{
			{
				Name:        "subscription_id",
				Type:        schema.TypeString,
				Resolver:    client.SubscriptionIDResolver,
				Description: `Azure subscription ID`,
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.State"),
			},
			{
				Name:     "audit_actions_and_groups",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Properties.AuditActionsAndGroups"),
			},
			{
				Name:     "is_azure_monitor_target_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Properties.IsAzureMonitorTargetEnabled"),
			},
			{
				Name:     "is_managed_identity_in_use",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Properties.IsManagedIdentityInUse"),
			},
			{
				Name:     "is_storage_secondary_key_in_use",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Properties.IsStorageSecondaryKeyInUse"),
			},
			{
				Name:     "queue_delay_ms",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Properties.QueueDelayMs"),
			},
			{
				Name:     "retention_days",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Properties.RetentionDays"),
			},
			{
				Name:     "storage_account_access_key",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.StorageAccountAccessKey"),
			},
			{
				Name:     "storage_account_subscription_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.StorageAccountSubscriptionID"),
			},
			{
				Name:     "storage_endpoint",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.StorageEndpoint"),
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
				Name:     "kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Kind"),
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
				Name:     "database_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
			},
		},
	}
}
