# Table: aws_lightsail_load_balancer_tls_certificates

This table shows data for Lightsail Load Balancer TLS Certificates.

https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_LoadBalancerTlsCertificate.html

The primary key for this table is **arn**.

## Relations

This table depends on [aws_lightsail_load_balancers](aws_lightsail_load_balancers).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|load_balancer_arn|`utf8`|
|tags|`json`|
|arn (PK)|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|domain_name|`utf8`|
|domain_validation_records|`json`|
|failure_reason|`utf8`|
|is_attached|`bool`|
|issued_at|`timestamp[us, tz=UTC]`|
|issuer|`utf8`|
|key_algorithm|`utf8`|
|load_balancer_name|`utf8`|
|location|`json`|
|name|`utf8`|
|not_after|`timestamp[us, tz=UTC]`|
|not_before|`timestamp[us, tz=UTC]`|
|renewal_summary|`json`|
|resource_type|`utf8`|
|revocation_reason|`utf8`|
|revoked_at|`timestamp[us, tz=UTC]`|
|serial|`utf8`|
|signature_algorithm|`utf8`|
|status|`utf8`|
|subject|`utf8`|
|subject_alternative_names|`list<item: utf8, nullable>`|
|support_code|`utf8`|