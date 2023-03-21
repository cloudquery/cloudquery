# Table: aws_docdb_certificates

This table shows data for Amazon DocumentDB Certificates.

https://docs.aws.amazon.com/documentdb/latest/developerguide/API_Certificate.html

The composite primary key for this table is (**account_id**, **arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region|String|
|arn (PK)|String|
|certificate_arn|String|
|certificate_identifier|String|
|certificate_type|String|
|thumbprint|String|
|valid_from|Timestamp|
|valid_till|Timestamp|