# Table: aws_wafregional_web_acls

https://docs.aws.amazon.com/waf/latest/APIReference/API_wafRegional_WebACL.html

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
|resources_for_web_acl|StringArray|
|default_action|JSON|
|rules|JSON|
|web_acl_id|String|
|metric_name|String|
|name|String|