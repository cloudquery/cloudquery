# Table: aws_xray_resource_policies

https://docs.aws.amazon.com/xray/latest/api/API_ResourcePolicy.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|last_updated_time|Timestamp|
|policy_document|String|
|policy_name|String|
|policy_revision_id|String|