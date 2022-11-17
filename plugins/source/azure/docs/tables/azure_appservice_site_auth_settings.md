# Table: azure_appservice_site_auth_settings

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v2#SiteAuthSettings

The primary key for this table is **id**.

## Relations
This table depends on [azure_appservice_sites](azure_appservice_sites.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|kind|String|
|aad_claims_authorization|String|
|additional_login_params|StringArray|
|allowed_audiences|StringArray|
|allowed_external_redirect_urls|StringArray|
|auth_file_path|String|
|client_id|String|
|client_secret|String|
|client_secret_certificate_thumbprint|String|
|client_secret_setting_name|String|
|config_version|String|
|default_provider|String|
|enabled|Bool|
|facebook_app_id|String|
|facebook_app_secret|String|
|facebook_app_secret_setting_name|String|
|facebook_o_auth_scopes|StringArray|
|git_hub_client_id|String|
|git_hub_client_secret|String|
|git_hub_client_secret_setting_name|String|
|git_hub_o_auth_scopes|StringArray|
|google_client_id|String|
|google_client_secret|String|
|google_client_secret_setting_name|String|
|google_o_auth_scopes|StringArray|
|is_auth_from_file|String|
|issuer|String|
|microsoft_account_client_id|String|
|microsoft_account_client_secret|String|
|microsoft_account_client_secret_setting_name|String|
|microsoft_account_o_auth_scopes|StringArray|
|runtime_version|String|
|token_refresh_extension_hours|Float|
|token_store_enabled|Bool|
|twitter_consumer_key|String|
|twitter_consumer_secret|String|
|twitter_consumer_secret_setting_name|String|
|unauthenticated_client_action|String|
|validate_issuer|Bool|
|id (PK)|String|
|name|String|
|type|String|
|site_id|String|