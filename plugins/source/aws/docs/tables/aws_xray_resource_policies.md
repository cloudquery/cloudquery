# Table: aws_xray_resource_policies

https://docs.aws.amazon.com/xray/latest/api/API_ResourcePolicy.html

The composite primary key for this table is (**account_id**, **region**, **policy_name**, **policy_revision_id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|policy_name (PK)|String|
|policy_revision_id (PK)|String|
|last_updated_time|Timestamp|
|policy_document|String|