# Table: aws_frauddetector_rules

https://docs.aws.amazon.com/frauddetector/latest/api/API_RuleDetail.html

The primary key for this table is **arn**.

## Relations
This table depends on [aws_frauddetector_detectors](aws_frauddetector_detectors.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|created_time|String|
|description|String|
|detector_id|String|
|expression|String|
|language|String|
|last_updated_time|String|
|outcomes|StringArray|
|rule_id|String|
|rule_version|String|