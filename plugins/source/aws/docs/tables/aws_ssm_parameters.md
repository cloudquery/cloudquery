# Table: aws_ssm_parameters

https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_ParameterMetadata.html

The composite primary key for this table is (**account_id**, **region**, **name**).



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|name (PK)|String|
|allowed_pattern|String|
|data_type|String|
|description|String|
|key_id|String|
|last_modified_date|Timestamp|
|last_modified_user|String|
|policies|JSON|
|tier|String|
|type|String|
|version|Int|