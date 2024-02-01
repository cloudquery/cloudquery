# Table: aws_organizations_account_parents

This table shows data for Organizations Account Parents.

https://docs.aws.amazon.com/organizations/latest/APIReference/API_ListParents.html
The 'request_account_id' column is added to show from where the request was made.

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**request_account_id**, **id**, **parent_id**, **type**).
## Relations

This table depends on [aws_organizations_accounts](aws_organizations_accounts.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|request_account_id|`utf8`|
|id|`utf8`|
|parent_id|`utf8`|
|type|`utf8`|