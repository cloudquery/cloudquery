# Table: aws_docdb_certificates

This table shows data for Amazon DocumentDB Certificates.

https://docs.aws.amazon.com/documentdb/latest/developerguide/API_Certificate.html

The composite primary key for this table is (**account_id**, **arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|certificate_arn|`utf8`|
|certificate_identifier|`utf8`|
|certificate_type|`utf8`|
|thumbprint|`utf8`|
|valid_from|`timestamp[us, tz=UTC]`|
|valid_till|`timestamp[us, tz=UTC]`|