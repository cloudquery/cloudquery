# Table: aws_ssm_parameters


The composite primary key for this table is (**account_id**, **region**, **name**).


## Columns
| Name          | Type          |
| ------------- | ------------- |
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
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|