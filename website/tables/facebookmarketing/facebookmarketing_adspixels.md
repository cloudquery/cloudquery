# Table: facebookmarketing_adspixels

https://developers.facebook.com/docs/graph-api/reference/ads-pixel/#Reading

The composite primary key for this table is (**account_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|automatic_matching_fields|StringArray|
|can_proxy|Bool|
|code|String|
|config|String|
|creation_time|Timestamp|
|creator|JSON|
|data_use_setting|String|
|description|String|
|duplicate_entries|Int|
|enable_auto_assign_to_accounts|Bool|
|enable_automatic_matching|Bool|
|event_stats|String|
|event_time_max|Int|
|event_time_min|Int|
|first_party_cookie_status|String|
|id (PK)|String|
|is_consolidated_container|Bool|
|is_created_by_business|Bool|
|is_crm|Bool|
|is_mta_use|Bool|
|is_restricted_use|Bool|
|is_unavailable|Bool|
|last_fired_time|Timestamp|
|last_upload_app|String|
|last_upload_app_changed_time|Int|
|match_rate_approx|Int|
|matched_entries|Int|
|name|String|
|owner_ad_account|JSON|
|owner_business|JSON|
|usage|JSON|
|valid_entries|Int|