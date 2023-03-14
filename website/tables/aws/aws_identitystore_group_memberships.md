# Table: aws_identitystore_group_memberships

This table shows data for Identity Store Group Memberships.

https://docs.aws.amazon.com/singlesignon/latest/IdentityStoreAPIReference/API_GroupMembership.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_identitystore_groups](aws_identitystore_groups).

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