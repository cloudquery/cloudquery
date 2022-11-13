// Auto generated code - DO NOT EDIT.

package sql

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql"
)

func managedDatabases() *schema.Table {
	return &schema.Table{
		Name:        "azure_sql_managed_databases",
		Description: `https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql#ManagedDatabase`,
		Resolver:    fetchSQLManagedDatabases,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "sql_managed_instance_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
			},
			{
				Name:     "collation",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Collation"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "creation_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreationDate"),
			},
			{
				Name:     "earliest_restore_point",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("EarliestRestorePoint"),
			},
			{
				Name:     "restore_point_in_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("RestorePointInTime"),
			},
			{
				Name:     "default_secondary_location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DefaultSecondaryLocation"),
			},
			{
				Name:     "catalog_collation",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CatalogCollation"),
			},
			{
				Name:     "create_mode",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CreateMode"),
			},
			{
				Name:     "storage_container_uri",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StorageContainerURI"),
			},
			{
				Name:     "source_database_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SourceDatabaseID"),
			},
			{
				Name:     "restorable_dropped_database_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RestorableDroppedDatabaseID"),
			},
			{
				Name:     "storage_container_sas_token",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StorageContainerSasToken"),
			},
			{
				Name:     "failover_group_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FailoverGroupID"),
			},
			{
				Name:     "recoverable_database_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RecoverableDatabaseID"),
			},
			{
				Name:     "long_term_retention_backup_resource_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LongTermRetentionBackupResourceID"),
			},
			{
				Name:     "auto_complete_restore",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AutoCompleteRestore"),
			},
			{
				Name:     "last_backup_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LastBackupName"),
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

		Relations: []*schema.Table{
			managedDatabaseVulnerabilityAssessments(),
			managedDatabaseVulnerabilityAssessmentScans(),
		},
	}
}

func fetchSQLManagedDatabases(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().SQL.ManagedDatabases

	instance := parent.Item.(sql.ManagedInstance)
	resourceDetails, err := client.ParseResourceID(*instance.ID)
	if err != nil {
		return err
	}
	response, err := svc.ListByInstance(ctx, resourceDetails.ResourceGroup, *instance.Name)

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
