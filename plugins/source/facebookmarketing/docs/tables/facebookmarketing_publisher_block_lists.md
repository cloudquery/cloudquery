# Table: facebookmarketing_publisher_block_lists

This table shows data for Facebook Marketing Publisher Block Lists.

https://developers.facebook.com/docs/marketing-api/reference/publisher-block-list/

The composite primary key for this table is (**account_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|business_owner_id|`utf8`|
|id (PK)|`utf8`|
|is_auto_blocking_on|`bool`|
|is_eligible_at_campaign_level|`bool`|
|last_update_time|`timestamp[us, tz=UTC]`|
|last_update_user|`utf8`|
|name|`utf8`|
|owner_ad_account_id|`utf8`|