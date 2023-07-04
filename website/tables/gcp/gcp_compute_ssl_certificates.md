# Table: gcp_compute_ssl_certificates

This table shows data for GCP Compute SSL Certificates.

https://cloud.google.com/compute/docs/reference/rest/v1/sslCertificates#SslCertificate

The primary key for this table is **self_link**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id|`utf8`|
|certificate|`utf8`|
|creation_timestamp|`utf8`|
|description|`utf8`|
|expire_time|`utf8`|
|id|`int64`|
|kind|`utf8`|
|managed|`json`|
|name|`utf8`|
|private_key|`utf8`|
|region|`utf8`|
|self_link (PK)|`utf8`|
|self_managed|`json`|
|subject_alternative_names|`list<item: utf8, nullable>`|
|type|`utf8`|