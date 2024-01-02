# Table: reports

This table shows data for Reports.

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|created_by_user|`json`|
|created_at|`int64`|
|last_updated_by_user|`json`|
|last_updated_at|`int64`|
|kind|`utf8`|
|root_version_id|`utf8`|
|stable_version_id|`utf8`|
|owned_by_user|`json`|
|share|`utf8`|
|target|`utf8`|
|scope|`utf8`|
|level|`utf8`|
|id|`utf8`|
|type|`utf8`|
|name|`utf8`|
|shared|`bool`|
|definition|`json`|
|aggregation|`json`|
|last_run_at|`int64`|
|last_success_run_at|`int64`|
|last_success_run_obj|`utf8`|
|aggregation_list|`json`|