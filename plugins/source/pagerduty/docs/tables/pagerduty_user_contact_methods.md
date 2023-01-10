# Table: pagerduty_user_contact_methods

https://developer.pagerduty.com/api-reference/50d46c0eb020d-list-a-user-s-contact-methods

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
|html_url|String|
|send_html_email|Bool|
|type|String|
|summary|String|
|self|String|
|label|String|
|address|String|
|send_short_email|Bool|
|blacklisted|Bool|
|country_code|Int|
|enabled|Bool|