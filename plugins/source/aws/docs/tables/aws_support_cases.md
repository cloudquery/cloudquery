# Table: aws_support_cases

This table shows data for Support Cases.

https://docs.aws.amazon.com/awssupport/latest/APIReference/API_DescribeCases.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**, **case_id**).
## Relations

The following tables depend on aws_support_cases:
  - [aws_support_case_communications](aws_support_case_communications.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|case_id|`utf8`|
|category_code|`utf8`|
|cc_email_addresses|`list<item: utf8, nullable>`|
|display_id|`utf8`|
|language|`utf8`|
|recent_communications|`json`|
|service_code|`utf8`|
|severity_code|`utf8`|
|status|`utf8`|
|subject|`utf8`|
|submitted_by|`utf8`|
|time_created|`utf8`|