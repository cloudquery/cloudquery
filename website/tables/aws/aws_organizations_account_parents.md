# Table: aws_organizations_account_parents

This table shows data for Organizations Account Parents.

https://docs.aws.amazon.com/organizations/latest/APIReference/API_ListParents.html
The 'request_account_id' column is added to show from where the request was made.

The composite primary key for this table is (**request_account_id**, **id**, **parent_id**, **type**).

## Relations

This table depends on [aws_organizations_accounts](aws_organizations_accounts).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|request_account_id (PK)|`utf8`|
|id (PK)|`utf8`|
|parent_id (PK)|`utf8`|
|type (PK)|`utf8`|