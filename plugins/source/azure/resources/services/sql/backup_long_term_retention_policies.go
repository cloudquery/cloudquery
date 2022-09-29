// Auto generated code - DO NOT EDIT.

package sql

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql"
)

func backupLongTermRetentionPolicies() *schema.Table {
	return &schema.Table{
		Name:     "azure_sql_backup_long_term_retention_policies",
		Resolver: fetchSQLBackupLongTermRetentionPolicies,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "sql_database_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
			},
			{
				Name:     "weekly_retention",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("WeeklyRetention"),
			},
			{
				Name:     "monthly_retention",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MonthlyRetention"),
			},
			{
				Name:     "yearly_retention",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("YearlyRetention"),
			},
			{
				Name:     "week_of_year",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("WeekOfYear"),
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

func fetchSQLBackupLongTermRetentionPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().SQL.BackupLongTermRetentionPolicies

	server := parent.Parent.Item.(sql.Server)
	database := parent.Item.(sql.Database)
	resourceDetails, err := client.ParseResourceID(*database.ID)
	if err != nil {
		return err
	}
	response, err := svc.ListByDatabase(ctx, resourceDetails.ResourceGroup, *server.Name, *database.Name)
	if err != nil {
		return err
	}
	res <- response
	return nil
}
