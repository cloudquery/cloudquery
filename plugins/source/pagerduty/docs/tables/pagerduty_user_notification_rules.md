# Table: pagerduty_user_notification_rules

https://developer.pagerduty.com/api-reference/043092de7e3e1-list-a-user-s-notification-rules

The primary key for this table is **id**.

## Relations

This table depends on [pagerduty_users](pagerduty_users.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|type|String|
|summary|String|
|self|String|
|html_url|String|
|start_delay_in_minutes|Int|
|created_at|Timestamp|
|contact_method|JSON|
|urgency|String|