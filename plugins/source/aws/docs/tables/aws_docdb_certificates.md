# Table: aws_docdb_certificates

This table shows data for Amazon DocumentDB Certificates.

https://docs.aws.amazon.com/documentdb/latest/developerguide/API_Certificate.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|certificate_arn|`utf8`|
|certificate_identifier|`utf8`|
|certificate_type|`utf8`|
|thumbprint|`utf8`|
|valid_from|`timestamp[us, tz=UTC]`|
|valid_till|`timestamp[us, tz=UTC]`|