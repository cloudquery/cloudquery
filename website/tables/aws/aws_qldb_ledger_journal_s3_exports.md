# Table: aws_qldb_ledger_journal_s3_exports

This table shows data for Quantum Ledger Database (QLDB) Ledger Journal S3 Exports.

https://docs.aws.amazon.com/qldb/latest/developerguide/API_JournalS3ExportDescription.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_qldb_ledgers](aws_qldb_ledgers).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|ledger_arn|`utf8`|
|exclusive_end_time|`timestamp[us, tz=UTC]`|
|export_creation_time|`timestamp[us, tz=UTC]`|
|export_id|`utf8`|
|inclusive_start_time|`timestamp[us, tz=UTC]`|
|ledger_name|`utf8`|
|role_arn|`utf8`|
|s3_export_configuration|`json`|
|status|`utf8`|
|output_format|`utf8`|