# Table: facebookmarketing_applications

This table shows data for Facebookmarketing Applications.

https://developers.facebook.com/docs/graph-api/reference/application#Reading

The composite primary key for this table is (**account_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|aam_rules|String|
|an_ad_space_limit|Int|
|an_platforms|StringArray|
|android_key_hash|StringArray|
|app_domains|StringArray|
|app_events_feature_bitmask|Int|
|app_events_session_timeout|Int|
|app_install_tracked|Bool|
|app_name|String|
|app_type|Int|
|auth_dialog_data_help_url|String|
|auth_dialog_headline|String|
|auth_dialog_perms_explanation|String|
|auth_referral_default_activity_privacy|String|
|auth_referral_enabled|Int|
|auth_referral_extended_perms|StringArray|
|auth_referral_friend_perms|StringArray|
|auth_referral_response_type|String|
|auth_referral_user_perms|StringArray|
|auto_event_setup_enabled|Bool|
|canvas_fluid_height|Bool|
|canvas_fluid_width|Int|
|canvas_url|String|
|category|String|
|company|String|
|configured_ios_sso|Bool|
|contact_email|String|
|created_time|Timestamp|
|creator_uid|String|
|deauth_callback_url|String|
|default_share_mode|String|
|description|String|
|financial_id|String|
|gdpv4_chrome_custom_tabs_enabled|Bool|
|gdpv4_enabled|Bool|
|gdpv4_nux_content|String|
|gdpv4_nux_enabled|Bool|
|has_messenger_product|Bool|
|hosting_url|String|
|icon_url|String|
|id (PK)|String|
|ios_bundle_id|StringArray|
|ios_sfvc_attr|Bool|
|ios_supports_native_proxy_auth_flow|Bool|
|ios_supports_system_auth|Bool|
|ipad_app_store_id|String|
|iphone_app_store_id|String|
|link|String|
|logging_token|String|
|logo_url|String|
|migrations|JSON|
|mobile_profile_section_url|String|
|mobile_web_url|String|
|name|String|
|namespace|String|
|page_tab_default_name|String|
|page_tab_url|String|
|photo_url|String|
|privacy_policy_url|String|
|profile_section_url|String|
|property_id|String|
|real_time_mode_devices|StringArray|
|restrictive_data_filter_params|String|
|restrictive_data_filter_rules|String|
|sdk_update_message|String|
|seamless_login|Int|
|secure_canvas_url|String|
|secure_page_tab_url|String|
|server_ip_whitelist|String|
|smart_login_bookmark_icon_url|String|
|smart_login_menu_icon_url|String|
|social_discovery|Int|
|subcategory|String|
|suggested_events_setting|String|
|supported_platforms|StringArray|
|supports_attribution|Bool|
|supports_implicit_sdk_logging|Bool|
|suppress_native_ios_gdp|Bool|
|terms_of_service_url|String|
|url_scheme_suffix|String|
|user_support_email|String|
|user_support_url|String|
|website_url|String|