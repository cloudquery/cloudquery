# Table: aws_support_trusted_advisor_check_summaries

https://docs.aws.amazon.com/awssupport/latest/APIReference/API_DescribeTrustedAdvisorCheckSummaries.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_support_trusted_advisor_checks](aws_support_trusted_advisor_checks.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|category_specific_summary|JSON|
|check_id|String|
|resources_summary|JSON|
|status|String|
|timestamp|String|
|has_flagged_resources|Bool|