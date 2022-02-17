package sql

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func SqlDatabases() *schema.Table {
	return &schema.Table{
		Name:        "azure_sql_databases",
		Description: "Azure sql database",
		Resolver:    fetchSqlDatabases,
		Options:     schema.TableCreationOptions{PrimaryKeys: []string{"server_cq_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "server_cq_id",
				Description: "Azure sql server cloudquery id",
				Type:        schema.TypeUUID,
				Resolver:    schema.ParentIdResolver,
			},
			{
				Name:        "transparent_data_encryption",
				Description: "TransparentDataEncryption represents a database transparent data encryption configuration",
				Type:        schema.TypeJSON,
				Resolver:    ResolveSqlDatabaseTransparentDataEncryption,
			},
			{
				Name:        "sku_name",
				Description: "The name of the SKU, typically, a letter + Number code, eg",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Sku.Name"),
			},
			{
				Name:        "sku_tier",
				Description: "The tier or edition of the particular SKU, eg",
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
				Description: "Kind of database",
				Type:        schema.TypeString,
			},
			{
				Name:        "managed_by",
				Description: "Resource that manages the database",
				Type:        schema.TypeString,
			},
			{
				Name:        "create_mode",
				Description: "Specifies the mode of database creation  Default: regular database creation  Copy: creates a database as a copy of an existing database",
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
				Description: "The name of the sample schema to apply when creating this database",
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
				Description: "The status of the database",
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
				Name:     "creation_date_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("DatabaseProperties.CreationDate.Time"),
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
				Name:     "restore_point_in_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("DatabaseProperties.RestorePointInTime.Time"),
			},
			{
				Name:     "source_database_deletion_date_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("DatabaseProperties.SourceDatabaseDeletionDate.Time"),
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
				Description: "Collation of the metadata catalog",
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
				Description: "The license type to apply for this database",
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
				Name:     "earliest_restore_date_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("DatabaseProperties.EarliestRestoreDate.Time"),
			},
			{
				Name:        "read_scale",
				Description: "The state of read-only routing",
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
				Description: "The secondary type of the database if it is a secondary",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DatabaseProperties.SecondaryType"),
			},
			{
				Name:        "current_sku_name",
				Description: "The name of the SKU, typically, a letter + Number code, eg",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DatabaseProperties.CurrentSku.Name"),
			},
			{
				Name:        "current_sku_tier",
				Description: "The tier or edition of the particular SKU, eg",
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
				Description: "Time in minutes after which database is automatically paused",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("DatabaseProperties.AutoPauseDelay"),
			},
			{
				Name:        "storage_account_type",
				Description: "The storage account type used to store backups for this database",
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
				Description: "Maintenance configuration id assigned to the database",
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
				Description: "DatabaseBlobAuditingPolicy a database blob auditing policy",
				Resolver:    fetchSqlDatabaseDbBlobAuditingPolicies,
				Columns: []schema.Column{
					{
						Name:        "database_cq_id",
						Description: "Unique CloudQuery ID of azure_sql_databases table (FK)",
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
						Description: "Specifies the state of the policy",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DatabaseBlobAuditingPolicyProperties.State"),
					},
					{
						Name:        "storage_endpoint",
						Description: "Specifies the blob storage endpoint (eg",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DatabaseBlobAuditingPolicyProperties.StorageEndpoint"),
					},
					{
						Name:        "storage_account_access_key",
						Description: "Specifies the identifier key of the auditing storage account If state is Enabled and storageEndpoint is specified, not specifying the storageAccountAccessKey will use SQL server system-assigned managed identity to access the storage Prerequisites for using managed identity authentication: 1",
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
						Description: "this will audit all the queries and stored procedures executed against the database, as well as successful and failed logins:  BATCH_COMPLETED_GROUP, SUCCESSFUL_DATABASE_AUTHENTICATION_GROUP, FAILED_DATABASE_AUTHENTICATION_GROUP  This above combination is also the set that is configured by default when enabling auditing from the Azure portal  The supported action groups to audit are (note: choose only specific groups that cover your auditing needs",
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
						Description: "Specifies whether audit events are sent to Azure Monitor",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("DatabaseBlobAuditingPolicyProperties.IsAzureMonitorTargetEnabled"),
					},
					{
						Name:        "queue_delay_ms",
						Description: "Specifies the amount of time in milliseconds that can elapse before audit actions are forced to be processed The default minimum value is 1000 (1 second)",
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
			{
				Name:        "azure_sql_database_db_vulnerability_assessments",
				Description: "DatabaseVulnerabilityAssessment a database vulnerability assessment",
				Resolver:    fetchSqlDatabaseDbVulnerabilityAssessments,
				Columns: []schema.Column{
					{
						Name:        "database_cq_id",
						Description: "Unique CloudQuery ID of azure_sql_databases table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "storage_container_path",
						Description: "A blob storage container path to hold the scan results (eg",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DatabaseVulnerabilityAssessmentProperties.StorageContainerPath"),
					},
					{
						Name:        "storage_container_sas_key",
						Description: "A shared access signature (SAS Key) that has read and write access to the blob container specified in 'storageContainerPath' parameter",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DatabaseVulnerabilityAssessmentProperties.StorageContainerSasKey"),
					},
					{
						Name:        "storage_account_access_key",
						Description: "Specifies the identifier key of the storage account for vulnerability assessment scan results",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DatabaseVulnerabilityAssessmentProperties.StorageAccountAccessKey"),
					},
					{
						Name:        "recurring_scans_is_enabled",
						Description: "Recurring scans state",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("DatabaseVulnerabilityAssessmentProperties.RecurringScans.IsEnabled"),
					},
					{
						Name:        "recurring_scans_email_subscription_admins",
						Description: "Specifies that the schedule scan notification will be is sent to the subscription administrators",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("DatabaseVulnerabilityAssessmentProperties.RecurringScans.EmailSubscriptionAdmins"),
					},
					{
						Name:        "recurring_scans_emails",
						Description: "Specifies an array of e-mail addresses to which the scan notification is sent",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("DatabaseVulnerabilityAssessmentProperties.RecurringScans.Emails"),
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
			{
				Name:        "azure_sql_database_db_vulnerability_assessment_scans",
				Description: "VulnerabilityAssessmentScanRecord a vulnerability assessment scan record",
				Resolver:    fetchSqlDatabaseDbVulnerabilityAssessmentScans,
				Columns: []schema.Column{
					{
						Name:        "database_cq_id",
						Description: "Unique CloudQuery ID of azure_sql_databases table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "scan_id",
						Description: "The scan ID",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VulnerabilityAssessmentScanRecordProperties.ScanID"),
					},
					{
						Name:        "trigger_type",
						Description: "The scan trigger type",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VulnerabilityAssessmentScanRecordProperties.TriggerType"),
					},
					{
						Name:        "state",
						Description: "The scan status",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VulnerabilityAssessmentScanRecordProperties.State"),
					},
					{
						Name:     "start_time",
						Type:     schema.TypeTimestamp,
						Resolver: schema.PathResolver("VulnerabilityAssessmentScanRecordProperties.StartTime.Time"),
					},
					{
						Name:     "end_time",
						Type:     schema.TypeTimestamp,
						Resolver: schema.PathResolver("VulnerabilityAssessmentScanRecordProperties.EndTime.Time"),
					},
					{
						Name:        "errors",
						Description: "The scan errors",
						Type:        schema.TypeJSON,
						Resolver:    resolveSqlDatabaseDbVulnerabilityAssessmentScansErrors,
					},
					{
						Name:        "storage_container_path",
						Description: "The scan results storage container path",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VulnerabilityAssessmentScanRecordProperties.StorageContainerPath"),
					},
					{
						Name:        "number_of_failed_security_checks",
						Description: "The number of failed security checks",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("VulnerabilityAssessmentScanRecordProperties.NumberOfFailedSecurityChecks"),
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
			{
				Name:        "azure_sql_database_db_threat_detection_policies",
				Description: "DatabaseSecurityAlertPolicy contains information about a database Threat Detection policy",
				Resolver:    fetchSqlDatabaseDbThreatDetectionPolicies,
				Columns: []schema.Column{
					{
						Name:        "database_cq_id",
						Description: "Unique CloudQuery ID of azure_sql_databases table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "location",
						Description: "The geo-location where the resource lives",
						Type:        schema.TypeString,
					},
					{
						Name:        "kind",
						Description: "Resource kind",
						Type:        schema.TypeString,
					},
					{
						Name:        "state",
						Description: "Specifies the state of the policy",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DatabaseSecurityAlertPolicyProperties.State"),
					},
					{
						Name:        "disabled_alerts",
						Description: "Specifies the semicolon-separated list of alerts that are disabled, or empty string to disable no alerts",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DatabaseSecurityAlertPolicyProperties.DisabledAlerts"),
					},
					{
						Name:        "email_addresses",
						Description: "Specifies the semicolon-separated list of e-mail addresses to which the alert is sent",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DatabaseSecurityAlertPolicyProperties.EmailAddresses"),
					},
					{
						Name:        "email_account_admins",
						Description: "Specifies that the alert is sent to the account administrators",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DatabaseSecurityAlertPolicyProperties.EmailAccountAdmins"),
					},
					{
						Name:        "storage_endpoint",
						Description: "Specifies the blob storage endpoint (eg",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DatabaseSecurityAlertPolicyProperties.StorageEndpoint"),
					},
					{
						Name:        "storage_account_access_key",
						Description: "Specifies the identifier key of the Threat Detection audit storage account",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DatabaseSecurityAlertPolicyProperties.StorageAccountAccessKey"),
					},
					{
						Name:        "retention_days",
						Description: "Specifies the number of days to keep in the Threat Detection audit logs",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("DatabaseSecurityAlertPolicyProperties.RetentionDays"),
					},
					{
						Name:        "use_server_default",
						Description: "Specifies whether to use the default server policy",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DatabaseSecurityAlertPolicyProperties.UseServerDefault"),
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

func fetchSqlDatabases(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
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
func ResolveSqlDatabaseTransparentDataEncryption(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	svc := meta.(*client.Client).Services().SQL.TransparentDataEncryptions
	database, ok := resource.Item.(sql.Database)
	if !ok {
		return fmt.Errorf("expected sql.Database but got %T", resource.Item)
	}
	details, err := client.ParseResourceID(*database.ID)
	if err != nil {
		return err
	}
	server, ok := resource.Parent.Item.(sql.Server)
	if !ok {
		return fmt.Errorf("not a sql.Server instance: %T", resource.Parent.Item)
	}
	result, err := svc.Get(ctx, details.ResourceGroup, *server.Name, *database.Name)
	if err != nil {
		return err
	}

	data, err := json.Marshal(result)
	if err != nil {
		return err
	}

	return resource.Set(c.Name, data)
}
func fetchSqlDatabaseDbBlobAuditingPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().SQL.DatabaseBlobAuditingPolicies
	database := parent.Item.(sql.Database)
	details, err := client.ParseResourceID(*database.ID)
	if err != nil {
		return err
	}
	server, ok := parent.Parent.Item.(sql.Server)
	if !ok {
		return fmt.Errorf("not a sql.Server instance: %T", parent.Parent.Item)
	}
	result, err := svc.ListByDatabase(ctx, details.ResourceGroup, *server.Name, *database.Name)
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
func fetchSqlDatabaseDbVulnerabilityAssessments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().SQL.DatabaseVulnerabilityAssessments
	database := parent.Item.(sql.Database)
	details, err := client.ParseResourceID(*database.ID)
	if err != nil {
		return err
	}
	server, ok := parent.Parent.Item.(sql.Server)
	if !ok {
		return fmt.Errorf("not a sql.Server instance: %T", parent.Parent.Item)
	}
	result, err := svc.ListByDatabase(ctx, details.ResourceGroup, *server.Name, *database.Name)
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
func fetchSqlDatabaseDbVulnerabilityAssessmentScans(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().SQL.DatabaseVulnerabilityAssessmentScans
	database := parent.Item.(sql.Database)
	details, err := client.ParseResourceID(*database.ID)
	if err != nil {
		return err
	}
	server, ok := parent.Parent.Item.(sql.Server)
	if !ok {
		return fmt.Errorf("not a sql.Server instance: %T", parent.Parent.Item)
	}
	result, err := svc.ListByDatabase(ctx, details.ResourceGroup, *server.Name, *database.Name)
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
func resolveSqlDatabaseDbVulnerabilityAssessmentScansErrors(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(sql.VulnerabilityAssessmentScanRecord)
	if !ok {
		return fmt.Errorf("expected sql.VulnerabilityAssessmentScanRecord but got %T", resource.Item)
	}

	if p.Errors == nil {
		return nil
	}

	parsed := make([]map[string]interface{}, 0, len(*p.Errors))

	for _, e := range *p.Errors {
		parsed = append(parsed, map[string]interface{}{
			"code":    *e.Code,
			"message": *e.Message,
		})
	}

	data, err := json.Marshal(parsed)
	if err != nil {
		return err
	}

	return resource.Set(c.Name, data)
}
func fetchSqlDatabaseDbThreatDetectionPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().SQL.DatabaseThreatDetectionPolicies
	database := parent.Item.(sql.Database)
	details, err := client.ParseResourceID(*database.ID)
	if err != nil {
		return err
	}
	server, ok := parent.Parent.Item.(sql.Server)
	if !ok {
		return fmt.Errorf("not a sql.Server instance: %T", parent.Parent.Item)
	}
	result, err := svc.Get(ctx, details.ResourceGroup, *server.Name, *database.Name)
	if err != nil {
		return err
	}
	res <- result
	return nil
}
