# Table: aws_lightsail_certificates

This table shows data for Lightsail Certificates.

https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_Certificate.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|tags|`json`|
|arn|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|domain_name|`utf8`|
|domain_validation_records|`json`|
|eligible_to_renew|`utf8`|
|in_use_resource_count|`int64`|
|issued_at|`timestamp[us, tz=UTC]`|
|issuer_ca|`utf8`|
|key_algorithm|`utf8`|
|name|`utf8`|
|not_after|`timestamp[us, tz=UTC]`|
|not_before|`timestamp[us, tz=UTC]`|
|renewal_summary|`json`|
|request_failure_reason|`utf8`|
|revocation_reason|`utf8`|
|revoked_at|`timestamp[us, tz=UTC]`|
|serial_number|`utf8`|
|status|`utf8`|
|subject_alternative_names|`list<item: utf8, nullable>`|
|support_code|`utf8`|