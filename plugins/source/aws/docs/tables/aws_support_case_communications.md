# Table: aws_support_case_communications

https://docs.aws.amazon.com/awssupport/latest/APIReference/API_DescribeCommunications.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_support_cases](aws_support_cases.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|attachment_set|JSON|
|body|String|
|case_id|String|
|submitted_by|String|
|time_created|String|