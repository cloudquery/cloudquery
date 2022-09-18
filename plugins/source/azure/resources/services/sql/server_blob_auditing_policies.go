// Auto generated code - DO NOT EDIT.

package sql

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql"
)

func serverBlobAuditingPolicies() *schema.Table {
	return &schema.Table{
		Name:     "azure_sql_server_blob_auditing_policies",
		Resolver: fetchSQLServerBlobAuditingPolicies,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "sql_server_id",
				Type:     schema.TypeUUID,
				Resolver: schema.ParentIDResolver,
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

func fetchSQLServerBlobAuditingPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().SQL.ServerBlobAuditingPolicies

	server := parent.Item.(sql.Server)
	resourceDetails, err := client.ParseResourceID(*server.ID)
	if err != nil {
		return err
	}
	response, err := svc.ListByServer(ctx, resourceDetails.ResourceGroup, *server.Name)

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
