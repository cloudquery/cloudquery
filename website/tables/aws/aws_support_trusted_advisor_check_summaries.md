# Table: aws_support_trusted_advisor_check_summaries

This table shows data for Support Trusted Advisor Check Summaries.

https://docs.aws.amazon.com/awssupport/latest/APIReference/API_DescribeTrustedAdvisorCheckSummaries.html

The composite primary key for this table is (**account_id**, **region**, **language_code**, **check_id**).

## Relations

This table depends on [aws_support_trusted_advisor_checks](aws_support_trusted_advisor_checks).

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
|resources_summary|JSON|
|status|String|
|timestamp|String|
|has_flagged_resources|Bool|