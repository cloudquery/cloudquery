# Table: aws_xray_sampling_rules

This table shows data for AWS X-Ray Sampling Rules.

https://docs.aws.amazon.com/xray/latest/api/API_SamplingRuleRecord.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|created_at|`timestamp[us, tz=UTC]`|
|modified_at|`timestamp[us, tz=UTC]`|
|sampling_rule|`json`|