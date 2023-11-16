# Table: datadog_hosts

This table shows data for Datadog Hosts.

The composite primary key for this table is (**account_name**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_name (PK)|`utf8`|
|aliases|`list<item: utf8, nullable>`|
|apps|`list<item: utf8, nullable>`|
|aws_name|`utf8`|
|host_name|`utf8`|
|id (PK)|`int64`|
|is_muted|`bool`|
|last_reported_time|`int64`|
|meta|`json`|
|metrics|`json`|
|mute_timeout|`int64`|
|name|`utf8`|
|sources|`list<item: utf8, nullable>`|
|tags_by_source|`json`|
|up|`bool`|
|additional_properties|`json`|