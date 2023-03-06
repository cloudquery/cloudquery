# Table: aws_identitystore_groups

https://docs.aws.amazon.com/singlesignon/latest/IdentityStoreAPIReference/API_Group.html

The primary key for this table is **_cq_id**.

## Relations

The following tables depend on aws_identitystore_groups:
  - [aws_identitystore_group_memberships](aws_identitystore_group_memberships)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|group_id|String|
|identity_store_id|String|
|description|String|
|display_name|String|
|external_ids|JSON|