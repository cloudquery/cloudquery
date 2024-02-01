# Table: aws_inspector_findings

This table shows data for Inspector Findings.

https://docs.aws.amazon.com/inspector/v1/APIReference/API_Finding.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|attributes|`json`|
|user_attributes|`json`|
|created_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|
|asset_attributes|`json`|
|asset_type|`utf8`|
|confidence|`int64`|
|description|`utf8`|
|id|`utf8`|
|indicator_of_compromise|`bool`|
|numeric_severity|`float64`|
|recommendation|`utf8`|
|schema_version|`int64`|
|service|`utf8`|
|service_attributes|`json`|
|severity|`utf8`|
|title|`utf8`|