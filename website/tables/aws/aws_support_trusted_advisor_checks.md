# Table: aws_support_trusted_advisor_checks

This table shows data for Support Trusted Advisor Checks.

https://docs.aws.amazon.com/awssupport/latest/APIReference/API_DescribeTrustedAdvisorChecks.html

The composite primary key for this table is (**account_id**, **region**, **language_code**, **id**).

## Relations

The following tables depend on aws_support_trusted_advisor_checks:
  - [aws_support_trusted_advisor_check_results](aws_support_trusted_advisor_check_results)
  - [aws_support_trusted_advisor_check_summaries](aws_support_trusted_advisor_check_summaries)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|language_code (PK)|`utf8`|
|category|`utf8`|
|description|`utf8`|
|id (PK)|`utf8`|
|metadata|`list<item: utf8, nullable>`|
|name|`utf8`|