# Table: aws_identitystore_group_memberships

This table shows data for Identity Store Group Memberships.

https://docs.aws.amazon.com/singlesignon/latest/IdentityStoreAPIReference/API_GroupMembership.html

The composite primary key for this table is (**request_account_id**, **request_region**, **group_arn**, **arn**).

## Relations

This table depends on [aws_identitystore_groups](aws_identitystore_groups.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|request_account_id (PK)|`utf8`|
|request_region (PK)|`utf8`|
|group_arn (PK)|`utf8`|
|arn (PK)|`utf8`|
|member_id|`utf8`|
|identity_store_id|`utf8`|
|group_id|`utf8`|
|membership_id|`utf8`|