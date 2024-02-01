# Table: aws_support_trusted_advisor_checks

This table shows data for Support Trusted Advisor Checks.

https://docs.aws.amazon.com/awssupport/latest/APIReference/API_DescribeTrustedAdvisorChecks.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**, **language_code**, **id**).
## Relations

The following tables depend on aws_support_trusted_advisor_checks:
  - [aws_support_trusted_advisor_check_results](aws_support_trusted_advisor_check_results.md)
  - [aws_support_trusted_advisor_check_summaries](aws_support_trusted_advisor_check_summaries.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|language_code|`utf8`|
|category|`utf8`|
|description|`utf8`|
|id|`utf8`|
|metadata|`list<item: utf8, nullable>`|
|name|`utf8`|