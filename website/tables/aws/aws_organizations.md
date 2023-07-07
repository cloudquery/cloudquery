# Table: aws_organizations

This table shows data for Organizations.

https://docs.aws.amazon.com/organizations/latest/APIReference/API_Organization.html

The composite primary key for this table is (**account_id**, **arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|arn (PK)|`utf8`|
|feature_set|`utf8`|
|id|`utf8`|
|master_account_arn|`utf8`|
|master_account_email|`utf8`|
|master_account_id|`utf8`|