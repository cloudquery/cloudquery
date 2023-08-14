# Table: aws_rds_cluster_parameters

This table shows data for Amazon Relational Database Service (RDS) Cluster Parameters.

https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_Parameter.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_rds_engine_versions](aws_rds_engine_versions).

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
|supported_engine_modes|`list<item: utf8, nullable>`|