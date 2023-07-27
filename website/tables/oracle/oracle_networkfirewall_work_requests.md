# Table: oracle_networkfirewall_work_requests

This table shows data for Oracle Network Firewall Work Requests.

The composite primary key for this table is (**region**, **compartment_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|region (PK)|`utf8`|
|compartment_id (PK)|`utf8`|
|operation_type|`utf8`|
|status|`utf8`|
|id (PK)|`utf8`|
|resources|`json`|
|percent_complete|`float64`|
|time_accepted|`timestamp[us, tz=UTC]`|
|time_started|`timestamp[us, tz=UTC]`|
|time_finished|`timestamp[us, tz=UTC]`|