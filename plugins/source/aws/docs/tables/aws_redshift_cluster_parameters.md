# Table: aws_redshift_cluster_parameters

https://docs.aws.amazon.com/redshift/latest/APIReference/API_Parameter.html

The composite primary key for this table is (**cluster_arn**, **parameter_name**).

## Relations
This table depends on [aws_redshift_cluster_parameter_groups](aws_redshift_cluster_parameter_groups.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|cluster_arn (PK)|String|
|parameter_name (PK)|String|
|allowed_values|String|
|apply_type|String|
|data_type|String|
|description|String|
|is_modifiable|Bool|
|minimum_engine_version|String|
|parameter_value|String|
|source|String|