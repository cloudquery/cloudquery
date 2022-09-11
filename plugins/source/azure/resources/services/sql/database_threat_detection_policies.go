// Auto generated code - DO NOT EDIT.

package sql

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/sql/mgmt/sql"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/pkg/errors"
)

func databaseThreatDetectionPolicies() *schema.Table {
	return &schema.Table{
		Name:     "azure_sql_database_threat_detection_policies",
		Resolver: fetchSQLDatabaseThreatDetectionPolicies,
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
				Name:     "location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Location"),
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
				Name:     "disabled_alerts",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DisabledAlerts"),
			},
			{
				Name:     "email_addresses",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EmailAddresses"),
			},
			{
				Name:     "email_account_admins",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EmailAccountAdmins"),
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
				Name:     "use_server_default",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("UseServerDefault"),
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

func fetchSQLDatabaseThreatDetectionPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().SQL.DatabaseThreatDetectionPolicies

	server := parent.Parent.Item.(sql.Server)
	database := parent.Item.(sql.Database)
	resourceDetails, err := client.ParseResourceID(*database.ID)
	if err != nil {
		return errors.WithStack(err)
	}
	response, err := svc.Get(ctx, resourceDetails.ResourceGroup, *server.Name, *database.Name)
	if err != nil {
		return errors.WithStack(err)
	}
	res <- response
	return nil
}
