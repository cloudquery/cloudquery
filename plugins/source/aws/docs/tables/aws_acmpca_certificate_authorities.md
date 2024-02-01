# Table: aws_acmpca_certificate_authorities

This table shows data for AWS Certificate Manager Private Certificate Authority (ACM PCA) Certificate Authorities.

https://docs.aws.amazon.com/privateca/latest/APIReference/API_CertificateAuthority.html

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
|certificate_authority_configuration|`json`|
|created_at|`timestamp[us, tz=UTC]`|
|failure_reason|`utf8`|
|key_storage_security_standard|`utf8`|
|last_state_change_at|`timestamp[us, tz=UTC]`|
|not_after|`timestamp[us, tz=UTC]`|
|not_before|`timestamp[us, tz=UTC]`|
|owner_account|`utf8`|
|restorable_until|`timestamp[us, tz=UTC]`|
|revocation_configuration|`json`|
|serial|`utf8`|
|status|`utf8`|
|type|`utf8`|
|usage_mode|`utf8`|