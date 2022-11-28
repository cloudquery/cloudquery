# Table: aws_kms_aliases

https://docs.aws.amazon.com/kms/latest/APIReference/API_AliasListEntry.html

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
|alias_name|String|
|creation_date|Timestamp|
|last_updated_date|Timestamp|
|target_key_id|String|