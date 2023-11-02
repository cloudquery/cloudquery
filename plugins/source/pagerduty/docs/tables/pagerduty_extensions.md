# Table: pagerduty_extensions

This table shows data for PagerDuty Extensions.

https://developer.pagerduty.com/api-reference/26b46f0092a55-list-extensions

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|html_url|`utf8`|
|endpoint_url|`utf8`|
|type|`utf8`|
|summary|`utf8`|
|self|`utf8`|
|name|`utf8`|
|extension_objects|`json`|
|extension_schema|`json`|
|temporarily_disabled|`bool`|