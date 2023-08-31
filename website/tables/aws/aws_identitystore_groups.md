# Table: aws_identitystore_groups

This table shows data for Identity Store Groups.

https://docs.aws.amazon.com/singlesignon/latest/IdentityStoreAPIReference/API_Group.html

The primary key for this table is **_cq_id**.

## Relations

The following tables depend on aws_identitystore_groups:
  - [aws_identitystore_group_memberships](aws_identitystore_group_memberships)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|group_id|`utf8`|
|identity_store_id|`utf8`|
|description|`utf8`|
|display_name|`utf8`|
|external_ids|`json`|