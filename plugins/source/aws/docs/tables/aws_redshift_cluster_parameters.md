# Table: aws_redshift_cluster_parameters

This table shows data for Redshift Cluster Parameters.

https://docs.aws.amazon.com/redshift/latest/APIReference/API_Parameter.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**cluster_arn**, **parameter_name**).
## Relations

This table depends on [aws_redshift_cluster_parameter_groups](aws_redshift_cluster_parameter_groups.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|cluster_arn|`utf8`|
|allowed_values|`utf8`|
|apply_type|`utf8`|
|data_type|`utf8`|
|description|`utf8`|
|is_modifiable|`bool`|
|minimum_engine_version|`utf8`|
|parameter_name|`utf8`|
|parameter_value|`utf8`|
|source|`utf8`|