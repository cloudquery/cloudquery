// Auto generated code - DO NOT EDIT.

package sql

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql"
)

func databases() *schema.Table {
	return &schema.Table{
		Name:        "azure_sql_databases",
		Description: `https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql#Database`,
		Resolver:    fetchSQLDatabases,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "sql_server_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
			},
			{
				Name:     "sku",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Sku"),
			},
			{
				Name:     "kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Kind"),
			},
			{
				Name:     "managed_by",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ManagedBy"),
			},
			{
				Name:     "create_mode",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CreateMode"),
			},
			{
				Name:     "collation",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Collation"),
			},
			{
				Name:     "max_size_bytes",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MaxSizeBytes"),
			},
			{
				Name:     "sample_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SampleName"),
			},
			{
				Name:     "elastic_pool_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ElasticPoolID"),
			},
			{
				Name:     "source_database_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SourceDatabaseID"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "database_id",
				Type:     schema.TypeUUID,
				Resolver: schema.PathResolver("DatabaseID"),
			},
			{
				Name:     "creation_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreationDate"),
			},
			{
				Name:     "current_service_objective_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CurrentServiceObjectiveName"),
			},
			{
				Name:     "requested_service_objective_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RequestedServiceObjectiveName"),
			},
			{
				Name:     "default_secondary_location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DefaultSecondaryLocation"),
			},
			{
				Name:     "failover_group_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FailoverGroupID"),
			},
			{
				Name:     "restore_point_in_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("RestorePointInTime"),
			},
			{
				Name:     "source_database_deletion_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("SourceDatabaseDeletionDate"),
			},
			{
				Name:     "recovery_services_recovery_point_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RecoveryServicesRecoveryPointID"),
			},
			{
				Name:     "long_term_retention_backup_resource_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LongTermRetentionBackupResourceID"),
			},
			{
				Name:     "recoverable_database_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RecoverableDatabaseID"),
			},
			{
				Name:     "restorable_dropped_database_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RestorableDroppedDatabaseID"),
			},
			{
				Name:     "catalog_collation",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CatalogCollation"),
			},
			{
				Name:     "zone_redundant",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ZoneRedundant"),
			},
			{
				Name:     "license_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LicenseType"),
			},
			{
				Name:     "max_log_size_bytes",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MaxLogSizeBytes"),
			},
			{
				Name:     "earliest_restore_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("EarliestRestoreDate"),
			},
			{
				Name:     "read_scale",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReadScale"),
			},
			{
				Name:     "high_availability_replica_count",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("HighAvailabilityReplicaCount"),
			},
			{
				Name:     "secondary_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SecondaryType"),
			},
			{
				Name:     "current_sku",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CurrentSku"),
			},
			{
				Name:     "auto_pause_delay",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("AutoPauseDelay"),
			},
			{
				Name:     "storage_account_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StorageAccountType"),
			},
			{
				Name:     "min_capacity",
				Type:     schema.TypeFloat,
				Resolver: schema.PathResolver("MinCapacity"),
			},
			{
				Name:     "paused_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("PausedDate"),
			},
			{
				Name:     "resumed_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("ResumedDate"),
			},
			{
				Name:     "maintenance_configuration_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MaintenanceConfigurationID"),
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
			databaseBlobAuditingPolicies(),
			databaseVulnerabilityAssessments(),
			databaseVulnerabilityAssessmentScans(),
			backupLongTermRetentionPolicies(),
			databaseThreatDetectionPolicies(),
			transparentDataEncryptions(),
		},
	}
}

func fetchSQLDatabases(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().SQL.Databases

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
