# Table: aws_lightsail_load_balancer_tls_certificates

https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_LoadBalancerTlsCertificate.html

The primary key for this table is **_cq_id**.

## Relations
This table depends on [aws_lightsail_load_balancers](aws_lightsail_load_balancers.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|load_balancer_arn|String|
|arn|String|
|created_at|Timestamp|
|domain_name|String|
|domain_validation_records|JSON|
|failure_reason|String|
|is_attached|Bool|
|issued_at|Timestamp|
|issuer|String|
|key_algorithm|String|
|load_balancer_name|String|
|location|JSON|
|name|String|
|not_after|Timestamp|
|not_before|Timestamp|
|renewal_summary|JSON|
|resource_type|String|
|revocation_reason|String|
|revoked_at|Timestamp|
|serial|String|
|signature_algorithm|String|
|status|String|
|subject|String|
|subject_alternative_names|StringArray|
|support_code|String|
|tags|JSON|