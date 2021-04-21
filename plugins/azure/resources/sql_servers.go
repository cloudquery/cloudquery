package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func SQLServers() *schema.Table {
	return &schema.Table{
		Name:         "azure_sql_servers",
		Resolver:     fetchSqlServers,
		Multiplex:    client.SubscriptionMultiplex,
		DeleteFilter: client.DeleteSubscriptionFilter,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name: "kind",
				Type: schema.TypeString,
			},
			{
				Name:     "fully_qualified_domain_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServerProperties.FullyQualifiedDomainName"),
			},
			{
				Name:     "version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServerProperties.Version"),
			},
			{
				Name:     "administrator_login",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServerProperties.AdministratorLogin"),
			},
			{
				Name:     "administrator_login_password",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServerProperties.AdministratorLoginPassword"),
			},
			{
				Name:     "external_administrator_sid",
				Type:     schema.TypeUUID,
				Resolver: schema.PathResolver("ServerProperties.ExternalAdministratorSid"),
			},
			{
				Name:     "external_administrator_login",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServerProperties.ExternalAdministratorLogin"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServerProperties.State"),
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
			sqlDatabases(),
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
