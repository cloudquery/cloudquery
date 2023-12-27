# Table: aws_iot_certificates

This table shows data for AWS IoT Certificates.

https://docs.aws.amazon.com/iot/latest/apireference/API_CertificateDescription.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|policies|`list<item: utf8, nullable>`|
|arn (PK)|`utf8`|
|ca_certificate_id|`utf8`|
|certificate_arn|`utf8`|
|certificate_id|`utf8`|
|certificate_mode|`utf8`|
|certificate_pem|`utf8`|
|creation_date|`timestamp[us, tz=UTC]`|
|customer_version|`int64`|
|generation_id|`utf8`|
|last_modified_date|`timestamp[us, tz=UTC]`|
|owned_by|`utf8`|
|previous_owned_by|`utf8`|
|status|`utf8`|
|transfer_data|`json`|
|validity|`json`|