# Table: aws_ssm_associations

https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_Association.html

The composite primary key for this table is (**account_id**, **region**, **association_id**).



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|association_id (PK)|String|
|association_name|String|
|association_version|String|
|document_version|String|
|instance_id|String|
|last_execution_date|Timestamp|
|name|String|
|overview|JSON|
|schedule_expression|String|
|schedule_offset|Int|
|target_maps|JSON|
|targets|JSON|