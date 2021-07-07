package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func SQLServers() *schema.Table {
	return &schema.Table{
		Name:         "azure_sql_servers",
		Description:  "Azure sql server",
		Resolver:     fetchSqlServers,
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
				Name:        "kind",
				Description: "Kind of sql server  This is metadata used for the Azure portal experience",
				Type:        schema.TypeString,
			},
			{
				Name:        "fully_qualified_domain_name",
				Description: "The fully qualified domain name of the server",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServerProperties.FullyQualifiedDomainName"),
			},
			{
				Name:        "version",
				Description: "The version of the server Possible values include: 'TwoFullStopZero', 'OneTwoFullStopZero'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServerProperties.Version"),
			},
			{
				Name:        "administrator_login",
				Description: "Administrator username for the server Can only be specified when the server is being created (and is required for creation)",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServerProperties.AdministratorLogin"),
			},
			{
				Name:        "administrator_login_password",
				Description: "The administrator login password (required for server creation)",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServerProperties.AdministratorLoginPassword"),
			},
			{
				Name:        "external_administrator_sid",
				Description: "The ID of the Active Azure Directory object with admin permissions on this server Legacy parameter, always null To check for Active Directory admin, query /servers/{serverName}/administrators",
				Type:        schema.TypeUUID,
				Resolver:    schema.PathResolver("ServerProperties.ExternalAdministratorSid"),
			},
			{
				Name:        "external_administrator_login",
				Description: "The display name of the Azure Active Directory object with admin permissions on this server Legacy parameter, always null To check for Active Directory admin, query /servers/{serverName}/administrators",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServerProperties.ExternalAdministratorLogin"),
			},
			{
				Name:        "state",
				Description: "The state of the server Possible values include: 'ServerStateReady', 'ServerStateDisabled'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServerProperties.State"),
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
			SQLDatabases(),
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchSqlServers(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	svc := meta.(*client.Client).Services().SQL.Servers
	servers, err := svc.List(ctx)
	if err != nil {
		return err
	}
	res <- *servers.Value
	return nil
}
