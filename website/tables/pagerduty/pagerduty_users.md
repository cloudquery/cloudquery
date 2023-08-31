# Table: pagerduty_users

This table shows data for PagerDuty Users.

https://developer.pagerduty.com/api-reference/c96e889522dd6-list-users

The primary key for this table is **id**.

## Relations

The following tables depend on pagerduty_users:
  - [pagerduty_user_contact_methods](pagerduty_user_contact_methods)
  - [pagerduty_user_notification_rules](pagerduty_user_notification_rules)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|html_url|`utf8`|
|avatar_url|`utf8`|
|type|`utf8`|
|summary|`utf8`|
|self|`utf8`|
|name|`utf8`|
|email|`utf8`|
|time_zone|`utf8`|
|color|`utf8`|
|role|`utf8`|
|description|`utf8`|
|invitation_sent|`bool`|
|contact_methods|`json`|
|notification_rules|`json`|
|job_title|`utf8`|
|teams|`json`|