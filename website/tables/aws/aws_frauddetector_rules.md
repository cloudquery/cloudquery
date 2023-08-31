# Table: aws_frauddetector_rules

This table shows data for Amazon Fraud Detector Rules.

https://docs.aws.amazon.com/frauddetector/latest/api/API_RuleDetail.html

The primary key for this table is **arn**.

## Relations

This table depends on [aws_frauddetector_detectors](aws_frauddetector_detectors).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|created_time|`utf8`|
|description|`utf8`|
|detector_id|`utf8`|
|expression|`utf8`|
|language|`utf8`|
|last_updated_time|`utf8`|
|outcomes|`list<item: utf8, nullable>`|
|rule_id|`utf8`|
|rule_version|`utf8`|