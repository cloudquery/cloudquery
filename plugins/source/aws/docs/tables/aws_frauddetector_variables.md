# Table: aws_frauddetector_variables

This table shows data for Amazon Fraud Detector Variables.

https://docs.aws.amazon.com/frauddetector/latest/api/API_Variable.html

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
|tags|`json`|
|created_time|`utf8`|
|data_source|`utf8`|
|data_type|`utf8`|
|default_value|`utf8`|
|description|`utf8`|
|last_updated_time|`utf8`|
|name|`utf8`|
|variable_type|`utf8`|