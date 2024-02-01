# Table: aws_support_severity_levels

This table shows data for Support Severity Levels.

https://docs.aws.amazon.com/awssupport/latest/APIReference/API_DescribeSeverityLevels.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**, **language_code**, **code**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|language_code|`utf8`|
|code|`utf8`|
|name|`utf8`|