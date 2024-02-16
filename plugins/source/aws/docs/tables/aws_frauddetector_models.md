# Table: aws_frauddetector_models

This table shows data for Amazon Fraud Detector Models.

https://docs.aws.amazon.com/frauddetector/latest/api/API_Model.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.
## Relations

The following tables depend on aws_frauddetector_models:
  - [aws_frauddetector_model_versions](aws_frauddetector_model_versions.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|created_time|`utf8`|
|description|`utf8`|
|event_type_name|`utf8`|
|last_updated_time|`utf8`|
|model_id|`utf8`|
|model_type|`utf8`|