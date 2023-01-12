# Table: oracle_objectstorage_work_requests

The composite primary key for this table is (**region**, **compartment_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|region (PK)|String|
|compartment_id (PK)|String|
|id (PK)|String|
|operation_type|String|
|status|String|
|resources|JSON|
|percent_complete|Float|
|time_accepted|Timestamp|
|time_started|Timestamp|
|time_finished|Timestamp|