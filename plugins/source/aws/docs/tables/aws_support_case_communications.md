# Table: aws_support_case_communications

This table shows data for Support Case Communications.

https://docs.aws.amazon.com/awssupport/latest/APIReference/API_DescribeCommunications.html

The composite primary key for this table is (**case_id**, **submitted_by**, **time_created**).

## Relations

This table depends on [aws_support_cases](aws_support_cases.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|attachment_set|`json`|
|body|`utf8`|
|case_id (PK)|`utf8`|
|submitted_by (PK)|`utf8`|
|time_created (PK)|`utf8`|