# Table: aws_qldb_ledgers

This table shows data for Quantum Ledger Database (QLDB) Ledgers.

https://docs.aws.amazon.com/qldb/latest/developerguide/API_DescribeLedger.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.
## Relations

The following tables depend on aws_qldb_ledgers:
  - [aws_qldb_ledger_journal_kinesis_streams](aws_qldb_ledger_journal_kinesis_streams.md)
  - [aws_qldb_ledger_journal_s3_exports](aws_qldb_ledger_journal_s3_exports.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|tags|`json`|
|arn|`utf8`|
|creation_date_time|`timestamp[us, tz=UTC]`|
|deletion_protection|`bool`|
|encryption_description|`json`|
|name|`utf8`|
|permissions_mode|`utf8`|
|state|`utf8`|