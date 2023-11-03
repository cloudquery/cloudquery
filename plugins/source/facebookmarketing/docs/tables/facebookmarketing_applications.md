# Table: facebookmarketing_applications

This table shows data for Facebook Marketing Applications.

https://developers.facebook.com/docs/graph-api/reference/application#Reading

The composite primary key for this table is (**account_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|aam_rules|`utf8`|
|an_ad_space_limit|`int64`|
|an_platforms|`list<item: utf8, nullable>`|
|android_key_hash|`list<item: utf8, nullable>`|
|app_domains|`list<item: utf8, nullable>`|
|app_events_feature_bitmask|`int64`|
|app_events_session_timeout|`int64`|
|app_install_tracked|`bool`|
|app_name|`utf8`|
|app_type|`int64`|
|auth_dialog_data_help_url|`utf8`|
|auth_dialog_headline|`utf8`|
|auth_dialog_perms_explanation|`utf8`|
|auth_referral_default_activity_privacy|`utf8`|
|auth_referral_enabled|`int64`|
|auth_referral_extended_perms|`list<item: utf8, nullable>`|
|auth_referral_friend_perms|`list<item: utf8, nullable>`|
|auth_referral_response_type|`utf8`|
|auth_referral_user_perms|`list<item: utf8, nullable>`|
|auto_event_setup_enabled|`bool`|
|canvas_fluid_height|`bool`|
|canvas_fluid_width|`int64`|
|canvas_url|`utf8`|
|category|`utf8`|
|company|`utf8`|
|configured_ios_sso|`bool`|
|contact_email|`utf8`|
|created_time|`timestamp[us, tz=UTC]`|
|creator_uid|`utf8`|
|deauth_callback_url|`utf8`|
|default_share_mode|`utf8`|
|description|`utf8`|
|financial_id|`utf8`|
|gdpv4_chrome_custom_tabs_enabled|`bool`|
|gdpv4_enabled|`bool`|
|gdpv4_nux_content|`utf8`|
|gdpv4_nux_enabled|`bool`|
|has_messenger_product|`bool`|
|hosting_url|`utf8`|
|icon_url|`utf8`|
|id (PK)|`utf8`|
|ios_bundle_id|`list<item: utf8, nullable>`|
|ios_sfvc_attr|`bool`|
|ios_supports_native_proxy_auth_flow|`bool`|
|ios_supports_system_auth|`bool`|
|ipad_app_store_id|`utf8`|
|iphone_app_store_id|`utf8`|
|link|`utf8`|
|logging_token|`utf8`|
|logo_url|`utf8`|
|migrations|`json`|
|mobile_profile_section_url|`utf8`|
|mobile_web_url|`utf8`|
|name|`utf8`|
|namespace|`utf8`|
|page_tab_default_name|`utf8`|
|page_tab_url|`utf8`|
|photo_url|`utf8`|
|privacy_policy_url|`utf8`|
|profile_section_url|`utf8`|
|property_id|`utf8`|
|real_time_mode_devices|`list<item: utf8, nullable>`|
|restrictive_data_filter_params|`utf8`|
|restrictive_data_filter_rules|`utf8`|
|sdk_update_message|`utf8`|
|seamless_login|`int64`|
|secure_canvas_url|`utf8`|
|secure_page_tab_url|`utf8`|
|server_ip_whitelist|`utf8`|
|smart_login_bookmark_icon_url|`utf8`|
|smart_login_menu_icon_url|`utf8`|
|social_discovery|`int64`|
|subcategory|`utf8`|
|suggested_events_setting|`utf8`|
|supported_platforms|`list<item: utf8, nullable>`|
|supports_attribution|`bool`|
|supports_implicit_sdk_logging|`bool`|
|suppress_native_ios_gdp|`bool`|
|terms_of_service_url|`utf8`|
|url_scheme_suffix|`utf8`|
|user_support_email|`utf8`|
|user_support_url|`utf8`|
|website_url|`utf8`|