// Auto generated code - DO NOT EDIT.

package mariadb

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/pkg/errors"
)

func Servers() *schema.Table {
	return &schema.Table{
		Name:      "azure_mariadb_servers",
		Resolver:  fetchMariaDBServers,
		Multiplex: client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "sku",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Sku"),
			},
			{
				Name:     "administrator_login",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AdministratorLogin"),
			},
			{
				Name:     "version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Version"),
			},
			{
				Name:     "ssl_enforcement",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SslEnforcement"),
			},
			{
				Name:     "user_visible_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("UserVisibleState"),
			},
			{
				Name:     "fully_qualified_domain_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FullyQualifiedDomainName"),
			},
			{
				Name:     "earliest_restore_date",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EarliestRestoreDate"),
			},
			{
				Name:     "storage_profile",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("StorageProfile"),
			},
			{
				Name:     "replication_role",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReplicationRole"),
			},
			{
				Name:     "master_server_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MasterServerID"),
			},
			{
				Name:     "replica_capacity",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ReplicaCapacity"),
			},
			{
				Name:     "public_network_access",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PublicNetworkAccess"),
			},
			{
				Name:     "private_endpoint_connections",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PrivateEndpointConnections"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
			{
				Name:     "location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Location"),
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

		Relations: []*schema.Table{
			configurations(),
		},
	}
}

func fetchMariaDBServers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().MariaDB.Servers

	response, err := svc.List(ctx)
	if err != nil {
		return errors.WithStack(err)
	}
	if response.Value == nil {
		return nil
	}
	res <- *response.Value

	return nil
}
