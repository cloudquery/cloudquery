# Table: facebookmarketing_customconversions

This table shows data for Facebook Marketing Custom Conversions.

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|aggregation_rule|`utf8`|
|business|`json`|
|creation_time|`timestamp[us, tz=UTC]`|
|custom_event_type|`utf8`|
|data_sources|`json`|
|default_conversion_value|`int64`|
|description|`utf8`|
|event_source_type|`utf8`|
|first_fired_time|`timestamp[us, tz=UTC]`|
|id (PK)|`utf8`|
|is_archived|`bool`|
|is_unavailable|`bool`|
|last_fired_time|`timestamp[us, tz=UTC]`|
|name|`utf8`|
|offline_conversion_data_set|`json`|
|pixel|`json`|
|retention_days|`int64`|
|rule|`utf8`|