# Table: aws_acm_certificates

This table shows data for Amazon Certificate Manager (ACM) Certificates.

https://docs.aws.amazon.com/acm/latest/APIReference/API_CertificateDetail.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|tags|`json`|
|certificate_authority_arn|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|domain_name|`utf8`|
|domain_validation_options|`json`|
|extended_key_usages|`json`|
|failure_reason|`utf8`|
|imported_at|`timestamp[us, tz=UTC]`|
|in_use_by|`list<item: utf8, nullable>`|
|issued_at|`timestamp[us, tz=UTC]`|
|issuer|`utf8`|
|key_algorithm|`utf8`|
|key_usages|`json`|
|not_after|`timestamp[us, tz=UTC]`|
|not_before|`timestamp[us, tz=UTC]`|
|options|`json`|
|renewal_eligibility|`utf8`|
|renewal_summary|`json`|
|revocation_reason|`utf8`|
|revoked_at|`timestamp[us, tz=UTC]`|
|serial|`utf8`|
|signature_algorithm|`utf8`|
|status|`utf8`|
|subject|`utf8`|
|subject_alternative_names|`list<item: utf8, nullable>`|
|type|`utf8`|