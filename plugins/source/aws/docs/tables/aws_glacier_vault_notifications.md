# Table: aws_glacier_vault_notifications



The primary key for this table is **vault_arn**.

## Relations
This table depends on [aws_glacier_vaults](aws_glacier_vaults.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|vault_arn (PK)|String|
|events|StringArray|
|sns_topic|String|