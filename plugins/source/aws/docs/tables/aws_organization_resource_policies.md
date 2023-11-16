# Table: aws_organization_resource_policies

This table shows data for Organization Resource Policies.

https://docs.aws.amazon.com/organizations/latest/APIReference/API_DescribeResourcePolicy.html

The primary key for this table is **account_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|content|`utf8`|
|resource_policy_summary|`json`|