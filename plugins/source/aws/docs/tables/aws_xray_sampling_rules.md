# Table: aws_xray_sampling_rules

https://docs.aws.amazon.com/xray/latest/api/API_SamplingRuleRecord.html

The primary key for this table is **arn**.



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
|tags|JSON|
|created_at|Timestamp|
|modified_at|Timestamp|
|sampling_rule|JSON|