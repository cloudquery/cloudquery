# Source Plugin: googleanalytics

## Tables

All tables propagated by `googleanalytics` plugin are dynamic.
The following rules are adhered to while creating table scheme and propagating data.

### Table name

Every table name is constructed as `ga_` prefix followed by a report name in snake case.
An error will be reported if you have several reports that have different names
but the snake case transformation produces the same string.

### Schema

The composite primary key for this table is (**property_id**, **date**, **dimension_hash**).
It supports incremental syncs based on the **date** column.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|property_id (PK)|String|
|date (PK) (Incremental Key)|Timestamp|
|dimensions|JSON|
|dimension_hash (PK)|ByteArray|
|metrics|JSON|
|data_loss|Bool|