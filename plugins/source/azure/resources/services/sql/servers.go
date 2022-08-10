package sql

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
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
				Name:          "identity_principal_id",
				Description:   "The Azure Active Directory principal id",
				Type:          schema.TypeUUID,
				Resolver:      schema.PathResolver("Identity.PrincipalID"),
				IgnoreInTests: true,
			},
			{
				Name:        "identity_type",
				Description: "The identity type.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Identity.Type"),
			},
			{
				Name:          "identity_tenant_id",
				Description:   "The Azure Active Directory tenant id",
				Type:          schema.TypeUUID,
				Resolver:      schema.PathResolver("Identity.TenantID"),
				IgnoreInTests: true,
			},
			{
				Name:        "kind",
				Description: "Kind of sql server.",
				Type:        schema.TypeString,
			},
			{
				Name:        "administrator_login",
				Description: "Administrator username for the server Once created it cannot be changed",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServerProperties.AdministratorLogin"),
			},
			{
				Name:          "administrator_login_password",
				Description:   "The administrator login password (required for server creation)",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("ServerProperties.AdministratorLoginPassword"),
				IgnoreInTests: true,
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
				Name:          "minimal_tls_version",
				Description:   "Minimal TLS version Allowed values: '10', '11', '12'",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("ServerProperties.MinimalTLSVersion"),
				IgnoreInTests: true,
			},
			{
				Name:        "public_network_access",
				Description: "Whether or not public endpoint access is allowed for this server.",
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
			SqlDatabases(),
			SQLServerEncryptionProtectors(),
			{
				Name:          "azure_sql_server_private_endpoint_connections",
				Description:   "List of private endpoint connections on a server",
				Resolver:      fetchSqlServerPrivateEndpointConnections,
				IgnoreInTests: true,
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
						Description: "The private link service connection status.",
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
						Description: "The actions required for private link service connection.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Properties.PrivateLinkServiceConnectionState.ActionsRequired"),
					},
					{
						Name:        "provisioning_state",
						Description: "State of the private endpoint connection.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Properties.ProvisioningState"),
					},
				},
			},
			{
				Name:          "azure_sql_server_firewall_rules",
				Description:   "The list of server firewall rules.",
				Resolver:      fetchSqlServerFirewallRules,
				IgnoreInTests: true,
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
						Description: "The start IP address of the firewall rule.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("FirewallRuleProperties.StartIPAddress"),
					},
					{
						Name:        "end_ip_address",
						Description: "The end IP address of the firewall rule. Must be IPv4.",
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
				Name:          "azure_sql_server_admins",
				Description:   "ServerAzureADAdministrator azure Active Directory administrator",
				Resolver:      fetchSqlServerAdmins,
				IgnoreInTests: true,
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
			{
				Name:        "azure_sql_server_db_blob_auditing_policies",
				Description: "Database blob auditing policy",
				Resolver:    fetchSqlServerDbBlobAuditingPolicies,
				Columns: []schema.Column{
					{
						Name:        "server_cq_id",
						Description: "Unique ID of azure_sql_servers table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "state",
						Description: "Specifies the state of the policy.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ServerBlobAuditingPolicyProperties.State"),
					},
					{
						Name:        "storage_endpoint",
						Description: "Specifies the blob storage endpoint.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ServerBlobAuditingPolicyProperties.StorageEndpoint"),
					},
					{
						Name:          "storage_account_access_key",
						Description:   "Specifies the identifier key of the auditing storage account.",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("ServerBlobAuditingPolicyProperties.StorageAccountAccessKey"),
						IgnoreInTests: true,
					},
					{
						Name:        "retention_days",
						Description: "Specifies the number of days to keep in the audit logs in the storage account",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("ServerBlobAuditingPolicyProperties.RetentionDays"),
					},
					{
						Name:        "audit_actions_and_groups",
						Description: "Specifies the Actions-Groups and Actions to audit.",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("ServerBlobAuditingPolicyProperties.AuditActionsAndGroups"),
					},
					{
						Name:        "storage_account_subscription_id",
						Description: "Specifies the blob storage subscription Id",
						Type:        schema.TypeUUID,
						Resolver:    schema.PathResolver("ServerBlobAuditingPolicyProperties.StorageAccountSubscriptionID"),
					},
					{
						Name:        "is_storage_secondary_key_in_use",
						Description: "Specifies whether storageAccountAccessKey value is the storage's secondary key",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("ServerBlobAuditingPolicyProperties.IsStorageSecondaryKeyInUse"),
					},
					{
						Name:        "is_azure_monitor_target_enabled",
						Description: "Specifies whether audit events are sent to Azure Monitor.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("ServerBlobAuditingPolicyProperties.IsAzureMonitorTargetEnabled"),
					},
					{
						Name:          "queue_delay_ms",
						Description:   "Specifies the amount of time in milliseconds that can elapse before audit actions are forced to be processed.",
						Type:          schema.TypeInt,
						Resolver:      schema.PathResolver("ServerBlobAuditingPolicyProperties.QueueDelayMs"),
						IgnoreInTests: true,
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
				Name:        "azure_sql_server_devops_audit_settings",
				Description: "ServerDevOpsAuditingSettings a server DevOps auditing settings",
				Resolver:    fetchSqlServerDevopsAuditSettings,
				Columns: []schema.Column{
					{
						Name:        "server_cq_id",
						Description: "Unique ID of azure_sql_servers table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:          "created_by",
						Description:   "A string identifier for the identity that created the resource",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("SystemData.CreatedBy"),
						IgnoreInTests: true,
					},
					{
						Name:        "created_by_type",
						Description: "The type of identity that created the resource.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SystemData.CreatedByType"),
					},
					{
						Name:        "created_at_time",
						Description: "The timestamp of resource creation (UTC).",
						Type:        schema.TypeTimestamp,
						Resolver:    schema.PathResolver("SystemData.CreatedAt.Time"),
					},
					{
						Name:          "last_modified_by",
						Description:   "A string identifier for the identity that last modified the resource",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("SystemData.LastModifiedBy"),
						IgnoreInTests: true,
					},
					{
						Name:        "last_modified_by_type",
						Description: "The type of identity that last modified the resource",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("SystemData.LastModifiedByType"),
					},
					{
						Name:        "last_modified_at_time",
						Description: "The timestamp of last modification (UTC).",
						Type:        schema.TypeTimestamp,
						Resolver:    schema.PathResolver("SystemData.LastModifiedAt.Time"),
					},
					{
						Name:        "is_azure_monitor_target_enabled",
						Description: "Specifies whether DevOps audit events are sent to Azure Monitor",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("ServerDevOpsAuditSettingsProperties.IsAzureMonitorTargetEnabled"),
					},
					{
						Name:        "state",
						Description: "Specifies the state of the audit If state is Enabled, storageEndpoint or isAzureMonitorTargetEnabled are required Possible values include: 'BlobAuditingPolicyStateEnabled', 'BlobAuditingPolicyStateDisabled'",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ServerDevOpsAuditSettingsProperties.State"),
					},
					{
						Name:        "storage_endpoint",
						Description: "Specifies the blob storage endpoint.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ServerDevOpsAuditSettingsProperties.StorageEndpoint"),
					},
					{
						Name:          "storage_account_access_key",
						Description:   "Specifies the identifier key of the auditing storage account.",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("ServerDevOpsAuditSettingsProperties.StorageAccountAccessKey"),
						IgnoreInTests: true,
					},
					{
						Name:        "storage_account_subscription_id",
						Description: "Specifies the blob storage subscription Id",
						Type:        schema.TypeUUID,
						Resolver:    schema.PathResolver("ServerDevOpsAuditSettingsProperties.StorageAccountSubscriptionID"),
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
				Name:        "azure_sql_server_vulnerability_assessments",
				Description: "Server vulnerability assessment",
				Resolver:    fetchSqlServerVulnerabilityAssessments,
				Columns: []schema.Column{
					{
						Name:        "server_cq_id",
						Description: "Unique CloudQuery ID of azure_sql_servers table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:          "storage_container_path",
						Description:   "A blob storage container path to hold the scan results.",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("ServerVulnerabilityAssessmentProperties.StorageContainerPath"),
						IgnoreInTests: true,
					},
					{
						Name:          "storage_container_sas_key",
						Description:   "A shared access signature (SAS Key) that has read and write access to the blob container specified in 'storageContainerPath' parameter.",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("ServerVulnerabilityAssessmentProperties.StorageContainerSasKey"),
						IgnoreInTests: true,
					},
					{
						Name:          "storage_account_access_key",
						Description:   "Specifies the identifier key of the storage account for vulnerability assessment scan results.",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("ServerVulnerabilityAssessmentProperties.StorageAccountAccessKey"),
						IgnoreInTests: true,
					},
					{
						Name:        "recurring_scans_is_enabled",
						Description: "Recurring scans state",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("ServerVulnerabilityAssessmentProperties.RecurringScans.IsEnabled"),
					},
					{
						Name:        "recurring_scans_email_subscription_admins",
						Description: "Specifies that the schedule scan notification will be is sent to the subscription administrators",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("ServerVulnerabilityAssessmentProperties.RecurringScans.EmailSubscriptionAdmins"),
					},
					{
						Name:          "recurring_scans_emails",
						Description:   "Specifies an array of e-mail addresses to which the scan notification is sent",
						Type:          schema.TypeStringArray,
						Resolver:      schema.PathResolver("ServerVulnerabilityAssessmentProperties.RecurringScans.Emails"),
						IgnoreInTests: true,
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
				Name:          "azure_sql_server_virtual_network_rules",
				Description:   "List of virtual network for a server",
				Resolver:      fetchSqlServerVirtualNetworkRules,
				IgnoreInTests: true,
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
						Name:        "name",
						Description: "Resource name",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Name"),
					},
					{
						Name:        "type",
						Description: "The virtual network rule type",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Type"),
					},
					{
						Name:        "subnet_id",
						Description: "The ARM resource id of the virtual network subnet.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VirtualNetworkSubnetID"),
					},
					{
						Name:        "ignore_missing_vnet_service_endpoint",
						Description: "Create firewall rule before the virtual network has vnet service endpoint enabled.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("IgnoreMissingVnetServiceEndpoint"),
					},
					{
						Name:        "state",
						Description: "Virtual Network Rule State",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("State"),
					},
				},
			},
			{
				Name:          "azure_sql_server_security_alert_policy",
				Description:   "List the server's threat detection policies",
				Resolver:      fetchSqlServerSecurityAlertPolicies,
				IgnoreInTests: false,
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
						Name:        "name",
						Description: "Resource name",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Name"),
					},
					{
						Name:        "type",
						Description: "The virtual network rule type",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Type"),
					},
					{
						Name:        "state",
						Description: "Specifies the state of the policy, whether it is enabled or disabled or a policy has not been applied yet on the specific database. Possible values include: 'SecurityAlertPolicyStateNew', 'SecurityAlertPolicyStateEnabled', 'SecurityAlertPolicyStateDisabled'",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("State"),
					},
					{
						Name:        "disabled_alerts",
						Description: "Specifies an array of alerts that are disabled. Allowed values are: Sql_Injection, Sql_Injection_Vulnerability, Access_Anomaly, Data_Exfiltration, Unsafe_Action",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("DisabledAlerts"),
					},
					{
						Name:        "email_addresses",
						Description: "Specifies an array of e-mail addresses to which the alert is sent.",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("EmailAddresses"),
					},
					{
						Name:        "email_account_admins",
						Description: "Specifies that the alert is sent to the account administrators.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("EmailAccountAdmins"),
					},
					{
						Name:        "storage_endpoint",
						Description: "Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("StorageEndpoint"),
					},
					{
						Name:        "storage_account_access_key",
						Description: "Specifies the identifier key of the Threat Detection audit storage account.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("StorageAccountAccessKey"),
					},
					{
						Name:        "retention_days",
						Description: "Specifies the number of days to keep in the Threat Detection audit logs.",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("RetentionDays"),
					},
					{
						Name:        "creation_time",
						Description: "Specifies the UTC creation time of the policy.",
						Type:        schema.TypeTimestamp,
						Resolver:    schema.PathResolver("CreationTime.Time"),
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchSqlServers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().SQL.Servers
	servers, err := svc.List(ctx)
	if err != nil {
		return diag.WrapError(err)
	}
	for servers.NotDone() {
		res <- servers.Values()
		if err := servers.NextWithContext(ctx); err != nil {
			return diag.WrapError(err)
		}
	}
	return nil
}

