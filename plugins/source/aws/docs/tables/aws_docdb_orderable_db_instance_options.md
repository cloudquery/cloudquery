# Table: aws_docdb_orderable_db_instance_options

https://docs.aws.amazon.com/documentdb/latest/developerguide/API_OrderableDBInstanceOption.html

The primary key for this table is **_cq_id**.

## Relations
This table depends on [aws_docdb_engine_versions](aws_docdb_engine_versions.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|availability_zones|JSON|
|db_instance_class|String|
|engine|String|
|engine_version|String|
|license_model|String|
|vpc|Bool|