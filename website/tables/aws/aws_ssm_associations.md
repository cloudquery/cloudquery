# Table: aws_ssm_associations

This table shows data for AWS Systems Manager (SSM) Associations.

https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_Association.html

The composite primary key for this table is (**account_id**, **region**, **association_id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|association_id (PK)|`utf8`|
|association_name|`utf8`|
|association_version|`utf8`|
|document_version|`utf8`|
|instance_id|`utf8`|
|last_execution_date|`timestamp[us, tz=UTC]`|
|name|`utf8`|
|overview|`json`|
|schedule_expression|`utf8`|
|schedule_offset|`int64`|
|target_maps|`json`|
|targets|`json`|