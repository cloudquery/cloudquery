# Table: datadog_dashboards

This table shows data for Datadog Dashboards.

The composite primary key for this table is (**account_name**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_name (PK)|`utf8`|
|author_handle|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|description|`utf8`|
|id (PK)|`utf8`|
|is_read_only|`bool`|
|layout_type|`utf8`|
|modified_at|`timestamp[us, tz=UTC]`|
|title|`utf8`|
|url|`utf8`|
|additional_properties|`json`|