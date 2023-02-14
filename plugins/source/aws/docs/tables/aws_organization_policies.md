# Table: aws_organization_policies

https://docs.aws.amazon.com/organizations/latest/APIReference/API_Policy.html

The primary key for this table is **account_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|content|String|
|policy_summary|JSON|