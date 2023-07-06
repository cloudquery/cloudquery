# Table: facebookmarketing_saved_audiences

This table shows data for Facebook Marketing Saved Audiences.

https://developers.facebook.com/docs/marketing-api/reference/saved-audience/#Reading

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`json`|
|approximate_count_lower_bound|`int64`|
|approximate_count_upper_bound|`int64`|
|delete_time|`int64`|
|description|`utf8`|
|extra_info|`utf8`|
|id (PK)|`utf8`|
|name|`utf8`|
|operation_status|`json`|
|owner_business|`json`|
|page_deletion_marked_delete_time|`int64`|
|permission_for_actions|`json`|
|run_status|`utf8`|
|targeting|`json`|
|time_created|`timestamp[us, tz=UTC]`|
|time_updated|`timestamp[us, tz=UTC]`|