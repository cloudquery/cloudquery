# Table: aws_support_case_communications

This table shows data for Support Case Communications.

https://docs.aws.amazon.com/awssupport/latest/APIReference/API_DescribeCommunications.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_support_cases](aws_support_cases).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|attachment_set|`json`|
|body|`utf8`|
|case_id|`utf8`|
|submitted_by|`utf8`|
|time_created|`utf8`|