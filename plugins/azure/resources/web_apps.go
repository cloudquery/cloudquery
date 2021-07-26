package resources

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/web/mgmt/2020-12-01/web"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func WebApps() *schema.Table {
	return &schema.Table{
		Name:         "azure_web_apps",
		Description:  "Site a web app, a mobile app backend, or an API app",
		Resolver:     fetchWebApps,
		Multiplex:    client.SubscriptionMultiplex,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"subscription_id", "id"}},
		DeleteFilter: client.DeleteSubscriptionFilter,
		Columns: []schema.Column{
			{
				Name:        "subscription_id",
				Description: "Azure subscription id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAzureSubscription,
			},
			{
				Name:        "state",
				Description: "Current state of the app",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.State"),
			},
			{
				Name:        "host_names",
				Description: "Hostnames associated with the app",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("SiteProperties.HostNames"),
			},
			{
				Name:        "repository_site_name",
				Description: "Name of the repository site",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.RepositorySiteName"),
			},
			{
				Name:        "usage_state",
				Description: "State indicating whether the app has exceeded its quota usage Read-only Possible values include: 'UsageStateNormal', 'UsageStateExceeded'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.UsageState"),
			},
			{
				Name:        "enabled",
				Description: "otherwise, <code>false</code> Setting this value to false disables the app (takes the app offline)",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("SiteProperties.Enabled"),
			},
			{
				Name:        "enabled_host_names",
				Description: "Enabled hostnames for the appHostnames need to be assigned (see HostNames) AND enabled Otherwise, the app is not served on those hostnames",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("SiteProperties.EnabledHostNames"),
			},
			{
				Name:        "availability_state",
				Description: "Management information availability state for the app Possible values include: 'Normal', 'Limited', 'DisasterRecoveryMode'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.AvailabilityState"),
			},
			{
				Name:        "server_farm_id",
				Description: "Resource ID of the associated App Service plan, formatted as: \"/subscriptions/{subscriptionID}/resourceGroups/{groupName}/providers/MicrosoftWeb/serverfarms/{appServicePlanName}\"",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.ServerFarmID"),
			},
			{
				Name:        "reserved",
				Description: "otherwise, <code>false</code>",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("SiteProperties.Reserved"),
			},
			{
				Name:        "is_xenon",
				Description: "Obsolete: Hyper-V sandbox",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("SiteProperties.IsXenon"),
			},
			{
				Name:        "hyper_v",
				Description: "Hyper-V sandbox",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("SiteProperties.HyperV"),
			},
			{
				Name:     "last_modified_time_utc_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("SiteProperties.LastModifiedTimeUtc.Time"),
			},
			{
				Name:        "site_config",
				Description: "Configuration of the app",
				Type:        schema.TypeJSON,
				Resolver:    resolveWebAppSiteConfig,
			},
			{
				Name:        "traffic_manager_host_names",
				Description: "Azure Traffic Manager hostnames associated with the app Read-only",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("SiteProperties.TrafficManagerHostNames"),
			},
			{
				Name:        "scm_site_also_stopped",
				Description: "otherwise, <code>false</code> The default is <code>false</code>",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("SiteProperties.ScmSiteAlsoStopped"),
			},
			{
				Name:        "target_swap_slot",
				Description: "Specifies which deployment slot this app will swap into Read-only",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.TargetSwapSlot"),
			},
			{
				Name:        "hosting_environment_profile_id",
				Description: "Resource ID of the App Service Environment",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.HostingEnvironmentProfile.ID"),
			},
			{
				Name:        "hosting_environment_profile_name",
				Description: "Name of the App Service Environment",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.HostingEnvironmentProfile.Name"),
			},
			{
				Name:        "hosting_environment_profile_type",
				Description: "Resource type of the App Service Environment",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.HostingEnvironmentProfile.Type"),
			},
			{
				Name:        "client_affinity_enabled",
				Description: "<code>false</code> to stop sending session affinity cookies, which route client requests in the same session to the same instance Default is <code>true</code>",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("SiteProperties.ClientAffinityEnabled"),
			},
			{
				Name:        "client_cert_enabled",
				Description: "otherwise, <code>false</code> Default is <code>false</code>",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("SiteProperties.ClientCertEnabled"),
			},
			{
				Name:        "client_cert_mode",
				Description: "This composes with ClientCertEnabled setting - ClientCertEnabled: false means ClientCert is ignored - ClientCertEnabled: true and ClientCertMode: Required means ClientCert is required - ClientCertEnabled: true and ClientCertMode: Optional means ClientCert is optional or accepted Possible values include: 'Required', 'Optional', 'OptionalInteractiveUser'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.ClientCertMode"),
			},
			{
				Name:        "client_cert_exclusion_paths",
				Description: "client certificate authentication comma-separated exclusion paths",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.ClientCertExclusionPaths"),
			},
			{
				Name:        "host_names_disabled",
				Description: "otherwise, <code>false</code>  If <code>true</code>, the app is only accessible via API management process",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("SiteProperties.HostNamesDisabled"),
			},
			{
				Name:        "custom_domain_verification_id",
				Description: "Unique identifier that verifies the custom domains assigned to the app Customer will add this id to a txt record for verification",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.CustomDomainVerificationID"),
			},
			{
				Name:        "outbound_ip_addresses",
				Description: "List of IP addresses that the app uses for outbound connections (eg database access) Includes VIPs from tenants that site can be hosted with current settings Read-only",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.OutboundIPAddresses"),
			},
			{
				Name:        "possible_outbound_ip_addresses",
				Description: "List of IP addresses that the app uses for outbound connections (eg database access) Includes VIPs from all tenants except dataComponent Read-only",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.PossibleOutboundIPAddresses"),
			},
			{
				Name:        "container_size",
				Description: "Size of the function container",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("SiteProperties.ContainerSize"),
			},
			{
				Name:        "daily_memory_time_quota",
				Description: "Maximum allowed daily memory-time quota (applicable on dynamic apps only)",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("SiteProperties.DailyMemoryTimeQuota"),
			},
			{
				Name:     "suspended_till_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("SiteProperties.SuspendedTill.Time"),
			},
			{
				Name:        "max_number_of_workers",
				Description: "Maximum number of workers This only applies to Functions container",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("SiteProperties.MaxNumberOfWorkers"),
			},
			{
				Name:        "cloning_info_correlation_id",
				Description: "Correlation ID of cloning operation This ID ties multiple cloning operations together to use the same snapshot",
				Type:        schema.TypeUUID,
				Resolver:    schema.PathResolver("SiteProperties.CloningInfo.CorrelationID"),
			},
			{
				Name:        "cloning_info_overwrite",
				Description: "otherwise, <code>false</code>",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("SiteProperties.CloningInfo.Overwrite"),
			},
			{
				Name:        "cloning_info_clone_custom_host_names",
				Description: "otherwise, <code>false</code>",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("SiteProperties.CloningInfo.CloneCustomHostNames"),
			},
			{
				Name:        "cloning_info_clone_source_control",
				Description: "otherwise, <code>false</code>",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("SiteProperties.CloningInfo.CloneSourceControl"),
			},
			{
				Name:        "cloning_info_source_web_app_id",
				Description: "ARM resource ID of the source app App resource ID is of the form /subscriptions/{subId}/resourceGroups/{resourceGroupName}/providers/MicrosoftWeb/sites/{siteName} for production slots and /subscriptions/{subId}/resourceGroups/{resourceGroupName}/providers/MicrosoftWeb/sites/{siteName}/slots/{slotName} for other slots",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.CloningInfo.SourceWebAppID"),
			},
			{
				Name:        "cloning_info_source_web_app_location",
				Description: "Location of source app ex: West US or North Europe",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.CloningInfo.SourceWebAppLocation"),
			},
			{
				Name:        "cloning_info_hosting_environment",
				Description: "App Service Environment",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.CloningInfo.HostingEnvironment"),
			},
			{
				Name:        "cloning_info_app_settings_overrides",
				Description: "Application setting overrides for cloned app If specified, these settings override the settings cloned from source app Otherwise, application settings from source app are retained",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("SiteProperties.CloningInfo.AppSettingsOverrides"),
			},
			{
				Name:        "cloning_info_configure_load_balancing",
				Description: "<code>true</code> to configure load balancing for source and destination app",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("SiteProperties.CloningInfo.ConfigureLoadBalancing"),
			},
			{
				Name:        "cloning_info_traffic_manager_profile_id",
				Description: "ARM resource ID of the Traffic Manager profile to use, if it exists Traffic Manager resource ID is of the form /subscriptions/{subId}/resourceGroups/{resourceGroupName}/providers/MicrosoftNetwork/trafficManagerProfiles/{profileName}",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.CloningInfo.TrafficManagerProfileID"),
			},
			{
				Name:        "cloning_info_traffic_manager_profile_name",
				Description: "Name of Traffic Manager profile to create This is only needed if Traffic Manager profile does not already exist",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.CloningInfo.TrafficManagerProfileName"),
			},
			{
				Name:        "resource_group",
				Description: "Name of the resource group the app belongs to Read-only",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.ResourceGroup"),
			},
			{
				Name:        "is_default_container",
				Description: "<code>true</code> if the app is a default container; otherwise, <code>false</code>",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("SiteProperties.IsDefaultContainer"),
			},
			{
				Name:        "default_host_name",
				Description: "Default hostname of the app Read-only",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.DefaultHostName"),
			},
			{
				Name:     "slot_swap_status_timestamp_utc_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("SiteProperties.SlotSwapStatus.TimestampUtc.Time"),
			},
			{
				Name:        "slot_swap_status_source_slot_name",
				Description: "The source slot of the last swap operation",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.SlotSwapStatus.SourceSlotName"),
			},
			{
				Name:        "slot_swap_status_destination_slot_name",
				Description: "The destination slot of the last swap operation",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.SlotSwapStatus.DestinationSlotName"),
			},
			{
				Name:        "key_vault_reference_identity",
				Description: "Identity to use for Key Vault Reference authentication",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.KeyVaultReferenceIdentity"),
			},
			{
				Name:        "https_only",
				Description: "HttpsOnly: configures a web site to accept only https requests Issues redirect for http requests",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("SiteProperties.HTTPSOnly"),
			},
			{
				Name:        "redundancy_mode",
				Description: "Site redundancy mode Possible values include: 'RedundancyModeNone', 'RedundancyModeManual', 'RedundancyModeFailover', 'RedundancyModeActiveActive', 'RedundancyModeGeoRedundant'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.RedundancyMode"),
			},
			{
				Name:        "in_progress_operation_id",
				Description: "Specifies an operation id if this site has a pending operation",
				Type:        schema.TypeUUID,
				Resolver:    schema.PathResolver("SiteProperties.InProgressOperationID"),
			},
			{
				Name:        "storage_account_required",
				Description: "Checks if Customer provided storage account is required",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("SiteProperties.StorageAccountRequired"),
			},
			{
				Name:        "identity_type",
				Description: "Type of managed service identity Possible values include: 'ManagedServiceIdentityTypeSystemAssigned', 'ManagedServiceIdentityTypeUserAssigned', 'ManagedServiceIdentityTypeSystemAssignedUserAssigned', 'ManagedServiceIdentityTypeNone'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Identity.Type"),
			},
			{
				Name:        "identity_tenant_id",
				Description: "Tenant of managed service identity",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Identity.TenantID"),
			},
			{
				Name:        "identity_principal_id",
				Description: "Principal Id of managed service identity",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Identity.PrincipalID"),
			},
			{
				Name:        "identity_user_assigned_identities",
				Description: "The list of user assigned identities associated with the resource The user identity dictionary key references will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/MicrosoftManagedIdentity/userAssignedIdentities/{identityName}",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Identity.UserAssignedIdentities"),
			},
			{
				Name:        "id",
				Description: "Resource Id",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "Resource Name",
				Type:        schema.TypeString,
			},
			{
				Name:        "kind",
				Description: "Kind of resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "location",
				Description: "Resource Location",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "Resource type",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "Resource tags",
				Type:        schema.TypeJSON,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "azure_web_app_host_name_ssl_states",
				Description: "HostNameSslState SSL-enabled hostname",
				Resolver:    fetchWebAppHostNameSslStates,
				Columns: []schema.Column{
					{
						Name:        "app_cq_id",
						Description: "Unique CloudQuery ID of azure_web_apps table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "name",
						Description: "Hostname",
						Type:        schema.TypeString,
					},
					{
						Name:        "ssl_state",
						Description: "SSL type Possible values include: 'SslStateDisabled', 'SslStateSniEnabled', 'SslStateIPBasedEnabled'",
						Type:        schema.TypeString,
					},
					{
						Name:        "virtual_ip",
						Description: "Virtual IP address assigned to the hostname if IP based SSL is enabled",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VirtualIP"),
					},
					{
						Name:        "thumbprint",
						Description: "SSL certificate thumbprint",
						Type:        schema.TypeString,
					},
					{
						Name:        "to_update",
						Description: "Set to <code>true</code> to update existing hostname",
						Type:        schema.TypeBool,
					},
					{
						Name:        "host_type",
						Description: "Indicates whether the hostname is a standard or repository hostname Possible values include: 'HostTypeStandard', 'HostTypeRepository'",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:     "aws_web_app_publishing_profiles",
				Resolver: fetchWebAppPublishingProfiles,
				Columns: []schema.Column{
					{
						Name:        "app_cq_id",
						Description: "Unique CloudQuery ID of azure_web_apps table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name: "publish_url",
						Type: schema.TypeString,
					},
					{
						Name: "user_name",
						Type: schema.TypeString,
					},
					{
						Name:     "user_pwd",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("UserPWD"),
					},
				},
			},
			WebAppAuthSettings(),
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchWebApps(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	svc := meta.(*client.Client).Services().Web.Apps
	response, err := svc.List(ctx)
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
func resolveWebAppSiteConfig(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r, ok := resource.Item.(web.Site)
	if !ok {
		return fmt.Errorf("expected web.Site but got %T", resource.Item)
	}

	if r.SiteConfig == nil {
		return nil
	}

	data, err := json.Marshal(r.SiteConfig)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, data)
}
func fetchWebAppHostNameSslStates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(web.Site)
	if !ok {
		return fmt.Errorf("expected web.Site but got %T", parent.Item)
	}

	if p.HostNameSslStates == nil {
		return nil
	}

	res <- *p.HostNameSslStates
	return nil
}
func fetchWebAppPublishingProfiles(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(web.Site)
	if !ok {
		return fmt.Errorf("expected web.Site but got %T", parent.Item)
	}

	svc := meta.(*client.Client).Services().Web.Apps
	response, err := svc.ListPublishingProfileXMLWithSecrets(ctx, *p.ResourceGroup, *p.Name, web.CsmPublishingProfileOptions{})
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	if _, err = buf.ReadFrom(response.Body); err != nil {
		return err
	}
	var profileData PublishData
	if err = xml.Unmarshal(buf.Bytes(), &profileData); err != nil {
		return err
	}

	res <- profileData.PublishData
	return nil
}

// ====================================================================================================================
//                                                  User Defined Helpers
// ====================================================================================================================

type PublishProfile struct {
	PublishUrl string `xml:"publishUrl,attr"`
	UserName   string `xml:"userName,attr"`
	UserPWD    string `xml:"userPWD,attr"`
}
type PublishData struct {
	XMLName     xml.Name         `xml:"publishData"`
	PublishData []PublishProfile `xml:"publishProfile"`
}
