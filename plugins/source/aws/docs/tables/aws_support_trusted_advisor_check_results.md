# Table: aws_support_trusted_advisor_check_results

https://docs.aws.amazon.com/awssupport/latest/APIReference/API_DescribeTrustedAdvisorCheckResult.html

The composite primary key for this table is (**account_id**, **region**, **language_code**, **check_id**).

## Relations

This table depends on [aws_support_trusted_advisor_checks](aws_support_trusted_advisor_checks.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|language_code (PK)|String|
|category_specific_summary|JSON|
|check_id (PK)|String|
|flagged_resources|JSON|
|resources_summary|JSON|
|status|String|
|timestamp|String|