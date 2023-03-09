# Table: facebookmarketing_publisher_block_lists

https://developers.facebook.com/docs/marketing-api/reference/publisher-block-list/

The composite primary key for this table is (**account_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|business_owner_id|String|
|id (PK)|String|
|is_auto_blocking_on|Bool|
|is_eligible_at_campaign_level|Bool|
|last_update_time|Timestamp|
|last_update_user|String|
|name|String|
|owner_ad_account_id|String|