// Auto generated code - DO NOT EDIT.

package web

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/pkg/errors"

	"github.com/Azure/azure-sdk-for-go/services/web/mgmt/2020-12-01/web"
)

func siteAuthSettings() *schema.Table {
	return &schema.Table{
		Name:     "azure_web_site_auth_settings",
		Resolver: fetchWebSiteAuthSettings,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "cq_id_parent",
				Type:     schema.TypeUUID,
				Resolver: schema.ParentIDResolver,
			},
			{
				Name:     "enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Enabled"),
			},
			{
				Name:     "runtime_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RuntimeVersion"),
			},
			{
				Name:     "unauthenticated_client_action",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("UnauthenticatedClientAction"),
			},
			{
				Name:     "token_store_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("TokenStoreEnabled"),
			},
			{
				Name:     "allowed_external_redirect_urls",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("AllowedExternalRedirectUrls"),
			},
			{
				Name:     "default_provider",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DefaultProvider"),
			},
			{
				Name:     "token_refresh_extension_hours",
				Type:     schema.TypeFloat,
				Resolver: schema.PathResolver("TokenRefreshExtensionHours"),
			},
			{
				Name:     "client_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClientID"),
			},
			{
				Name:     "client_secret",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClientSecret"),
			},
			{
				Name:     "client_secret_setting_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClientSecretSettingName"),
			},
			{
				Name:     "client_secret_certificate_thumbprint",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClientSecretCertificateThumbprint"),
			},
			{
				Name:     "issuer",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Issuer"),
			},
			{
				Name:     "validate_issuer",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ValidateIssuer"),
			},
			{
				Name:     "allowed_audiences",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("AllowedAudiences"),
			},
			{
				Name:     "additional_login_params",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("AdditionalLoginParams"),
			},
			{
				Name:     "aad_claims_authorization",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AadClaimsAuthorization"),
			},
			{
				Name:     "google_client_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("GoogleClientID"),
			},
			{
				Name:     "google_client_secret",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("GoogleClientSecret"),
			},
			{
				Name:     "google_client_secret_setting_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("GoogleClientSecretSettingName"),
			},
			{
				Name:     "google_o_auth_scopes",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("GoogleOAuthScopes"),
			},
			{
				Name:     "facebook_app_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FacebookAppID"),
			},
			{
				Name:     "facebook_app_secret",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FacebookAppSecret"),
			},
			{
				Name:     "facebook_app_secret_setting_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FacebookAppSecretSettingName"),
			},
			{
				Name:     "facebook_o_auth_scopes",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("FacebookOAuthScopes"),
			},
			{
				Name:     "git_hub_client_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("GitHubClientID"),
			},
			{
				Name:     "git_hub_client_secret",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("GitHubClientSecret"),
			},
			{
				Name:     "git_hub_client_secret_setting_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("GitHubClientSecretSettingName"),
			},
			{
				Name:     "git_hub_o_auth_scopes",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("GitHubOAuthScopes"),
			},
			{
				Name:     "twitter_consumer_key",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TwitterConsumerKey"),
			},
			{
				Name:     "twitter_consumer_secret",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TwitterConsumerSecret"),
			},
			{
				Name:     "twitter_consumer_secret_setting_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TwitterConsumerSecretSettingName"),
			},
			{
				Name:     "microsoft_account_client_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MicrosoftAccountClientID"),
			},
			{
				Name:     "microsoft_account_client_secret",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MicrosoftAccountClientSecret"),
			},
			{
				Name:     "microsoft_account_client_secret_setting_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MicrosoftAccountClientSecretSettingName"),
			},
			{
				Name:     "microsoft_account_o_auth_scopes",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("MicrosoftAccountOAuthScopes"),
			},
			{
				Name:     "is_auth_from_file",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IsAuthFromFile"),
			},
			{
				Name:     "auth_file_path",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AuthFilePath"),
			},
			{
				Name:     "config_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ConfigVersion"),
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
				Name:     "kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Kind"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
		},
	}
}

func fetchWebSiteAuthSettings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Web.SiteAuthSettings

	site := parent.Item.(web.Site)
	response, err := svc.GetAuthSettings(ctx, *site.ResourceGroup, *site.Name)
	if err != nil {
		return errors.WithStack(err)
	}
	res <- response
	return nil
}
