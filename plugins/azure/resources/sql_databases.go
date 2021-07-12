package resources

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func SQLDatabases() *schema.Table {
	return &schema.Table{
		Name:         "azure_sql_databases",
		Description:  "Azure sql database",
		Resolver:     fetchSqlDatabases,
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
				Name:        "sku_name",
				Description: "The name of the SKU, typically, a letter + Number code, eg P3",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Sku.Name"),
			},
			{
				Name:        "sku_tier",
				Description: "The tier or edition of the particular SKU, eg Basic, Premium",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Sku.Tier"),
			},
			{
				Name:        "sku_size",
				Description: "Size of the particular SKU",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Sku.Size"),
			},
			{
				Name:        "sku_family",
				Description: "If the service has different generations of hardware, for the same SKU, then that can be captured here",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Sku.Family"),
			},
			{
				Name:        "sku_capacity",
				Description: "Capacity of the particular SKU",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Sku.Capacity"),
			},
			{
				Name:        "kind",
				Description: "Kind of database This is metadata used for the Azure portal experience",
				Type:        schema.TypeString,
			},
			{
				Name:        "managed_by",
				Description: "Resource that manages the database",
				Type:        schema.TypeString,
			},
			{
				Name:        "create_mode",
				Description: "Specifies the mode of database creation  Default: regular database creation  Copy: creates a database as a copy of an existing database sourceDatabaseId must be specified as the resource ID of the source database  Secondary: creates a database as a secondary replica of an existing database sourceDatabaseId must be specified as the resource ID of the existing primary database  PointInTimeRestore: Creates a database by restoring a point in time backup of an existing database sourceDatabaseId must be specified as the resource ID of the existing database, and restorePointInTime must be specified  Recovery: Creates a database by restoring a geo-replicated backup sourceDatabaseId must be specified as the recoverable database resource ID to restore  Restore: Creates a database by restoring a backup of a deleted database sourceDatabaseId must be specified If sourceDatabaseId is the database's original resource ID, then sourceDatabaseDeletionDate must be specified Otherwise sourceDatabaseId must be the restorable dropped database resource ID and sourceDatabaseDeletionDate is ignored restorePointInTime may also be specified to restore from an earlier point in time  RestoreLongTermRetentionBackup: Creates a database by restoring from a long term retention vault recoveryServicesRecoveryPointResourceId must be specified as the recovery point resource ID  Copy, Secondary, and RestoreLongTermRetentionBackup are not supported for DataWarehouse edition Possible values include: 'CreateModeDefault', 'CreateModeCopy', 'CreateModeSecondary', 'CreateModePointInTimeRestore', 'CreateModeRestore', 'CreateModeRecovery', 'CreateModeRestoreExternalBackup', 'CreateModeRestoreExternalBackupSecondary', 'CreateModeRestoreLongTermRetentionBackup', 'CreateModeOnlineSecondary'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DatabaseProperties.CreateMode"),
			},
			{
				Name:        "collation",
				Description: "The collation of the database",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DatabaseProperties.Collation"),
			},
			{
				Name:        "max_size_bytes",
				Description: "The max size of the database expressed in bytes",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("DatabaseProperties.MaxSizeBytes"),
			},
			{
				Name:        "sample_name",
				Description: "The name of the sample schema to apply when creating this database Possible values include: 'AdventureWorksLT', 'WideWorldImportersStd', 'WideWorldImportersFull'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DatabaseProperties.SampleName"),
			},
			{
				Name:        "elastic_pool_id",
				Description: "The resource identifier of the elastic pool containing this database",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DatabaseProperties.ElasticPoolID"),
			},
			{
				Name:        "source_database_id",
				Description: "The resource identifier of the source database associated with create operation of this database",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DatabaseProperties.SourceDatabaseID"),
			},
			{
				Name:        "status",
				Description: "The status of the database Possible values include: 'DatabaseStatusOnline', 'DatabaseStatusRestoring', 'DatabaseStatusRecoveryPending', 'DatabaseStatusRecovering', 'DatabaseStatusSuspect', 'DatabaseStatusOffline', 'DatabaseStatusStandby', 'DatabaseStatusShutdown', 'DatabaseStatusEmergencyMode', 'DatabaseStatusAutoClosed', 'DatabaseStatusCopying', 'DatabaseStatusCreating', 'DatabaseStatusInaccessible', 'DatabaseStatusOfflineSecondary', 'DatabaseStatusPausing', 'DatabaseStatusPaused', 'DatabaseStatusResuming', 'DatabaseStatusScaling', 'DatabaseStatusOfflineChangingDwPerformanceTiers', 'DatabaseStatusOnlineChangingDwPerformanceTiers', 'DatabaseStatusDisabled'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DatabaseProperties.Status"),
			},
			{
				Name:        "database_id",
				Description: "The ID of the database",
				Type:        schema.TypeUUID,
				Resolver:    schema.PathResolver("DatabaseProperties.DatabaseID"),
			},
			{
				Name:        "creation_date_time",
				Description: "The creation date of the database (ISO8601 format)",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("DatabaseProperties.CreationDate.Time"),
			},
			{
				Name:        "current_service_objective_name",
				Description: "The current service level objective name of the database",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DatabaseProperties.CurrentServiceObjectiveName"),
			},
			{
				Name:        "requested_service_objective_name",
				Description: "The requested service level objective name of the database",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DatabaseProperties.RequestedServiceObjectiveName"),
			},
			{
				Name:        "default_secondary_location",
				Description: "The default secondary region for this database",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DatabaseProperties.DefaultSecondaryLocation"),
			},
			{
				Name:        "failover_group_id",
				Description: "Failover Group resource identifier that this database belongs to",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DatabaseProperties.FailoverGroupID"),
			},
			{
				Name:        "restore_point_in_time",
				Description: "Specifies the point in time (ISO8601 format) of the source database that will be restored to create the new database",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("DatabaseProperties.RestorePointInTime.Time"),
			},
			{
				Name:        "source_database_deletion_date_time",
				Description: "Specifies the time that the database was deleted",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("DatabaseProperties.SourceDatabaseDeletionDate.Time"),
			},
			{
				Name:        "recovery_services_recovery_point_id",
				Description: "The resource identifier of the recovery point associated with create operation of this database",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DatabaseProperties.RecoveryServicesRecoveryPointID"),
			},
			{
				Name:        "long_term_retention_backup_resource_id",
				Description: "The resource identifier of the long term retention backup associated with create operation of this database",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DatabaseProperties.LongTermRetentionBackupResourceID"),
			},
			{
				Name:        "recoverable_database_id",
				Description: "The resource identifier of the recoverable database associated with create operation of this database",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DatabaseProperties.RecoverableDatabaseID"),
			},
			{
				Name:        "restorable_dropped_database_id",
				Description: "The resource identifier of the restorable dropped database associated with create operation of this database",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DatabaseProperties.RestorableDroppedDatabaseID"),
			},
			{
				Name:        "catalog_collation",
				Description: "Collation of the metadata catalog Possible values include: 'DATABASEDEFAULT', 'SQLLatin1GeneralCP1CIAS'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DatabaseProperties.CatalogCollation"),
			},
			{
				Name:        "zone_redundant",
				Description: "Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("DatabaseProperties.ZoneRedundant"),
			},
			{
				Name:        "license_type",
				Description: "The license type to apply for this database `LicenseIncluded` if you need a license, or `BasePrice` if you have a license and are eligible for the Azure Hybrid Benefit Possible values include: 'LicenseIncluded', 'BasePrice'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DatabaseProperties.LicenseType"),
			},
			{
				Name:        "max_log_size_bytes",
				Description: "The max log size for this database",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("DatabaseProperties.MaxLogSizeBytes"),
			},
			{
				Name:        "earliest_restore_date_time",
				Description: "This records the earliest start date and time that restore is available for this database (ISO8601 format)",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("DatabaseProperties.EarliestRestoreDate.Time"),
			},
			{
				Name:        "read_scale",
				Description: "The state of read-only routing If enabled, connections that have application intent set to readonly in their connection string may be routed to a readonly secondary replica in the same region Possible values include: 'DatabaseReadScaleEnabled', 'DatabaseReadScaleDisabled'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DatabaseProperties.ReadScale"),
			},
			{
				Name:        "high_availability_replica_count",
				Description: "The number of secondary replicas associated with the database that are used to provide high availability",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("DatabaseProperties.HighAvailabilityReplicaCount"),
			},
			{
				Name:        "secondary_type",
				Description: "The secondary type of the database if it is a secondary  Valid values are Geo and Named Possible values include: 'Geo', 'Named'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DatabaseProperties.SecondaryType"),
			},
			{
				Name:        "current_sku_name",
				Description: "The name of the SKU, typically, a letter + Number code, eg P3",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DatabaseProperties.CurrentSku.Name"),
			},
			{
				Name:        "current_sku_tier",
				Description: "The tier or edition of the particular SKU, eg Basic, Premium",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DatabaseProperties.CurrentSku.Tier"),
			},
			{
				Name:        "current_sku_size",
				Description: "Size of the particular SKU",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DatabaseProperties.CurrentSku.Size"),
			},
			{
				Name:        "current_sku_family",
				Description: "If the service has different generations of hardware, for the same SKU, then that can be captured here",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DatabaseProperties.CurrentSku.Family"),
			},
			{
				Name:        "current_sku_capacity",
				Description: "Capacity of the particular SKU",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("DatabaseProperties.CurrentSku.Capacity"),
			},
			{
				Name:        "auto_pause_delay",
				Description: "Time in minutes after which database is automatically paused A value of -1 means that automatic pause is disabled",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("DatabaseProperties.AutoPauseDelay"),
			},
			{
				Name:        "storage_account_type",
				Description: "The storage account type used to store backups for this database Possible values include: 'GRS', 'LRS', 'ZRS'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DatabaseProperties.StorageAccountType"),
			},
			{
				Name:        "min_capacity",
				Description: "Minimal capacity that database will always have allocated, if not paused",
				Type:        schema.TypeFloat,
				Resolver:    schema.PathResolver("DatabaseProperties.MinCapacity"),
			},
			{
				Name:     "paused_date_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("DatabaseProperties.PausedDate.Time"),
			},
			{
				Name:     "resumed_date_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("DatabaseProperties.ResumedDate.Time"),
			},
			{
				Name:        "maintenance_configuration_id",
				Description: "Maintenance configuration id assigned to the database This configuration defines the period when the maintenance updates will occur",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DatabaseProperties.MaintenanceConfigurationID"),
			},
			{
				Name:        "location",
				Description: "Resource location",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "Resource tags",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "id",
				Description: "Resource ID",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "Resource name",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "Resource type",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchSqlDatabases(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	svc := meta.(*client.Client).Services().SQL.Databases
	server := parent.Item.(sql.Server)
	resourceDetails, err := client.ParseResourceID(*server.ID)
	if err != nil {
		return err
	}
	databases, err := svc.ListByServer(ctx, resourceDetails.ResourceGroup, *server.Name)
	if err != nil {
		return err
	}
	for databases.NotDone() {
		res <- databases.Values()
		if err := databases.NextWithContext(ctx); err != nil {
			return err
		}
	}
	return nil
}
