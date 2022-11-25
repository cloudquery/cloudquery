# Table: azure_web_site_auth_settings

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/web/mgmt/2020-12-01/web#SiteAuthSettings

The primary key for this table is **id**.

## Relations
This table depends on [azure_web_apps](azure_web_apps.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|web_app_id|String|
|enabled|Bool|
|runtime_version|String|
|unauthenticated_client_action|String|
|token_store_enabled|Bool|
|allowed_external_redirect_urls|StringArray|
|default_provider|String|
|token_refresh_extension_hours|Float|
|client_id|String|
|client_secret|String|
|client_secret_setting_name|String|
|client_secret_certificate_thumbprint|String|
|issuer|String|
|validate_issuer|Bool|
|allowed_audiences|StringArray|
|additional_login_params|StringArray|
|aad_claims_authorization|String|
|google_client_id|String|
|google_client_secret|String|
|google_client_secret_setting_name|String|
|google_o_auth_scopes|StringArray|
|facebook_app_id|String|
|facebook_app_secret|String|
|facebook_app_secret_setting_name|String|
|facebook_o_auth_scopes|StringArray|
|git_hub_client_id|String|
|git_hub_client_secret|String|
|git_hub_client_secret_setting_name|String|
|git_hub_o_auth_scopes|StringArray|
|twitter_consumer_key|String|
|twitter_consumer_secret|String|
|twitter_consumer_secret_setting_name|String|
|microsoft_account_client_id|String|
|microsoft_account_client_secret|String|
|microsoft_account_client_secret_setting_name|String|
|microsoft_account_o_auth_scopes|StringArray|
|is_auth_from_file|String|
|auth_file_path|String|
|config_version|String|
|id (PK)|String|
|name|String|
|kind|String|
|type|String|