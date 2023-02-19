# Table: aws_support_trusted_advisor_checks

https://docs.aws.amazon.com/awssupport/latest/APIReference/API_DescribeTrustedAdvisorChecks.html

The composite primary key for this table is (**account_id**, **region**, **language_code**, **id**).

## Relations

The following tables depend on aws_support_trusted_advisor_checks:
  - [aws_support_trusted_advisor_check_results](aws_support_trusted_advisor_check_results.md)
  - [aws_support_trusted_advisor_check_summaries](aws_support_trusted_advisor_check_summaries.md)

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
|category|String|
|description|String|
|id (PK)|String|
|metadata|StringArray|
|name|String|