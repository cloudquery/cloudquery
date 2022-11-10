// Auto generated code - DO NOT EDIT.

package sql

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql"
)

func serverSecurityAlertPolicies() *schema.Table {
	return &schema.Table{
		Name:        "azure_sql_server_security_alert_policies",
		Description: `https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql#ServerSecurityAlertPolicy`,
		Resolver:    fetchSQLServerSecurityAlertPolicies,
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
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "disabled_alerts",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("DisabledAlerts"),
			},
			{
				Name:     "email_addresses",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("EmailAddresses"),
			},
			{
				Name:     "email_account_admins",
				Type:     schema.TypeBool,
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
				Name:     "creation_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreationTime"),
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

func fetchSQLServerSecurityAlertPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().SQL.ServerSecurityAlertPolicies

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
