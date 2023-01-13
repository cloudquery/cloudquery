# Table: aws_waf_web_acls

https://docs.aws.amazon.com/waf/latest/APIReference/API_waf_WebACLSummary.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|arn (PK)|String|
|tags|JSON|
|default_action|JSON|
|rules|JSON|
|web_acl_id|String|
|metric_name|String|
|name|String|
|web_acl_arn|String|
|logging_configuration|JSON|