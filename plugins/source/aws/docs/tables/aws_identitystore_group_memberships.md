# Table: aws_identitystore_group_memberships



The primary key for this table is **_cq_id**.

## Relations
This table depends on [aws_identitystore_groups](aws_identitystore_groups.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|identity_store_id|String|
|group_id|String|
|membership_id|String|