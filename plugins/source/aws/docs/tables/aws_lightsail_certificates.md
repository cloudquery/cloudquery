# Table: aws_lightsail_certificates

https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_Certificate.html

The primary key for this table is **_cq_id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|tags|JSON|
|arn|String|
|created_at|Timestamp|
|domain_name|String|
|domain_validation_records|JSON|
|eligible_to_renew|String|
|in_use_resource_count|Int|
|issued_at|Timestamp|
|issuer_ca|String|
|key_algorithm|String|
|name|String|
|not_after|Timestamp|
|not_before|Timestamp|
|renewal_summary|JSON|
|request_failure_reason|String|
|revocation_reason|String|
|revoked_at|Timestamp|
|serial_number|String|
|status|String|
|subject_alternative_names|StringArray|
|support_code|String|