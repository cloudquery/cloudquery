# Table: cloudflare_access_applications

This table shows data for Cloudflare Access Applications.

The primary key for this table is **id**.

## Relations

The following tables depend on cloudflare_access_applications:
  - [cloudflare_access_applications_self_hosted_domains](cloudflare_access_applications_self_hosted_domains.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|zone_id|`utf8`|
|id (PK)|`utf8`|
|gateway_rules|`json`|
|allowed_idps|`list<item: utf8, nullable>`|
|custom_deny_message|`utf8`|
|logo_url|`utf8`|
|aud|`utf8`|
|domain|`utf8`|
|self_hosted_domains|`list<item: utf8, nullable>`|
|type|`utf8`|
|session_duration|`utf8`|
|same_site_cookie_attribute|`utf8`|
|custom_deny_url|`utf8`|
|custom_non_identity_deny_url|`utf8`|
|name|`utf8`|
|private_address|`utf8`|
|cors_headers|`json`|
|created_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|
|saas_app|`json`|
|auto_redirect_to_identity|`bool`|
|skip_interstitial|`bool`|
|app_launcher_visible|`bool`|
|enable_binding_cookie|`bool`|
|http_only_cookie_attribute|`bool`|
|service_auth_401_redirect|`bool`|
|path_cookie_attribute|`bool`|
|custom_pages|`list<item: utf8, nullable>`|
|tags|`list<item: utf8, nullable>`|
|access_app_launcher_customization|`json`|