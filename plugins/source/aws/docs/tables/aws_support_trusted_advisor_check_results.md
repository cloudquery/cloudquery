# Table: aws_support_trusted_advisor_check_results

This table shows data for Support Trusted Advisor Check Results.

https://docs.aws.amazon.com/awssupport/latest/APIReference/API_DescribeTrustedAdvisorCheckResult.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**, **language_code**, **check_id**).
## Relations

This table depends on [aws_support_trusted_advisor_checks](aws_support_trusted_advisor_checks.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|language_code|`utf8`|
|category_specific_summary|`json`|
|check_id|`utf8`|
|flagged_resources|`json`|
|resources_summary|`json`|
|status|`utf8`|
|timestamp|`utf8`|