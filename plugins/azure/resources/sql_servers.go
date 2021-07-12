package resources

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql"
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
				Name:        "identity_principal_id",
				Description: "The Azure Active Directory principal id",
				Type:        schema.TypeUUID,
				Resolver:    schema.PathResolver("Identity.PrincipalID"),
			},
			{
				Name:        "identity_type",
				Description: "The identity type Set this to 'SystemAssigned' in order to automatically create and assign an Azure Active Directory principal for the resource Possible values include: 'None', 'SystemAssigned', 'UserAssigned'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Identity.Type"),
			},
			{
				Name:        "identity_tenant_id",
				Description: "The Azure Active Directory tenant id",
				Type:        schema.TypeUUID,
				Resolver:    schema.PathResolver("Identity.TenantID"),
			},
			{
				Name:        "kind",
				Description: "Kind of sql server This is metadata used for the Azure portal experience",
				Type:        schema.TypeString,
			},
			{
				Name:        "administrator_login",
				Description: "Administrator username for the server Once created it cannot be changed",
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
				Name:        "version",
				Description: "The version of the server",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServerProperties.Version"),
			},
			{
				Name:        "state",
				Description: "The state of the server",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServerProperties.State"),
			},
			{
				Name:        "fully_qualified_domain_name",
				Description: "The fully qualified domain name of the server",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServerProperties.FullyQualifiedDomainName"),
			},
			{
				Name:        "minimal_tls_version",
				Description: "Minimal TLS version Allowed values: '10', '11', '12'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServerProperties.MinimalTLSVersion"),
			},
			{
				Name:        "public_network_access",
				Description: "Whether or not public endpoint access is allowed for this server  Value is optional but if passed in, must be 'Enabled' or 'Disabled' Possible values include: 'ServerPublicNetworkAccessEnabled', 'ServerPublicNetworkAccessDisabled'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServerProperties.PublicNetworkAccess"),
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
			{
				Name:        "azure_sql_server_private_endpoint_connections",
				Description: "List of private endpoint connections on a server",
				Resolver:    fetchSqlServerPrivateEndpointConnections,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"server_cq_id", "id"}},
				Columns: []schema.Column{
					{
						Name:        "server_cq_id",
						Description: "Unique ID of azure_sql_servers table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "id",
						Description: "Resource ID",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
					{
						Name:        "private_endpoint_id",
						Description: "Resource id of the private endpoint",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Properties.PrivateEndpoint.ID"),
					},
					{
						Name:        "private_link_service_connection_state_status",
						Description: "The private link service connection status Possible values include: 'Approved', 'Pending', 'Rejected', 'Disconnected'",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Properties.PrivateLinkServiceConnectionState.Status"),
					},
					{
						Name:        "private_link_service_connection_state_description",
						Description: "The private link service connection description",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Properties.PrivateLinkServiceConnectionState.Description"),
					},
					{
						Name:        "private_link_service_connection_state_actions_required",
						Description: "The actions required for private link service connection Possible values include: 'PrivateLinkServiceConnectionStateActionsRequireNone'",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Properties.PrivateLinkServiceConnectionState.ActionsRequired"),
					},
					{
						Name:        "provisioning_state",
						Description: "State of the private endpoint connection Possible values include: 'PrivateEndpointProvisioningStateApproving', 'PrivateEndpointProvisioningStateReady', 'PrivateEndpointProvisioningStateDropping', 'PrivateEndpointProvisioningStateFailed', 'PrivateEndpointProvisioningStateRejecting'",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Properties.ProvisioningState"),
					},
				},
			},
			{
				Name:        "azure_sql_server_firewall_rules",
				Description: "The list of server firewall rules.",
				Resolver:    fetchSqlServerFirewallRules,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"server_cq_id", "id"}},
				Columns: []schema.Column{
					{
						Name:        "server_cq_id",
						Description: "Unique ID of azure_sql_servers table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "kind",
						Description: "Kind of server that contains this firewall rule",
						Type:        schema.TypeString,
					},
					{
						Name:        "location",
						Description: "Location of the server that contains this firewall rule",
						Type:        schema.TypeString,
					},
					{
						Name:        "start_ip_address",
						Description: "The start IP address of the firewall rule Must be IPv4 format Use value '0000' to represent all Azure-internal IP addresses",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("FirewallRuleProperties.StartIPAddress"),
					},
					{
						Name:        "end_ip_address",
						Description: "The end IP address of the firewall rule Must be IPv4 format Must be greater than or equal to startIpAddress Use value '0000' to represent all Azure-internal IP addresses",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("FirewallRuleProperties.EndIPAddress"),
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
				Name:        "azure_sql_server_admins",
				Description: "ServerAzureADAdministrator azure Active Directory administrator",
				Resolver:    fetchSqlServerAdmins,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"server_cq_id", "id"}},
				Columns: []schema.Column{
					{
						Name:        "server_cq_id",
						Description: "Unique ID of azure_sql_servers table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "administrator_type",
						Description: "Type of the sever administrator",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("AdministratorProperties.AdministratorType"),
					},
					{
						Name:        "login",
						Description: "Login name of the server administrator",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("AdministratorProperties.Login"),
					},
					{
						Name:        "sid",
						Description: "SID (object ID) of the server administrator",
						Type:        schema.TypeUUID,
						Resolver:    schema.PathResolver("AdministratorProperties.Sid"),
					},
					{
						Name:        "tenant_id",
						Description: "Tenant ID of the administrator",
						Type:        schema.TypeUUID,
						Resolver:    schema.PathResolver("AdministratorProperties.TenantID"),
					},
					{
						Name:        "azure_ad_only_authentication",
						Description: "Azure Active Directory only Authentication enabled",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("AdministratorProperties.AzureADOnlyAuthentication"),
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
func fetchSqlServers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	svc := meta.(*client.Client).Services().SQL.Servers
	servers, err := svc.List(ctx)
	if err != nil {
		return err
	}
	for servers.NotDone() {
		res <- servers.Values()
		if err := servers.NextWithContext(ctx); err != nil {
			return err
		}
	}
	return nil
}

func fetchSqlServerPrivateEndpointConnections(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	server, ok := parent.Item.(sql.Server)
	if !ok {
		return fmt.Errorf("not an sql.Server instance: %#v", parent.Item)
	}
	if server.PrivateEndpointConnections != nil {
		res <- *server.PrivateEndpointConnections
	}
	return nil
}

func fetchSqlServerFirewallRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	svc := meta.(*client.Client).Services().SQL.Firewall
	server, ok := parent.Item.(sql.Server)
	if !ok {
		return fmt.Errorf("not an sql.Server instance: %#v", parent.Item)
	}
	details, err := client.ParseResourceID(*server.ID)
	if err != nil {
		return err
	}
	result, err := svc.ListByServer(ctx, details.ResourceGroup, *server.Name)
	if err != nil {
		return err
	}
	if result.Value != nil {
		res <- *result.Value
	}
	return nil
}

func fetchSqlServerAdmins(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	svc := meta.(*client.Client).Services().SQL.ServerAdmins
	server, ok := parent.Item.(sql.Server)
	if !ok {
		return fmt.Errorf("not an sql.Server instance: %#v", parent.Item)
	}
	details, err := client.ParseResourceID(*server.ID)
	if err != nil {
		return err
	}
	result, err := svc.ListByServer(ctx, details.ResourceGroup, *server.Name)
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
