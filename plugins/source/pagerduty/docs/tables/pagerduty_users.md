# Table: pagerduty_users

https://developer.pagerduty.com/api-reference/c96e889522dd6-list-users

The primary key for this table is **id**.

## Relations

The following tables depend on pagerduty_users:
  - [pagerduty_user_contact_methods](pagerduty_user_contact_methods.md)
  - [pagerduty_user_notification_rules](pagerduty_user_notification_rules.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|html_url|String|
|avatar_url|String|
|type|String|
|summary|String|
|self|String|
|name|String|
|email|String|
|time_zone|String|
|color|String|
|role|String|
|description|String|
|invitation_sent|Bool|
|contact_methods|JSON|
|notification_rules|JSON|
|job_title|String|
|teams|JSON|