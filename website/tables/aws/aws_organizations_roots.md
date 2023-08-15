# Table: aws_organizations_roots

This table shows data for Organizations Roots.

https://docs.aws.amazon.com/organizations/latest/APIReference/API_Root.html
The 'request_account_id' column is added to show from where the request was made.

The composite primary key for this table is (**request_account_id**, **arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|request_account_id (PK)|`utf8`|
|tags|`json`|
|arn (PK)|`utf8`|
|id|`utf8`|
|name|`utf8`|
|policy_types|`json`|