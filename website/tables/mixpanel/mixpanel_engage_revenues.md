# Table: mixpanel_engage_revenues

This table shows data for Mixpanel Engage Revenues.

The composite primary key for this table is (**project_id**, **date**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`int64`|
|date (PK)|`timestamp[us, tz=UTC]`|
|amount|`float64`|
|count|`int64`|
|paid_count|`int64`|