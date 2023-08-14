# Table: aws_docdb_orderable_db_instance_options

This table shows data for Amazon DocumentDB Orderable DB Instance Options.

https://docs.aws.amazon.com/documentdb/latest/developerguide/API_OrderableDBInstanceOption.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_docdb_engine_versions](aws_docdb_engine_versions).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|availability_zones|`json`|
|db_instance_class|`utf8`|
|engine|`utf8`|
|engine_version|`utf8`|
|license_model|`utf8`|
|vpc|`bool`|