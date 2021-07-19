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
				Description: "Specifies the mode of database creation.",
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
				Description: "The name of the sample schema to apply when creating this database.",
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
				Description: "The status of the database.",
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
				Description: "Collation of the metadata catalog.",
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
				Description: "The license type to apply for this database.",
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
				Description: "The state of read-only routing If enabled, connections that have application intent set to readonly in their connection string may be routed to a readonly secondary replica in the same region.",
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
				Description: "The secondary type of the database if it is a secondary.",
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
				Description: "Time in minutes after which database is automatically paused.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("DatabaseProperties.AutoPauseDelay"),
			},
			{
				Name:        "storage_account_type",
				Description: "The storage account type used to store backups for this database.",
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
				Description: "Maintenance configuration id assigned to the database.",
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
		Relations: []*schema.Table{
			{
				Name:        "azure_sql_database_db_blob_auditing_policies",
				Description: "Database blob auditing policy",
				Resolver:    fetchSqlDatabaseDbBlobAuditingPolicies,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"database_cq_id", "id"}},
				Columns: []schema.Column{
					{
						Name:        "database_cq_id",
						Description: "Unique ID of azure_sql_databases table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "kind",
						Description: "Resource kind",
						Type:        schema.TypeString,
					},
					{
						Name:        "state",
						Description: "Specifies the state of the policy.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DatabaseBlobAuditingPolicyProperties.State"),
					},
					{
						Name:        "storage_endpoint",
						Description: "Specifies the blob storage endpoint.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DatabaseBlobAuditingPolicyProperties.StorageEndpoint"),
					},
					{
						Name:        "storage_account_access_key",
						Description: "Specifies the identifier key of the auditing storage account.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DatabaseBlobAuditingPolicyProperties.StorageAccountAccessKey"),
					},
					{
						Name:        "retention_days",
						Description: "Specifies the number of days to keep in the audit logs in the storage account",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("DatabaseBlobAuditingPolicyProperties.RetentionDays"),
					},
					{
						Name:        "audit_actions_and_groups",
						Description: "Specifies the Actions-Groups and Actions to audit.",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("DatabaseBlobAuditingPolicyProperties.AuditActionsAndGroups"),
					},
					{
						Name:        "storage_account_subscription_id",
						Description: "Specifies the blob storage subscription Id",
						Type:        schema.TypeUUID,
						Resolver:    schema.PathResolver("DatabaseBlobAuditingPolicyProperties.StorageAccountSubscriptionID"),
					},
					{
						Name:        "is_storage_secondary_key_in_use",
						Description: "Specifies whether storageAccountAccessKey value is the storage's secondary key",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("DatabaseBlobAuditingPolicyProperties.IsStorageSecondaryKeyInUse"),
					},
					{
						Name:        "is_azure_monitor_target_enabled",
						Description: "Specifies whether audit events are sent to Azure Monitor.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("DatabaseBlobAuditingPolicyProperties.IsAzureMonitorTargetEnabled"),
					},
					{
						Name:        "queue_delay_ms",
						Description: "Specifies the amount of time in milliseconds that can elapse before audit actions are forced to be processed The default minimum value is 1000 (1 second) The maximum is 2,147,483,647",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("DatabaseBlobAuditingPolicyProperties.QueueDelayMs"),
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

func fetchSqlDatabaseDbBlobAuditingPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	svc := meta.(*client.Client).Services().SQL.DatabaseBlobAuditingPolicies
	database := parent.Item.(sql.Database)
	details, err := client.ParseResourceID(*database.ID)
	if err != nil {
		return err
	}
	serverName := parent.Parent.Get("name").(*string)
	result, err := svc.ListByDatabase(ctx, details.ResourceGroup, *serverName, *database.Name)
	if err != nil {
		return err
	}
	for result.NotDone() {
		res <- result.Values()
		if err := result.NextWithContext(ctx); err != nil {
			return err
		}
	}
	return nil
}
