# Table: github_hook_deliveries

This table shows data for Github Hook Deliveries.

The composite primary key for this table is (**org**, **hook_id**, **id**).

## Relations

This table depends on [github_hooks](github_hooks).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|org (PK)|`utf8`|
|hook_id (PK)|`int64`|
|id (PK)|`int64`|
|guid|`utf8`|
|delivered_at|`timestamp[us, tz=UTC]`|
|redelivery|`bool`|
|duration|`float64`|
|status|`utf8`|
|status_code|`int64`|
|event|`utf8`|
|action|`utf8`|
|installation_id|`int64`|
|repository_id|`int64`|
|request|`json`|
|response|`json`|