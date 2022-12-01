# Table: aws_qldb_ledgers



The primary key for this table is **arn**.

## Relations

The following tables depend on aws_qldb_ledgers:
  - [aws_qldb_ledger_journal_kinesis_streams](aws_qldb_ledger_journal_kinesis_streams.md)
  - [aws_qldb_ledger_journal_s3_exports](aws_qldb_ledger_journal_s3_exports.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|tags|JSON|
|arn (PK)|String|
|creation_date_time|Timestamp|
|deletion_protection|Bool|
|encryption_description|JSON|
|name|String|
|permissions_mode|String|
|state|String|
|result_metadata|JSON|