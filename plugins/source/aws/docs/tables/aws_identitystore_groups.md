# Table: aws_identitystore_groups

This table shows data for Identity Store Groups.

https://docs.aws.amazon.com/singlesignon/latest/IdentityStoreAPIReference/API_Group.html

The composite primary key for this table is (**request_account_id**, **request_region**, **arn**).

## Relations

The following tables depend on aws_identitystore_groups:
  - [aws_identitystore_group_memberships](aws_identitystore_group_memberships.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|request_account_id (PK)|`utf8`|
|request_region (PK)|`utf8`|
|arn (PK)|`utf8`|
|group_id|`utf8`|
|identity_store_id|`utf8`|
|description|`utf8`|
|display_name|`utf8`|
|external_ids|`json`|