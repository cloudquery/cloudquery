# Table: aws_organizations_organizational_units

This table shows data for Organizations Organizational Units.

https://docs.aws.amazon.com/organizations/latest/APIReference/API_OrganizationalUnit.html
The 'request_account_id' column is added to show from where the request was made.

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**request_account_id**, **arn**).
## Relations

The following tables depend on aws_organizations_organizational_units:
  - [aws_organizations_organizational_unit_parents](aws_organizations_organizational_unit_parents.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|request_account_id|`utf8`|
|arn|`utf8`|
|id|`utf8`|
|name|`utf8`|