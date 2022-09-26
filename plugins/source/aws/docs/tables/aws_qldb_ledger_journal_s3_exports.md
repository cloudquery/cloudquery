# Table: aws_qldb_ledger_journal_s3_exports


The primary key for this table is **_cq_id**.

## Relations
This table depends on [`aws_qldb_ledgers`](aws_qldb_ledgers.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
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
|_cq_id (PK)|UUID|
|_cq_fetch_time|Timestamp|