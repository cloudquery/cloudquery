# Table: aws_support_severity_levels

https://docs.aws.amazon.com/awssupport/latest/APIReference/API_DescribeSeverityLevels.html

The composite primary key for this table is (**account_id**, **region**, **language_code**, **code**).

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
|code (PK)|String|
|name|String|