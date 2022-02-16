package sql

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func sqlManagedDatabases() *schema.Table {
	return &schema.Table{
		Name:        "azure_sql_managed_databases",
		Description: "ManagedDatabase a managed database resource.",
		Resolver:    fetchSqlManagedDatabases,
		Columns: []schema.Column{
			{
				Name:        "collation",
				Description: "Collation of the managed database.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ManagedDatabaseProperties.Collation"),
			},
			{
				Name:        "status",
				Description: "Status of the database",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ManagedDatabaseProperties.Status"),
			},
			{
				Name:     "creation_date_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("ManagedDatabaseProperties.CreationDate.Time"),
			},
			{
				Name:     "earliest_restore_point_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("ManagedDatabaseProperties.EarliestRestorePoint.Time"),
			},
			{
				Name:     "restore_point_in_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("ManagedDatabaseProperties.RestorePointInTime.Time"),
			},
			{
				Name:        "default_secondary_location",
				Description: "Geo paired region.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ManagedDatabaseProperties.DefaultSecondaryLocation"),
			},
			{
				Name:        "catalog_collation",
				Description: "Collation of the metadata catalog",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ManagedDatabaseProperties.CatalogCollation"),
			},
			{
				Name:        "create_mode",
				Description: "Managed database create mode",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ManagedDatabaseProperties.CreateMode"),
			},
			{
				Name:        "storage_container_uri",
				Description: "Conditional",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ManagedDatabaseProperties.StorageContainerURI"),
			},
			{
				Name:        "source_database_id",
				Description: "The resource identifier of the source database associated with create operation of this database.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ManagedDatabaseProperties.SourceDatabaseID"),
			},
			{
				Name:        "restorable_dropped_database_id",
				Description: "The restorable dropped database resource id to restore when creating this database.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ManagedDatabaseProperties.RestorableDroppedDatabaseID"),
			},
			{
				Name:        "storage_container_sas_token",
				Description: "SAS token used to access resources",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ManagedDatabaseProperties.StorageContainerSasToken"),
			},
			{
				Name:        "failover_group_id",
				Description: "Instance Failover Group resource identifier that this managed database belongs to.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ManagedDatabaseProperties.FailoverGroupID"),
			},
			{
				Name:        "recoverable_database_id",
				Description: "The resource identifier of the recoverable database associated with create operation of this database.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ManagedDatabaseProperties.RecoverableDatabaseID"),
			},
			{
				Name:        "long_term_retention_backup_resource_id",
				Description: "The name of the Long Term Retention backup to be used for restore of this managed database.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ManagedDatabaseProperties.LongTermRetentionBackupResourceID"),
			},
			{
				Name:        "auto_complete_restore",
				Description: "Whether to auto complete restore of this managed database.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("ManagedDatabaseProperties.AutoCompleteRestore"),
			},
			{
				Name:        "last_backup_name",
				Description: "Last backup file name for restore of this managed database.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ManagedDatabaseProperties.LastBackupName"),
			},
			{
				Name:        "location",
				Description: "Resource location.",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "Resource tags.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "id",
				Description: "Resource ID.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "Resource name.",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "Resource type.",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "azure_sql_managed_database_vulnerability_assessments",
				Description: "DatabaseVulnerabilityAssessment a database vulnerability assessment.",
				Resolver:    fetchSqlManagedDatabaseVulnerabilityAssessments,
				Columns: []schema.Column{
					{
						Name:        "managed_database_cq_id",
						Description: "Unique CloudQuery ID of azure_sql_managed_databases table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "storage_container_path",
						Description: "A blob storage container path to hold the scan results",
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
						Description: "Recurring scans state.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("DatabaseVulnerabilityAssessmentProperties.RecurringScans.IsEnabled"),
					},
					{
						Name:        "recurring_scans_email_subscription_admins",
						Description: "Specifies that the schedule scan notification will be is sent to the subscription administrators.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("DatabaseVulnerabilityAssessmentProperties.RecurringScans.EmailSubscriptionAdmins"),
					},
					{
						Name:        "recurring_scans_emails",
						Description: "Specifies an array of e-mail addresses to which the scan notification is sent.",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("DatabaseVulnerabilityAssessmentProperties.RecurringScans.Emails"),
					},
					{
						Name:        "id",
						Description: "Resource ID.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
					{
						Name:        "name",
						Description: "Resource name.",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "Resource type.",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "azure_sql_managed_database_vulnerability_assessment_scans",
				Description: "VulnerabilityAssessmentScanRecord a vulnerability assessment scan record.",
				Resolver:    fetchSqlManagedDatabaseVulnerabilityAssessmentScans,
				Columns: []schema.Column{
					{
						Name:        "managed_database_cq_id",
						Description: "Unique CloudQuery ID of azure_sql_managed_databases table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "scan_id",
						Description: "The scan ID.",
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
						Description: "The scan errors.",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "storage_container_path",
						Description: "The scan results storage container path.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VulnerabilityAssessmentScanRecordProperties.StorageContainerPath"),
					},
					{
						Name:        "number_of_failed_security_checks",
						Description: "The number of failed security checks.",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("VulnerabilityAssessmentScanRecordProperties.NumberOfFailedSecurityChecks"),
					},
					{
						Name:        "id",
						Description: "Resource ID.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
					{
						Name:        "name",
						Description: "Resource name.",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "Resource type.",
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

func fetchSqlManagedDatabases(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().SQL.ManagedDatabases
	server, ok := parent.Item.(sql.ManagedInstance)
	if !ok {
		return fmt.Errorf("not an sql.ManagedInstance instance: %T", parent.Item)
	}
	resourceDetails, err := client.ParseResourceID(*server.ID)
	if err != nil {
		return err
	}
	databases, err := svc.ListByInstance(ctx, resourceDetails.ResourceGroup, *server.Name)
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
func fetchSqlManagedDatabaseVulnerabilityAssessments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().SQL.ManagedDatabaseVulnerabilityAssessments
	database, ok := parent.Item.(sql.ManagedDatabase)
	if !ok {
		return fmt.Errorf("not an sql.ManagedDatabase instance: %T", parent.Item)
	}
	details, err := client.ParseResourceID(*database.ID)
	if err != nil {
		return err
	}
	server, ok := parent.Parent.Item.(sql.ManagedInstance)
	if !ok {
		return fmt.Errorf("not a sql.ManagedInstance instance: %T", parent.Parent.Item)
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
func fetchSqlManagedDatabaseVulnerabilityAssessmentScans(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().SQL.ManagedDatabaseVulnerabilityAssessmentScans
	database, ok := parent.Item.(sql.ManagedDatabase)
	if !ok {
		return fmt.Errorf("not an sql.ManagedDatabase instance: %T", parent.Item)
	}
	details, err := client.ParseResourceID(*database.ID)
	if err != nil {
		return err
	}
	server, ok := parent.Parent.Item.(sql.ManagedInstance)
	if !ok {
		return fmt.Errorf("not a sql.ManagedInstance instance: %T", parent.Parent.Item)
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
