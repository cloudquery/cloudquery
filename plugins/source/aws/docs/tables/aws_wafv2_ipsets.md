# Table: aws_wafv2_ipsets

https://docs.aws.amazon.com/waf/latest/APIReference/API_IPSet.html

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
|addresses|InetArray|
|tags|JSON|
|arn (PK)|String|
|ip_address_version|String|
|id|String|
|name|String|
|description|String|