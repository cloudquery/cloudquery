package resources

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func AdServicePrincipals() *schema.Table {
	return &schema.Table{
		Name:         "azure_ad_service_principals",
		Description:  "ServicePrincipal active Directory service principal information",
		Resolver:     fetchAdServicePrincipals,
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
				Name:        "account_enabled",
				Description: "whether or not the service principal account is enabled",
				Type:        schema.TypeBool,
			},
			{
				Name:        "alternative_names",
				Description: "alternative names",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "app_display_name",
				Description: "The display name exposed by the associated application",
				Type:        schema.TypeString,
			},
			{
				Name:        "app_id",
				Description: "The application ID",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AppID"),
			},
			{
				Name:        "app_owner_tenant_id",
				Description: "Application owner id",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AppOwnerTenantID"),
			},
			{
				Name:        "app_role_assignment_required",
				Description: "Specifies whether an AppRoleAssignment to a user or group is required before Azure AD will issue a user or access token to the application",
				Type:        schema.TypeBool,
			},
			{
				Name:        "display_name",
				Description: "The display name of the service principal",
				Type:        schema.TypeString,
			},
			{
				Name:        "error_url",
				Description: "A URL provided by the author of the associated application to report errors when using the application",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ErrorURL"),
			},
			{
				Name:        "homepage",
				Description: "The URL to the homepage of the associated application",
				Type:        schema.TypeString,
			},
			{
				Name:        "logout_url",
				Description: "A URL provided by the author of the associated application to logout",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LogoutURL"),
			},
			{
				Name:        "preferred_token_signing_key_thumbprint",
				Description: "The thumbprint of preferred certificate to sign the token",
				Type:        schema.TypeString,
			},
			{
				Name:        "publisher_name",
				Description: "The publisher's name of the associated application",
				Type:        schema.TypeString,
			},
			{
				Name:        "reply_urls",
				Description: "The URLs that user tokens are sent to for sign in with the associated application  The redirect URIs that the oAuth 20 authorization code and access tokens are sent to for the associated application",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "saml_metadata_url",
				Description: "The URL to the SAML metadata of the associated application",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SamlMetadataURL"),
			},
			{
				Name:        "service_principal_names",
				Description: "A collection of service principal names",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "service_principal_type",
				Description: "the type of the service principal",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "Optional list of tags that you can apply to your service principals Not nullable",
				Type:        schema.TypeStringArray,
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
				Name:        "deletion_timestamp_time",
				Description: "The time at which the directory object was deleted.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("DeletionTimestamp.Time"),
			},
			{
				Name:        "object_type",
				Description: "Possible values include: 'ObjectTypeDirectoryObject', 'ObjectTypeApplication', 'ObjectTypeGroup', 'ObjectTypeServicePrincipal', 'ObjectTypeUser'",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "azure_ad_service_principal_app_roles",
				Description: "AppRole",
				Resolver:    fetchAdServicePrincipalAppRoles,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"service_principal_cq_id", "id"}},
				Columns: []schema.Column{
					{
						Name:        "service_principal_cq_id",
						Description: "Unique ID of azure_ad_service_principals table (FK)",
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
				Name:        "azure_ad_service_principal_key_credentials",
				Description: "KeyCredential active Directory Key Credential information",
				Resolver:    fetchAdServicePrincipalKeyCredentials,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"service_principal_cq_id", "key_id"}},
				Columns: []schema.Column{
					{
						Name:        "service_principal_cq_id",
						Description: "Unique ID of azure_ad_service_principals table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "additional_properties",
						Description: "Unmatched properties from the message are deserialized this collection",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "start_date",
						Description: "Start date.",
						Type:        schema.TypeTimestamp,
						Resolver:    schema.PathResolver("StartDate.Time"),
					},
					{
						Name:        "end_date",
						Description: "End date.",
						Type:        schema.TypeTimestamp,
						Resolver:    schema.PathResolver("EndDate.Time"),
					},
					{
						Name:        "key_value",
						Description: "Key value",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Value"),
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
						Name:        "key_type",
						Description: "Type Acceptable values are 'AsymmetricX509Cert' and 'Symmetric'",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Type"),
					},
					{
						Name:        "custom_key_identifier",
						Description: "Custom Key Identifier",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "azure_ad_service_principal_oauth2_permissions",
				Description: "OAuth2Permission represents an OAuth 20 delegated permission scope The specified OAuth 20 delegated permission scopes may be requested by client applications (through the requiredResourceAccess collection on the Application object) when calling a resource application The oauth2Permissions property of the ServicePrincipal entity and of the Application entity is a collection of OAuth2Permission",
				Resolver:    fetchAdServicePrincipalOauth2Permissions,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"service_principal_cq_id", "id"}},
				Columns: []schema.Column{
					{
						Name:        "service_principal_cq_id",
						Description: "Unique ID of azure_ad_service_principals table (FK)",
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
				Name:        "azure_ad_service_principal_password_credentials",
				Description: "PasswordCredential active Directory Password Credential information",
				Resolver:    fetchAdServicePrincipalPasswordCredentials,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"service_principal_cq_id", "key_id"}},
				Columns: []schema.Column{
					{
						Name:        "service_principal_cq_id",
						Description: "Unique ID of azure_ad_service_principals table (FK)",
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
						Name:        "key_value",
						Description: "Key value",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Value"),
					},
					{
						Name:        "custom_key_identifier",
						Description: "Custom Key Identifier",
						Type:        schema.TypeByteArray,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchAdServicePrincipals(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	svc := meta.(*client.Client).Services().AD.ServicePrincipals
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

func fetchAdServicePrincipalAppRoles(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	sp, ok := parent.Item.(graphrbac.ServicePrincipal)
	if !ok {
		return fmt.Errorf("not a graphrbac.ServicePrincipal instance: %#v", parent.Item)
	}
	if sp.AppRoles != nil {
		for _, item := range *sp.AppRoles {
			res <- item
		}
	}
	return nil
}

func fetchAdServicePrincipalKeyCredentials(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	sp, ok := parent.Item.(graphrbac.ServicePrincipal)
	if !ok {
		return fmt.Errorf("not a graphrbac.ServicePrincipal instance: %#v", parent.Item)
	}
	if sp.KeyCredentials != nil {
		for _, item := range *sp.KeyCredentials {
			res <- item
		}
	}
	return nil
}

func fetchAdServicePrincipalOauth2Permissions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	sp, ok := parent.Item.(graphrbac.ServicePrincipal)
	if !ok {
		return fmt.Errorf("not a graphrbac.ServicePrincipal instance: %#v", parent.Item)
	}
	if sp.Oauth2Permissions != nil {
		for _, item := range *sp.Oauth2Permissions {
			res <- item
		}
	}
	return nil
}

func fetchAdServicePrincipalPasswordCredentials(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	sp, ok := parent.Item.(graphrbac.ServicePrincipal)
	if !ok {
		return fmt.Errorf("not a graphrbac.ServicePrincipal instance: %#v", parent.Item)
	}
	if sp.PasswordCredentials != nil {
		for _, item := range *sp.PasswordCredentials {
			res <- item
		}
	}
	return nil
}
