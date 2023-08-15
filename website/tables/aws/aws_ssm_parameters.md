# Table: aws_ssm_parameters

This table shows data for AWS Systems Manager (SSM) Parameters.

https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_ParameterMetadata.html

The composite primary key for this table is (**account_id**, **region**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|name (PK)|`utf8`|
|allowed_pattern|`utf8`|
|data_type|`utf8`|
|description|`utf8`|
|key_id|`utf8`|
|last_modified_date|`timestamp[us, tz=UTC]`|
|last_modified_user|`utf8`|
|policies|`json`|
|tier|`utf8`|
|type|`utf8`|
|version|`int64`|