# Table: aws_acmpca_certificate_authorities

This table shows data for AWS Certificate Manager Private Certificate Authority (ACM PCA) Certificate Authorities.

https://docs.aws.amazon.com/privateca/latest/APIReference/API_CertificateAuthority.html

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
|certificate_authority_configuration|JSON|
|created_at|Timestamp|
|failure_reason|String|
|key_storage_security_standard|String|
|last_state_change_at|Timestamp|
|not_after|Timestamp|
|not_before|Timestamp|
|owner_account|String|
|restorable_until|Timestamp|
|revocation_configuration|JSON|
|serial|String|
|status|String|
|type|String|
|usage_mode|String|