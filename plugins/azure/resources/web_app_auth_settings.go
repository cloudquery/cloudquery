package resources

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/web/mgmt/2020-12-01/web"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func WebAppAuthSettings() *schema.Table {
	return &schema.Table{
		Name:        "azure_web_app_auth_settings",
		Description: "SiteAuthSettings configuration settings for the Azure App Service Authentication / Authorization feature",
		Resolver:    fetchWebAppAuthSettings,
		Options:     schema.TableCreationOptions{PrimaryKeys: []string{"app_cq_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "app_cq_id",
				Description: "Unique CloudQuery ID of azure_web_apps table (FK)",
				Type:        schema.TypeUUID,
				Resolver:    schema.ParentIdResolver,
			},
			{
				Name:        "app_id",
				Description: "Original resource id of the web app (FK)",
				Type:        schema.TypeString,
				Resolver:    schema.ParentResourceFieldResolver("id"),
			},
			{
				Name:        "enabled",
				Description: "If authorization for site is enabled the value is true",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("SiteAuthSettingsProperties.Enabled"),
			},
			{
				Name:        "runtime_version",
				Description: "The RuntimeVersion of the Authentication / Authorization feature in use for the current app The setting in this value can control the behavior of certain features in the Authentication / Authorization module",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteAuthSettingsProperties.RuntimeVersion"),
			},
			{
				Name:        "config_version",
				Description: "The ConfigVersion of the Authentication / Authorization feature in use for the current app The setting in this value can control the behavior of the control plane for Authentication / Authorization",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteAuthSettingsProperties.ConfigVersion"),
			},
			{
				Name:        "unauthenticated_client_action",
				Description: "The action to take when an unauthenticated client attempts to access the app Possible values include: 'RedirectToLoginPage', 'AllowAnonymous'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteAuthSettingsProperties.UnauthenticatedClientAction"),
			},
			{
				Name:        "token_store_enabled",
				Description: "otherwise, <code>false</code>  The default is <code>false</code>",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("SiteAuthSettingsProperties.TokenStoreEnabled"),
			},
			{
				Name:        "allowed_external_redirect_urls",
				Description: "External URLs that can be redirected to as part of logging in or logging out of the app Note that the query string part of the URL is ignored This is an advanced setting typically only needed by Windows Store application backends Note that URLs within the current domain are always implicitly allowed",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("SiteAuthSettingsProperties.AllowedExternalRedirectUrls"),
			},
			{
				Name:        "default_provider",
				Description: "The default authentication provider to use when multiple providers are configured This setting is only needed if multiple providers are configured and the unauthenticated client action is set to \"RedirectToLoginPage\" Possible values include: 'BuiltInAuthenticationProviderAzureActiveDirectory', 'BuiltInAuthenticationProviderFacebook', 'BuiltInAuthenticationProviderGoogle', 'BuiltInAuthenticationProviderMicrosoftAccount', 'BuiltInAuthenticationProviderTwitter', 'BuiltInAuthenticationProviderGithub'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteAuthSettingsProperties.DefaultProvider"),
			},
			{
				Name:        "token_refresh_extension_hours",
				Description: "The number of hours after session token expiration that a session token can be used to call the token refresh API The default is 72 hours",
				Type:        schema.TypeFloat,
				Resolver:    schema.PathResolver("SiteAuthSettingsProperties.TokenRefreshExtensionHours"),
			},
			{
				Name:        "client_id",
				Description: "The Client ID of this relying party application, known as the client_id This setting is required for enabling OpenID Connection authentication with Azure Active Directory or other 3rd party OpenID Connect providers More information on OpenID Connect: http://openidnet/specs/openid-connect-core-1_0html",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteAuthSettingsProperties.ClientID"),
			},
			{
				Name:        "client_secret",
				Description: "The Client Secret of this relying party application (in Azure Active Directory, this is also referred to as the Key) This setting is optional If no client secret is configured, the OpenID Connect implicit auth flow is used to authenticate end users Otherwise, the OpenID Connect Authorization Code Flow is used to authenticate end users More information on OpenID Connect: http://openidnet/specs/openid-connect-core-1_0html",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteAuthSettingsProperties.ClientSecret"),
			},
			{
				Name:        "client_secret_setting_name",
				Description: "The app setting name that contains the client secret of the relying party application",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteAuthSettingsProperties.ClientSecretSettingName"),
			},
			{
				Name:        "client_secret_certificate_thumbprint",
				Description: "An alternative to the client secret, that is the thumbprint of a certificate used for signing purposes This property acts as a replacement for the Client Secret It is also optional",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteAuthSettingsProperties.ClientSecretCertificateThumbprint"),
			},
			{
				Name:        "issuer",
				Description: "The OpenID Connect Issuer URI that represents the entity which issues access tokens for this application When using Azure Active Directory, this value is the URI of the directory tenant, eg https://stswindowsnet/{tenant-guid}/ This URI is a case-sensitive identifier for the token issuer More information on OpenID Connect Discovery: http://openidnet/specs/openid-connect-discovery-1_0html",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteAuthSettingsProperties.Issuer"),
			},
			{
				Name:        "validate_issuer",
				Description: "Gets a value indicating whether the issuer should be a valid HTTPS url and be validated as such",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("SiteAuthSettingsProperties.ValidateIssuer"),
			},
			{
				Name:        "allowed_audiences",
				Description: "Allowed audience values to consider when validating JWTs issued by Azure Active Directory Note that the <code>ClientID</code> value is always considered an allowed audience, regardless of this setting",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("SiteAuthSettingsProperties.AllowedAudiences"),
			},
			{
				Name:        "additional_login_params",
				Description: "Login parameters to send to the OpenID Connect authorization endpoint when a user logs in Each parameter must be in the form \"key=value\"",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("SiteAuthSettingsProperties.AdditionalLoginParams"),
			},
			{
				Name:        "aad_claims_authorization",
				Description: "Gets a JSON string containing the Azure AD Acl settings",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteAuthSettingsProperties.AadClaimsAuthorization"),
			},
			{
				Name:        "google_client_id",
				Description: "The OpenID Connect Client ID for the Google web application This setting is required for enabling Google Sign-In Google Sign-In documentation: https://developersgooglecom/identity/sign-in/web/",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteAuthSettingsProperties.GoogleClientID"),
			},
			{
				Name:        "google_client_secret",
				Description: "The client secret associated with the Google web application This setting is required for enabling Google Sign-In Google Sign-In documentation: https://developersgooglecom/identity/sign-in/web/",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteAuthSettingsProperties.GoogleClientSecret"),
			},
			{
				Name:        "google_client_secret_setting_name",
				Description: "The app setting name that contains the client secret associated with the Google web application",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteAuthSettingsProperties.GoogleClientSecretSettingName"),
			},
			{
				Name:        "google_oauth_scopes",
				Description: "The OAuth 20 scopes that will be requested as part of Google Sign-In authentication This setting is optional If not specified, \"openid\", \"profile\", and \"email\" are used as default scopes Google Sign-In documentation: https://developersgooglecom/identity/sign-in/web/",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("SiteAuthSettingsProperties.GoogleOAuthScopes"),
			},
			{
				Name:        "facebook_app_id",
				Description: "The App ID of the Facebook app used for login This setting is required for enabling Facebook Login Facebook Login documentation: https://developersfacebookcom/docs/facebook-login",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteAuthSettingsProperties.FacebookAppID"),
			},
			{
				Name:        "facebook_app_secret",
				Description: "The App Secret of the Facebook app used for Facebook Login This setting is required for enabling Facebook Login Facebook Login documentation: https://developersfacebookcom/docs/facebook-login",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteAuthSettingsProperties.FacebookAppSecret"),
			},
			{
				Name:        "facebook_app_secret_setting_name",
				Description: "The app setting name that contains the app secret used for Facebook Login",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteAuthSettingsProperties.FacebookAppSecretSettingName"),
			},
			{
				Name:        "facebook_oauth_scopes",
				Description: "The OAuth 20 scopes that will be requested as part of Facebook Login authentication This setting is optional Facebook Login documentation: https://developersfacebookcom/docs/facebook-login",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("SiteAuthSettingsProperties.FacebookOAuthScopes"),
			},
			{
				Name:        "git_hub_client_id",
				Description: "The Client Id of the GitHub app used for login This setting is required for enabling Github login",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteAuthSettingsProperties.GitHubClientID"),
			},
			{
				Name:        "git_hub_client_secret",
				Description: "The Client Secret of the GitHub app used for Github Login This setting is required for enabling Github login",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteAuthSettingsProperties.GitHubClientSecret"),
			},
			{
				Name:        "git_hub_client_secret_setting_name",
				Description: "The app setting name that contains the client secret of the Github app used for GitHub Login",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteAuthSettingsProperties.GitHubClientSecretSettingName"),
			},
			{
				Name:        "git_hub_oauth_scopes",
				Description: "The OAuth 20 scopes that will be requested as part of GitHub Login authentication This setting is optional",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("SiteAuthSettingsProperties.GitHubOAuthScopes"),
			},
			{
				Name:        "twitter_consumer_key",
				Description: "The OAuth 10a consumer key of the Twitter application used for sign-in This setting is required for enabling Twitter Sign-In Twitter Sign-In documentation: https://devtwittercom/web/sign-in",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteAuthSettingsProperties.TwitterConsumerKey"),
			},
			{
				Name:        "twitter_consumer_secret",
				Description: "The OAuth 10a consumer secret of the Twitter application used for sign-in This setting is required for enabling Twitter Sign-In Twitter Sign-In documentation: https://devtwittercom/web/sign-in",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteAuthSettingsProperties.TwitterConsumerSecret"),
			},
			{
				Name:        "twitter_consumer_secret_setting_name",
				Description: "The app setting name that contains the OAuth 10a consumer secret of the Twitter application used for sign-in",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteAuthSettingsProperties.TwitterConsumerSecretSettingName"),
			},
			{
				Name:        "microsoft_account_client_id",
				Description: "The OAuth 20 client ID that was created for the app used for authentication This setting is required for enabling Microsoft Account authentication Microsoft Account OAuth documentation: https://devonedrivecom/auth/msa_oauthhtm",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteAuthSettingsProperties.MicrosoftAccountClientID"),
			},
			{
				Name:        "microsoft_account_client_secret",
				Description: "The OAuth 20 client secret that was created for the app used for authentication This setting is required for enabling Microsoft Account authentication Microsoft Account OAuth documentation: https://devonedrivecom/auth/msa_oauthhtm",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteAuthSettingsProperties.MicrosoftAccountClientSecret"),
			},
			{
				Name:        "microsoft_account_client_secret_setting_name",
				Description: "The app setting name containing the OAuth 20 client secret that was created for the app used for authentication",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteAuthSettingsProperties.MicrosoftAccountClientSecretSettingName"),
			},
			{
				Name:        "microsoft_account_oauth_scopes",
				Description: "The OAuth 20 scopes that will be requested as part of Microsoft Account authentication This setting is optional If not specified, \"wlbasic\" is used as the default scope Microsoft Account Scopes and permissions documentation: https://msdnmicrosoftcom/en-us/library/dn631845aspx",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("SiteAuthSettingsProperties.MicrosoftAccountOAuthScopes"),
			},
			{
				Name:        "is_auth_from_file",
				Description: "\"true\" if the auth config settings should be read from a file, \"false\" otherwise",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteAuthSettingsProperties.IsAuthFromFile"),
			},
			{
				Name:        "auth_file_path",
				Description: "The path of the config file containing auth settings If the path is relative, base will the site's root directory",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteAuthSettingsProperties.AuthFilePath"),
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
				Name:        "type",
				Description: "Resource type",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchWebAppAuthSettings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(web.Site)
	if !ok {
		return fmt.Errorf("expected web.Site but got %T", parent.Item)
	}

	svc := meta.(*client.Client).Services().Web.Apps
	response, err := svc.GetAuthSettings(ctx, *p.ResourceGroup, *p.Name)
	if err != nil {
		return err
	}
	res <- response
	return nil
}
