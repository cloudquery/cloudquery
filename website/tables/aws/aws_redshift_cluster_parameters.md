# Table: aws_redshift_cluster_parameters

This table shows data for Redshift Cluster Parameters.

https://docs.aws.amazon.com/redshift/latest/APIReference/API_Parameter.html

The composite primary key for this table is (**cluster_arn**, **parameter_name**).

## Relations

This table depends on [aws_redshift_cluster_parameter_groups](aws_redshift_cluster_parameter_groups).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|account_id|utf8|
|region|utf8|
|cluster_arn (PK)|utf8|
|parameter_name (PK)|utf8|
|allowed_values|utf8|
|apply_type|utf8|
|data_type|utf8|
|description|utf8|
|is_modifiable|bool|
|minimum_engine_version|utf8|
|parameter_value|utf8|
|source|utf8|