# Table: aws_organizations_organizational_units

This table shows data for Organizations Organizational Units.

https://docs.aws.amazon.com/organizations/latest/APIReference/API_OrganizationalUnit.html
The 'request_account_id' column is added to show from where the request was made.

The composite primary key for this table is (**request_account_id**, **arn**).

## Relations

The following tables depend on aws_organizations_organizational_units:
  - [aws_organizations_organizational_unit_parents](aws_organizations_organizational_unit_parents)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|request_account_id (PK)|`utf8`|
|arn (PK)|`utf8`|
|id|`utf8`|
|name|`utf8`|