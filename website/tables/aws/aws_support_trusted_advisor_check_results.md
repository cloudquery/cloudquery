# Table: aws_support_trusted_advisor_check_results

This table shows data for Support Trusted Advisor Check Results.

https://docs.aws.amazon.com/awssupport/latest/APIReference/API_DescribeTrustedAdvisorCheckResult.html

The composite primary key for this table is (**account_id**, **region**, **language_code**, **check_id**).

## Relations

This table depends on [aws_support_trusted_advisor_checks](aws_support_trusted_advisor_checks).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|language_code (PK)|`utf8`|
|category_specific_summary|`json`|
|check_id (PK)|`utf8`|
|flagged_resources|`json`|
|resources_summary|`json`|
|status|`utf8`|
|timestamp|`utf8`|