package ad

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Applications() *schema.Table {
	return &schema.Table{
		Name:         "azure_ad_applications",
		Description:  "Application active Directory application information",
		Resolver:     fetchAdApplications,
		Multiplex:    client.SubscriptionMultiplex,
		DeleteFilter: client.DeleteSubscriptionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"subscription_id", "object_id"}},
		Columns: []schema.Column{
			{
				Name:        "subscription_id",
				Description: "Azure subscription id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAzureSubscription,
			},
			{
				Name:        "app_id",
				Description: "The application ID",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AppID"),
			},
			{
				Name:        "allow_guests_sign_in",
				Description: "A property on the application to indicate if the application accepts other IDPs or not or partially accepts",
				Type:        schema.TypeBool,
			},
			{
				Name:        "allow_passthrough_users",
				Description: "Indicates that the application supports pass through users who have no presence in the resource tenant",
				Type:        schema.TypeBool,
			},
			{
				Name:        "app_logo_url",
				Description: "The url for the application logo image stored in a CDN",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AppLogoURL"),
			},
			{
				Name:        "app_permissions",
				Description: "The application permissions",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "available_to_other_tenants",
				Description: "Whether the application is available to other tenants",
				Type:        schema.TypeBool,
			},
			{
				Name:        "display_name",
				Description: "The display name of the application",
				Type:        schema.TypeString,
			},
			{
				Name:        "error_url",
				Description: "A URL provided by the author of the application to report errors when using the application",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ErrorURL"),
			},
			{
				Name:        "group_membership_claims",
				Description: "Configures the groups claim issued in a user or OAuth 20 access token that the app expects Possible values include: 'None', 'SecurityGroup', 'All'",
				Type:        schema.TypeString,
			},
			{
				Name:        "homepage",
				Description: "The home page of the application",
				Type:        schema.TypeString,
			},
			{
				Name:        "identifier_uris",
				Description: "A collection of URIs for the application",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "informational_urls_terms_of_service",
				Description: "The terms of service URI",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("InformationalUrls.TermsOfService"),
			},
			{
				Name:        "informational_urls_marketing",
				Description: "The marketing URI",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("InformationalUrls.Marketing"),
			},
			{
				Name:        "informational_urls_privacy",
				Description: "The privacy policy URI",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("InformationalUrls.Privacy"),
			},
			{
				Name:        "informational_urls_support",
				Description: "The support URI",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("InformationalUrls.Support"),
			},
			{
				Name:        "is_device_only_auth_supported",
				Description: "Specifies whether this application supports device authentication without a user The default is false",
				Type:        schema.TypeBool,
			},
			{
				Name:        "known_client_applications",
				Description: "Client applications that are tied to this resource application Consent to any of the known client applications will result in implicit consent to the resource application through a combined consent dialog (showing the OAuth permission scopes required by the client and the resource)",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "logout_url",
				Description: "the url of the logout page",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LogoutURL"),
			},
			{
				Name:        "oauth2_allow_implicit_flow",
				Description: "Whether to allow implicit grant flow for OAuth2",
				Type:        schema.TypeBool,
			},
			{
				Name:        "oauth2_allow_url_path_matching",
				Description: "Specifies whether during a token Request Azure AD will allow path matching of the redirect URI against the applications collection of replyURLs The default is false",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Oauth2AllowURLPathMatching"),
			},
			{
				Name:        "oauth2_require_post_response",
				Description: "Specifies whether, as part of OAuth 20 token requests, Azure AD will allow POST requests, as opposed to GET requests The default is false, which specifies that only GET requests will be allowed",
				Type:        schema.TypeBool,
			},
			{
				Name:        "org_restrictions",
				Description: "A list of tenants allowed to access application",
				Type:        schema.TypeStringArray,
			},
			{
				Name:     "optional_claims",
				Type:     schema.TypeJSON,
				Resolver: resolveAdApplicationOptionalClaims,
			},
			{
				Name:        "public_client",
				Description: "Specifies whether this application is a public client (such as an installed application running on a mobile device) Default is false",
				Type:        schema.TypeBool,
			},
			{
				Name:        "publisher_domain",
				Description: "Reliable domain which can be used to identify an application",
				Type:        schema.TypeString,
			},
			{
				Name:        "reply_urls",
				Description: "A collection of reply URLs for the application",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "saml_metadata_url",
				Description: "The URL to the SAML metadata for the application",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SamlMetadataURL"),
			},
			{
				Name:        "sign_in_audience",
				Description: "Audience for signing in to the application (AzureADMyOrganization, AzureADAllOrganizations, AzureADAndMicrosoftAccounts)",
				Type:        schema.TypeString,
			},
			{
				Name:        "www_homepage",
				Description: "The primary Web page",
				Type:        schema.TypeString,
			},
			{
				Name:        "additional_properties",
				Description: "Unmatched properties from the message are deserialized this collection",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "object_id",
				Description: "The object ID",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ObjectID"),
			},
			{
				Name:     "deletion_timestamp_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("DeletionTimestamp.Time"),
			},
			{
				Name:        "object_type",
				Description: "Possible values include: 'ObjectTypeDirectoryObject', 'ObjectTypeApplication', 'ObjectTypeGroup', 'ObjectTypeServicePrincipal', 'ObjectTypeUser'",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "azure_ad_application_app_roles",
				Description: "AppRole",
				Resolver:    fetchAdApplicationAppRoles,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"application_cq_id", "id"}},
				Columns: []schema.Column{
					{
						Name:        "application_cq_id",
						Description: "Unique ID of azure_ad_applications table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "id",
						Description: "Unique role identifier inside the appRoles collection",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
					{
						Name:        "allowed_member_types",
						Description: "Specifies whether this app role definition can be assigned to users and groups by setting to 'User', or to other applications (that are accessing this application in daemon service scenarios) by setting to 'Application', or to both",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "description",
						Description: "Permission help text that appears in the admin app assignment and consent experiences",
						Type:        schema.TypeString,
					},
					{
						Name:        "display_name",
						Description: "Display name for the permission that appears in the admin consent and app assignment experiences",
						Type:        schema.TypeString,
					},
					{
						Name:        "is_enabled",
						Description: "When creating or updating a role definition, this must be set to true (which is the default) To delete a role, this must first be set to false At that point, in a subsequent call, this role may be removed",
						Type:        schema.TypeBool,
					},
					{
						Name:        "role_claim_value",
						Description: "Specifies the value of the roles claim that the application should expect in the authentication and access tokens",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Value"),
					},
				},
			},
			{
				Name:        "azure_ad_application_key_credentials",
				Description: "KeyCredential active Directory Key Credential information",
				Resolver:    fetchAdApplicationKeyCredentials,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"application_cq_id", "key_id"}},
				Columns: []schema.Column{
					{
						Name:        "application_cq_id",
						Description: "Unique ID of azure_ad_applications table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "additional_properties",
						Description: "Unmatched properties from the message are deserialized this collection",
						Type:        schema.TypeJSON,
					},
					{
						Name:     "start_date_time",
						Type:     schema.TypeTimestamp,
						Resolver: schema.PathResolver("StartDate.Time"),
					},
					{
						Name:     "end_date_time",
						Type:     schema.TypeTimestamp,
						Resolver: schema.PathResolver("EndDate.Time"),
					},
					{
						Name:        "value",
						Description: "Key value",
						Type:        schema.TypeString,
					},
					{
						Name:        "key_id",
						Description: "Key ID",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("KeyID"),
					},
					{
						Name:        "usage",
						Description: "Usage Acceptable values are 'Verify' and 'Sign'",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "Type Acceptable values are 'AsymmetricX509Cert' and 'Symmetric'",
						Type:        schema.TypeString,
					},
					{
						Name:        "custom_key_identifier",
						Description: "Custom Key Identifier",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "azure_ad_application_oauth2_permissions",
				Description: "OAuth2Permission represents an OAuth 20 delegated permission scope The specified OAuth 20 delegated permission scopes may be requested by client applications (through the requiredResourceAccess collection on the Application object) when calling a resource application The oauth2Permissions property of the ServicePrincipal entity and of the Application entity is a collection of OAuth2Permission",
				Resolver:    fetchAdApplicationOauth2Permissions,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"application_cq_id", "id"}},
				Columns: []schema.Column{
					{
						Name:        "application_cq_id",
						Description: "Unique ID of azure_ad_applications table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "admin_consent_description",
						Description: "Permission help text that appears in the admin consent and app assignment experiences",
						Type:        schema.TypeString,
					},
					{
						Name:        "admin_consent_display_name",
						Description: "Display name for the permission that appears in the admin consent and app assignment experiences",
						Type:        schema.TypeString,
					},
					{
						Name:        "id",
						Description: "Unique scope permission identifier inside the oauth2Permissions collection",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
					{
						Name:        "is_enabled",
						Description: "When creating or updating a permission, this property must be set to true (which is the default) To delete a permission, this property must first be set to false At that point, in a subsequent call, the permission may be removed",
						Type:        schema.TypeBool,
					},
					{
						Name:        "permission_type",
						Description: "Specifies whether this scope permission can be consented to by an end user, or whether it is a tenant-wide permission that must be consented to by a Company Administrator Possible values are \"User\" or \"Admin\"",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Type"),
					},
					{
						Name:        "user_consent_description",
						Description: "Permission help text that appears in the end user consent experience",
						Type:        schema.TypeString,
					},
					{
						Name:        "user_consent_display_name",
						Description: "Display name for the permission that appears in the end user consent experience",
						Type:        schema.TypeString,
					},
					{
						Name:        "scope_claim_value",
						Description: "The value of the scope claim that the resource application should expect in the OAuth 20 access token",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Value"),
					},
				},
			},
			{
				Name:        "azure_ad_application_password_credentials",
				Description: "PasswordCredential active Directory Password Credential information",
				Resolver:    fetchAdApplicationPasswordCredentials,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"application_cq_id", "key_id"}},
				Columns: []schema.Column{
					{
						Name:        "application_cq_id",
						Description: "Unique ID of azure_ad_applications table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "additional_properties",
						Description: "Unmatched properties from the message are deserialized this collection",
						Type:        schema.TypeJSON,
					},
					{
						Name:     "start_date_time",
						Type:     schema.TypeTimestamp,
						Resolver: schema.PathResolver("StartDate.Time"),
					},
					{
						Name:     "end_date_time",
						Type:     schema.TypeTimestamp,
						Resolver: schema.PathResolver("EndDate.Time"),
					},
					{
						Name:        "key_id",
						Description: "Key ID",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("KeyID"),
					},
					{
						Name:        "value",
						Description: "Key value",
						Type:        schema.TypeString,
					},
					{
						Name:        "custom_key_identifier",
						Description: "Custom Key Identifier",
						Type:        schema.TypeByteArray,
					},
				},
			},
			{
				Name:        "azure_ad_application_pre_authorized_applications",
				Description: "PreAuthorizedApplication contains information about pre authorized client application",
				Resolver:    fetchAdApplicationPreAuthorizedApplications,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"application_cq_id", "app_id"}},
				Columns: []schema.Column{
					{
						Name:        "application_cq_id",
						Description: "Unique ID of azure_ad_applications table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "app_id",
						Description: "Represents the application id",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("AppID"),
					},
					{
						Name:        "permissions",
						Description: "Collection of required app permissions/entitlements from the resource application",
						Type:        schema.TypeJSON,
						Resolver:    resolveAdApplicationPreAuthorizedApplicationPermissions,
					},
					{
						Name:        "extensions",
						Description: "Collection of extensions from the resource application",
						Type:        schema.TypeJSON,
						Resolver:    resolveAdApplicationPreAuthorizedApplicationExtensions,
					},
				},
			},
			{
				Name:        "azure_ad_application_required_resource_accesses",
				Description: "RequiredResourceAccess specifies the set of OAuth 20 permission scopes and app roles under the specified resource that an application requires access to The specified OAuth 20 permission scopes may be requested by client applications (through the requiredResourceAccess collection) when calling a resource application The requiredResourceAccess property of the Application entity is a collection of RequiredResourceAccess",
				Resolver:    fetchAdApplicationRequiredResourceAccesses,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"application_cq_id", "resource_app_id"}},
				Columns: []schema.Column{
					{
						Name:        "application_cq_id",
						Description: "Unique ID of azure_ad_applications table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "additional_properties",
						Description: "Unmatched properties from the message are deserialized this collection",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "resource_access",
						Description: "The list of OAuth20 permission scopes and app roles that the application requires from the specified resource",
						Type:        schema.TypeJSON,
						Resolver:    resolveAdApplicationRequiredResourceAccessResourceAccess,
					},
					{
						Name:        "resource_app_id",
						Description: "The unique identifier for the resource that the application requires access to This should be equal to the appId declared on the target resource application",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ResourceAppID"),
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchAdApplications(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().AD.Applications
	response, err := svc.List(ctx, "")
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

func resolveAdApplicationOptionalClaims(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	app, ok := resource.Item.(graphrbac.Application)
	if !ok {
		return fmt.Errorf("not a graphrbac.Application instance: %#v", resource.Item)
	}
	out, err := json.Marshal(app.OptionalClaims)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, out)
}

func fetchAdApplicationAppRoles(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	app, ok := parent.Item.(graphrbac.Application)
	if !ok {
		return fmt.Errorf("not a graphrbac.Application instance: %#v", parent.Item)
	}
	if app.AppRoles == nil {
		return nil
	}
	for _, item := range *app.AppRoles {
		res <- item
	}
	return nil
}

func fetchAdApplicationKeyCredentials(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	app, ok := parent.Item.(graphrbac.Application)
	if !ok {
		return fmt.Errorf("not a graphrbac.Application instance: %#v", parent.Item)
	}
	if app.KeyCredentials == nil {
		return nil
	}
	for _, item := range *app.KeyCredentials {
		res <- item
	}
	return nil
}

func fetchAdApplicationOauth2Permissions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	app, ok := parent.Item.(graphrbac.Application)
	if !ok {
		return fmt.Errorf("not a graphrbac.Application instance: %#v", parent.Item)
	}
	if app.Oauth2Permissions == nil {
		return nil
	}
	for _, item := range *app.Oauth2Permissions {
		res <- item
	}
	return nil
}

func fetchAdApplicationPasswordCredentials(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	app, ok := parent.Item.(graphrbac.Application)
	if !ok {
		return fmt.Errorf("not a graphrbac.Application instance: %#v", parent.Item)
	}
	if app.PasswordCredentials == nil {
		return nil
	}
	for _, item := range *app.PasswordCredentials {
		res <- item
	}
	return nil
}

func fetchAdApplicationPreAuthorizedApplications(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	app, ok := parent.Item.(graphrbac.Application)
	if !ok {
		return fmt.Errorf("not a graphrbac.Application instance: %#v", parent.Item)
	}
	if app.PreAuthorizedApplications == nil {
		return nil
	}
	for _, item := range *app.PreAuthorizedApplications {
		res <- item
	}
	return nil
}

func resolveAdApplicationPreAuthorizedApplicationPermissions(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	paa, ok := resource.Item.(graphrbac.PreAuthorizedApplication)
	if !ok {
		return fmt.Errorf("not a graphrbac.PreAuthorizedApplication instance: %#v", resource.Item)
	}
	out, err := json.Marshal(paa.Permissions)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, out)
}

func resolveAdApplicationPreAuthorizedApplicationExtensions(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	paa, ok := resource.Item.(graphrbac.PreAuthorizedApplication)
	if !ok {
		return fmt.Errorf("not a graphrbac.PreAuthorizedApplication instance: %#v", resource.Item)
	}
	out, err := json.Marshal(paa.Extensions)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, out)
}

func fetchAdApplicationRequiredResourceAccesses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	app, ok := parent.Item.(graphrbac.Application)
	if !ok {
		return fmt.Errorf("not a graphrbac.Application instance: %#v", parent.Item)
	}
	if app.RequiredResourceAccess == nil {
		return nil
	}
	for _, item := range *app.RequiredResourceAccess {
		res <- item
	}
	return nil
}

func resolveAdApplicationRequiredResourceAccessResourceAccess(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	rra, ok := resource.Item.(graphrbac.RequiredResourceAccess)
	if !ok {
		return fmt.Errorf("not a graphrbac.RequiredResourceAccess instance: %#v", resource.Item)
	}
	out, err := json.Marshal(rra.ResourceAccess)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, out)
}
