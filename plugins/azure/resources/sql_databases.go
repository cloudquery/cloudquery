package resources

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/sql/mgmt/2014-04-01/sql"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func sqlDatabases() *schema.Table {
	return &schema.Table{
		Name:         "azure_sql_databases",
		Resolver:     fetchSqlDatabases,
		Multiplex:    client.SubscriptionMultiplex,
		DeleteFilter: client.DeleteSubscriptionFilter,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "server_id",
				Type:     schema.TypeUUID,
				Resolver: schema.ParentIdResolver,
			},
			{
				Name: "kind",
				Type: schema.TypeString,
			},
			{
				Name:     "collation",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DatabaseProperties.Collation"),
			},
			{
				Name:     "creation_date_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("DatabaseProperties.CreationDate.Time"),
			},
			{
				Name:     "containment_state",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("DatabaseProperties.ContainmentState"),
			},
			{
				Name:     "current_service_objective_id",
				Type:     schema.TypeUUID,
				Resolver: schema.PathResolver("DatabaseProperties.CurrentServiceObjectiveID"),
			},
			{
				Name:     "database_id",
				Type:     schema.TypeUUID,
				Resolver: schema.PathResolver("DatabaseProperties.DatabaseID"),
			},
			{
				Name:     "earliest_restore_date_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("DatabaseProperties.EarliestRestoreDate.Time"),
			},
			{
				Name:     "create_mode",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DatabaseProperties.CreateMode"),
			},
			{
				Name:     "source_database_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DatabaseProperties.SourceDatabaseID"),
			},
			{
				Name:     "source_database_deletion_date_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("DatabaseProperties.SourceDatabaseDeletionDate.Time"),
			},
			{
				Name:     "restore_point_in_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("DatabaseProperties.RestorePointInTime.Time"),
			},
			{
				Name:     "recovery_services_recovery_point_resource_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DatabaseProperties.RecoveryServicesRecoveryPointResourceID"),
			},
			{
				Name:     "edition",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DatabaseProperties.Edition"),
			},
			{
				Name:     "max_size_bytes",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DatabaseProperties.MaxSizeBytes"),
			},
			{
				Name:     "requested_service_objective_id",
				Type:     schema.TypeUUID,
				Resolver: schema.PathResolver("DatabaseProperties.RequestedServiceObjectiveID"),
			},
			{
				Name:     "requested_service_objective_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DatabaseProperties.RequestedServiceObjectiveName"),
			},
			{
				Name:     "service_level_objective",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DatabaseProperties.ServiceLevelObjective"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DatabaseProperties.Status"),
			},
			{
				Name:     "elastic_pool_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DatabaseProperties.ElasticPoolName"),
			},
			{
				Name:     "default_secondary_location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DatabaseProperties.DefaultSecondaryLocation"),
			},
			{
				Name:     "failover_group_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DatabaseProperties.FailoverGroupID"),
			},
			{
				Name:     "read_scale",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DatabaseProperties.ReadScale"),
			},
			{
				Name:     "sample_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DatabaseProperties.SampleName"),
			},
			{
				Name:     "zone_redundant",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("DatabaseProperties.ZoneRedundant"),
			},
			{
				Name: "location",
				Type: schema.TypeString,
			},
			{
				Name: "tags",
				Type: schema.TypeJSON,
			},
			{
				Name:     "resource_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name: "type",
				Type: schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "azure_sql_database_transparent_data_encryptions",
				Resolver: fetchSqlDatabaseTransparentDataEncryptions,
				Columns: []schema.Column{
					{
						Name:     "database_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "location",
						Type: schema.TypeString,
					},
					{
						Name:     "status",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("TransparentDataEncryptionProperties.Status"),
					},
					{
						Name:     "resource_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ID"),
					},
					{
						Name: "name",
						Type: schema.TypeString,
					},
					{
						Name: "type",
						Type: schema.TypeString,
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
	svc := meta.(*client.Client).Services().SQL.Database
	server := parent.Item.(sql.Server)
	resourceDetails, err := client.ParseResourceID(*server.ID)
	if err != nil {
		return err
	}
	databases, err := svc.ListByServer(ctx, resourceDetails.ResourceGroup, *server.Name, "true", "")
	if err != nil {
		return err
	}
	if databases.Value == nil {
		return nil
	}
	res <- *databases.Value
	return nil
}
func fetchSqlDatabaseTransparentDataEncryptions(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	database := parent.Item.(sql.Database)
	if database.TransparentDataEncryption == nil {
		return nil
	}
	res <- *database.TransparentDataEncryption
	return nil
}
