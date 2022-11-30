# Table: aws_inspector_findings

https://docs.aws.amazon.com/inspector/v1/APIReference/API_Finding.html

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
|attributes|JSON|
|user_attributes|JSON|
|created_at|Timestamp|
|updated_at|Timestamp|
|asset_attributes|JSON|
|asset_type|String|
|confidence|Int|
|description|String|
|id|String|
|indicator_of_compromise|Bool|
|numeric_severity|Float|
|recommendation|String|
|schema_version|Int|
|service|String|
|service_attributes|JSON|
|severity|String|
|title|String|