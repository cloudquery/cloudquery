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

#### Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|property_id (PK)|utf8|
|date (PK) (Incremental Key)|timestamp[us, tz=UTC]|
|dimensions|json|
|dimension_hash (PK)|binary|
|metrics|json|
|data_loss|bool|

### `dat_loss` column

Data loss indicates if the `(other)` value is present due to hitting Google Analytics limits.
To learn more, read [this article](https://support.google.com/analytics/answer/13331684).