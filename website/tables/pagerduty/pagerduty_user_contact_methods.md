# Table: pagerduty_user_contact_methods

This table shows data for PagerDuty User Contact Methods.

https://developer.pagerduty.com/api-reference/50d46c0eb020d-list-a-user-s-contact-methods

The primary key for this table is **id**.

## Relations

This table depends on [pagerduty_users](pagerduty_users).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|html_url|`utf8`|
|send_html_email|`bool`|
|type|`utf8`|
|summary|`utf8`|
|self|`utf8`|
|label|`utf8`|
|address|`utf8`|
|send_short_email|`bool`|
|blacklisted|`bool`|
|country_code|`int64`|
|enabled|`bool`|