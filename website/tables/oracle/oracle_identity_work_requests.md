# Table: oracle_identity_work_requests

This table shows data for Oracle Identity Work Requests.

The composite primary key for this table is (**region**, **compartment_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|region (PK)|`utf8`|
|compartment_id (PK)|`utf8`|
|id (PK)|`utf8`|
|operation_type|`utf8`|
|status|`utf8`|
|resources|`json`|
|errors|`json`|
|time_accepted|`timestamp[us, tz=UTC]`|
|time_started|`timestamp[us, tz=UTC]`|
|time_finished|`timestamp[us, tz=UTC]`|
|percent_complete|`float64`|