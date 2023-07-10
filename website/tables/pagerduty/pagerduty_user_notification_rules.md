# Table: pagerduty_user_notification_rules

This table shows data for PagerDuty User Notification Rules.

https://developer.pagerduty.com/api-reference/043092de7e3e1-list-a-user-s-notification-rules

The primary key for this table is **id**.

## Relations

This table depends on [pagerduty_users](pagerduty_users).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|type|`utf8`|
|summary|`utf8`|
|self|`utf8`|
|html_url|`utf8`|
|start_delay_in_minutes|`int64`|
|created_at|`timestamp[us, tz=UTC]`|
|contact_method|`json`|
|urgency|`utf8`|