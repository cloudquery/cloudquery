# Table: aws_support_severity_levels

This table shows data for Support Severity Levels.

https://docs.aws.amazon.com/awssupport/latest/APIReference/API_DescribeSeverityLevels.html

The composite primary key for this table is (**account_id**, **region**, **language_code**, **code**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|language_code (PK)|`utf8`|
|code (PK)|`utf8`|
|name|`utf8`|