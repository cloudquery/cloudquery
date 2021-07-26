
# Table: aws_web_app_auth_settings
SiteAuthSettings configuration settings for the Azure App Service Authentication / Authorization feature
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|app_cq_id|uuid|Unique CloudQuery ID of azure_web_apps table (FK)|
|app_id|text|Original resource id of the web app (FK)|
|enabled|boolean|If authorization for site is enabled the value is true|
|runtime_version|text|The RuntimeVersion of the Authentication / Authorization feature in use for the current app The setting in this value can control the behavior of certain features in the Authentication / Authorization module|
|config_version|text|The ConfigVersion of the Authentication / Authorization feature in use for the current app The setting in this value can control the behavior of the control plane for Authentication / Authorization|
|unauthenticated_client_action|text|The action to take when an unauthenticated client attempts to access the app Possible values include: 'RedirectToLoginPage', 'AllowAnonymous'|
|token_store_enabled|boolean|otherwise, <code>false</code>  The default is <code>false</code>|
|allowed_external_redirect_urls|text[]|External URLs that can be redirected to as part of logging in or logging out of the app Note that the query string part of the URL is ignored This is an advanced setting typically only needed by Windows Store application backends Note that URLs within the current domain are always implicitly allowed|
|default_provider|text|The default authentication provider to use when multiple providers are configured This setting is only needed if multiple providers are configured and the unauthenticated client action is set to "RedirectToLoginPage" Possible values include: 'BuiltInAuthenticationProviderAzureActiveDirectory', 'BuiltInAuthenticationProviderFacebook', 'BuiltInAuthenticationProviderGoogle', 'BuiltInAuthenticationProviderMicrosoftAccount', 'BuiltInAuthenticationProviderTwitter', 'BuiltInAuthenticationProviderGithub'|
|token_refresh_extension_hours|float|The number of hours after session token expiration that a session token can be used to call the token refresh API The default is 72 hours|
|client_id|text|The Client ID of this relying party application, known as the client_id This setting is required for enabling OpenID Connection authentication with Azure Active Directory or other 3rd party OpenID Connect providers More information on OpenID Connect: http://openidnet/specs/openid-connect-core-1_0html|
|client_secret|text|The Client Secret of this relying party application (in Azure Active Directory, this is also referred to as the Key) This setting is optional If no client secret is configured, the OpenID Connect implicit auth flow is used to authenticate end users Otherwise, the OpenID Connect Authorization Code Flow is used to authenticate end users More information on OpenID Connect: http://openidnet/specs/openid-connect-core-1_0html|
|client_secret_setting_name|text|The app setting name that contains the client secret of the relying party application|
|client_secret_certificate_thumbprint|text|An alternative to the client secret, that is the thumbprint of a certificate used for signing purposes This property acts as a replacement for the Client Secret It is also optional|
|issuer|text|The OpenID Connect Issuer URI that represents the entity which issues access tokens for this application When using Azure Active Directory, this value is the URI of the directory tenant, eg https://stswindowsnet/{tenant-guid}/ This URI is a case-sensitive identifier for the token issuer More information on OpenID Connect Discovery: http://openidnet/specs/openid-connect-discovery-1_0html|
|validate_issuer|boolean|Gets a value indicating whether the issuer should be a valid HTTPS url and be validated as such|
|allowed_audiences|text[]|Allowed audience values to consider when validating JWTs issued by Azure Active Directory Note that the <code>ClientID</code> value is always considered an allowed audience, regardless of this setting|
|additional_login_params|text[]|Login parameters to send to the OpenID Connect authorization endpoint when a user logs in Each parameter must be in the form "key=value"|
|aad_claims_authorization|text|Gets a JSON string containing the Azure AD Acl settings|
|google_client_id|text|The OpenID Connect Client ID for the Google web application This setting is required for enabling Google Sign-In Google Sign-In documentation: https://developersgooglecom/identity/sign-in/web/|
|google_client_secret|text|The client secret associated with the Google web application This setting is required for enabling Google Sign-In Google Sign-In documentation: https://developersgooglecom/identity/sign-in/web/|
|google_client_secret_setting_name|text|The app setting name that contains the client secret associated with the Google web application|
|google_oauth_scopes|text[]|The OAuth 20 scopes that will be requested as part of Google Sign-In authentication This setting is optional If not specified, "openid", "profile", and "email" are used as default scopes Google Sign-In documentation: https://developersgooglecom/identity/sign-in/web/|
|facebook_app_id|text|The App ID of the Facebook app used for login This setting is required for enabling Facebook Login Facebook Login documentation: https://developersfacebookcom/docs/facebook-login|
|facebook_app_secret|text|The App Secret of the Facebook app used for Facebook Login This setting is required for enabling Facebook Login Facebook Login documentation: https://developersfacebookcom/docs/facebook-login|
|facebook_app_secret_setting_name|text|The app setting name that contains the app secret used for Facebook Login|
|facebook_oauth_scopes|text[]|The OAuth 20 scopes that will be requested as part of Facebook Login authentication This setting is optional Facebook Login documentation: https://developersfacebookcom/docs/facebook-login|
|git_hub_client_id|text|The Client Id of the GitHub app used for login This setting is required for enabling Github login|
|git_hub_client_secret|text|The Client Secret of the GitHub app used for Github Login This setting is required for enabling Github login|
|git_hub_client_secret_setting_name|text|The app setting name that contains the client secret of the Github app used for GitHub Login|
|git_hub_oauth_scopes|text[]|The OAuth 20 scopes that will be requested as part of GitHub Login authentication This setting is optional|
|twitter_consumer_key|text|The OAuth 10a consumer key of the Twitter application used for sign-in This setting is required for enabling Twitter Sign-In Twitter Sign-In documentation: https://devtwittercom/web/sign-in|
|twitter_consumer_secret|text|The OAuth 10a consumer secret of the Twitter application used for sign-in This setting is required for enabling Twitter Sign-In Twitter Sign-In documentation: https://devtwittercom/web/sign-in|
|twitter_consumer_secret_setting_name|text|The app setting name that contains the OAuth 10a consumer secret of the Twitter application used for sign-in|
|microsoft_account_client_id|text|The OAuth 20 client ID that was created for the app used for authentication This setting is required for enabling Microsoft Account authentication Microsoft Account OAuth documentation: https://devonedrivecom/auth/msa_oauthhtm|
|microsoft_account_client_secret|text|The OAuth 20 client secret that was created for the app used for authentication This setting is required for enabling Microsoft Account authentication Microsoft Account OAuth documentation: https://devonedrivecom/auth/msa_oauthhtm|
|microsoft_account_client_secret_setting_name|text|The app setting name containing the OAuth 20 client secret that was created for the app used for authentication|
|microsoft_account_oauth_scopes|text[]|The OAuth 20 scopes that will be requested as part of Microsoft Account authentication This setting is optional If not specified, "wlbasic" is used as the default scope Microsoft Account Scopes and permissions documentation: https://msdnmicrosoftcom/en-us/library/dn631845aspx|
|is_auth_from_file|text|"true" if the auth config settings should be read from a file, "false" otherwise|
|auth_file_path|text|The path of the config file containing auth settings If the path is relative, base will the site's root directory|
|id|text|Resource Id|
|name|text|Resource Name|
|kind|text|Kind of resource|
|type|text|Resource type|
