// Auto generated code - DO NOT EDIT.

package sql

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/sql/mgmt/sql"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/pkg/errors"
)

func databaseBlobAuditingPolicies() *schema.Table {
	return &schema.Table{
		Name:     "azure_sql_database_blob_auditing_policies",
		Resolver: fetchSQLDatabaseBlobAuditingPolicies,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "cq_id_parent",
				Type:     schema.TypeUUID,
				Resolver: schema.ParentIdResolver,
			},
			{
				Name:     "kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Kind"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "storage_endpoint",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StorageEndpoint"),
			},
			{
				Name:     "storage_account_access_key",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StorageAccountAccessKey"),
			},
			{
				Name:     "retention_days",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("RetentionDays"),
			},
			{
				Name:     "audit_actions_and_groups",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("AuditActionsAndGroups"),
			},
			{
				Name:     "is_storage_secondary_key_in_use",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsStorageSecondaryKeyInUse"),
			},
			{
				Name:     "is_azure_monitor_target_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsAzureMonitorTargetEnabled"),
			},
			{
				Name:     "queue_delay_ms",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("QueueDelayMs"),
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
		},
	}
}

func fetchSQLDatabaseBlobAuditingPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().SQL.DatabaseBlobAuditingPolicies

	server := parent.Parent.Item.(sql.Server)
	database := parent.Item.(sql.Database)
	resourceDetails, err := client.ParseResourceID(*database.ID)
	if err != nil {
		return errors.WithStack(err)
	}
	response, err := svc.ListByDatabase(ctx, resourceDetails.ResourceGroup, *server.Name, *database.Name)

	if err != nil {
		return errors.WithStack(err)
	}

	for response.NotDone() {
		res <- response.Values()
		if err := response.NextWithContext(ctx); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}
