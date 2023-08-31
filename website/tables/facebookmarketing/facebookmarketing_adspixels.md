# Table: facebookmarketing_adspixels

This table shows data for Facebook Marketing Ads Pixels.

https://developers.facebook.com/docs/graph-api/reference/ads-pixel/#Reading

The composite primary key for this table is (**account_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|automatic_matching_fields|`list<item: utf8, nullable>`|
|can_proxy|`bool`|
|code|`utf8`|
|config|`utf8`|
|creation_time|`timestamp[us, tz=UTC]`|
|creator|`json`|
|data_use_setting|`utf8`|
|description|`utf8`|
|duplicate_entries|`int64`|
|enable_auto_assign_to_accounts|`bool`|
|enable_automatic_matching|`bool`|
|event_stats|`utf8`|
|event_time_max|`int64`|
|event_time_min|`int64`|
|first_party_cookie_status|`utf8`|
|id (PK)|`utf8`|
|is_consolidated_container|`bool`|
|is_created_by_business|`bool`|
|is_crm|`bool`|
|is_mta_use|`bool`|
|is_restricted_use|`bool`|
|is_unavailable|`bool`|
|last_fired_time|`timestamp[us, tz=UTC]`|
|last_upload_app|`utf8`|
|last_upload_app_changed_time|`int64`|
|match_rate_approx|`int64`|
|matched_entries|`int64`|
|name|`utf8`|
|owner_ad_account|`json`|
|owner_business|`json`|
|usage|`json`|
|valid_entries|`int64`|