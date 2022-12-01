# Table: aws_qldb_ledger_journal_s3_exports

https://docs.aws.amazon.com/qldb/latest/developerguide/API_JournalS3ExportDescription.html

The primary key for this table is **_cq_id**.

## Relations
This table depends on [aws_qldb_ledgers](aws_qldb_ledgers.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|ledger_arn|String|
|exclusive_end_time|Timestamp|
|export_creation_time|Timestamp|
|export_id|String|
|inclusive_start_time|Timestamp|
|ledger_name|String|
|role_arn|String|
|s3_export_configuration|JSON|
|status|String|
|output_format|String|