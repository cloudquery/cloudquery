# Table: aws_docdb_cluster_parameters

This table shows data for Amazon DocumentDB Cluster Parameters.

https://docs.aws.amazon.com/documentdb/latest/developerguide/API_Parameter.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_docdb_engine_versions](aws_docdb_engine_versions).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|allowed_values|`utf8`|
|apply_method|`utf8`|
|apply_type|`utf8`|
|data_type|`utf8`|
|description|`utf8`|
|is_modifiable|`bool`|
|minimum_engine_version|`utf8`|
|parameter_name|`utf8`|
|parameter_value|`utf8`|
|source|`utf8`|