# Table: aws_qldb_ledger_journal_kinesis_streams

This table shows data for Quantum Ledger Database (QLDB) Ledger Journal Kinesis Streams.

https://docs.aws.amazon.com/qldb/latest/developerguide/API_JournalKinesisStreamDescription.html

The primary key for this table is **arn**.

## Relations

This table depends on [aws_qldb_ledgers](aws_qldb_ledgers).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|ledger_arn|`utf8`|
|kinesis_configuration|`json`|
|ledger_name|`utf8`|
|role_arn|`utf8`|
|status|`utf8`|
|stream_id|`utf8`|
|stream_name|`utf8`|
|arn (PK)|`utf8`|
|creation_time|`timestamp[us, tz=UTC]`|
|error_cause|`utf8`|
|exclusive_end_time|`timestamp[us, tz=UTC]`|
|inclusive_start_time|`timestamp[us, tz=UTC]`|