func fetchSqlServerPrivateEndpointConnections(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	server := parent.Item.(sql.Server)
	if server.PrivateEndpointConnections != nil {
		res <- *server.PrivateEndpointConnections
	}
	return nil
}

func fetchSqlServerVirtualNetworkRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().SQL.VirtualNetworkRules
	server := parent.Item.(sql.Server)
	details, err := client.ParseResourceID(*server.ID)
	if err != nil {
		return diag.WrapError(err)
	}
	result, err := svc.ListByServer(ctx, details.ResourceGroup, *server.Name)
	if err != nil {
		return diag.WrapError(err)
	}
	for result.NotDone() {
		res <- result.Values()
		if err := result.NextWithContext(ctx); err != nil {
			return diag.WrapError(err)
		}
	}
	return nil
}

func fetchSqlServerFirewallRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().SQL.Firewall
	server := parent.Item.(sql.Server)
	details, err := client.ParseResourceID(*server.ID)
	if err != nil {
		return diag.WrapError(err)
	}
	result, err := svc.ListByServer(ctx, details.ResourceGroup, *server.Name)
	if err != nil {
		return diag.WrapError(err)
	}
	if result.Value != nil {
		res <- *result.Value
	}
	return nil
}

func fetchSqlServerAdmins(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().SQL.ServerAdmins
	server := parent.Item.(sql.Server)
	details, err := client.ParseResourceID(*server.ID)
	if err != nil {
		return diag.WrapError(err)
	}
	result, err := svc.ListByServer(ctx, details.ResourceGroup, *server.Name)
	if err != nil {
		return diag.WrapError(err)
	}
	for result.NotDone() {
		res <- result.Values()
		if err := result.NextWithContext(ctx); err != nil {
			return diag.WrapError(err)
		}
	}
	return nil
}

func fetchSqlServerDbBlobAuditingPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().SQL.ServerBlobAuditingPolicies
	s := parent.Item.(sql.Server)
	details, err := client.ParseResourceID(*s.ID)
	if err != nil {
		return diag.WrapError(err)
	}
	result, err := svc.ListByServer(ctx, details.ResourceGroup, *s.Name)
	if err != nil {
		return diag.WrapError(err)
	}
	for result.NotDone() {
		res <- result.Values()
		if err := result.NextWithContext(ctx); err != nil {
			return diag.WrapError(err)
		}
	}
	return nil
}

func fetchSqlServerDevopsAuditSettings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().SQL.ServerDevOpsAuditSettings
	s := parent.Item.(sql.Server)
	details, err := client.ParseResourceID(*s.ID)
	if err != nil {
		return diag.WrapError(err)
	}
	result, err := svc.ListByServer(ctx, details.ResourceGroup, *s.Name)
	if err != nil {
		return diag.WrapError(err)
	}
	for result.NotDone() {
		res <- result.Values()
		if err := result.NextWithContext(ctx); err != nil {
			return diag.WrapError(err)
		}
	}
	return nil
}

func fetchSqlServerVulnerabilityAssessments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().SQL.ServerVulnerabilityAssessments
	s := parent.Item.(sql.Server)
	details, err := client.ParseResourceID(*s.ID)
	if err != nil {
		return diag.WrapError(err)
	}
	result, err := svc.ListByServer(ctx, details.ResourceGroup, *s.Name)
	if err != nil {
		return diag.WrapError(err)
	}
	for result.NotDone() {
		res <- result.Values()
		if err := result.NextWithContext(ctx); err != nil {
			return diag.WrapError(err)
		}
	}
	return nil
}

func fetchSqlServerSecurityAlertPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().SQL.ServerSecurityAlertPolicies
	server := parent.Item.(sql.Server)
	details, err := client.ParseResourceID(*server.ID)
	if err != nil {
		return diag.WrapError(err)
	}
	result, err := svc.ListByServer(ctx, details.ResourceGroup, *server.Name)
	if err != nil {
		return diag.WrapError(err)
	}
	for result.NotDone() {
		res <- result.Values()
		if err = result.NextWithContext(ctx); err != nil {
			return diag.WrapError(err)
		}
	}
	return nil
}
