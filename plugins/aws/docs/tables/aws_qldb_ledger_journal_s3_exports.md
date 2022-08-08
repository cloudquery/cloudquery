
# Table: aws_qldb_ledger_journal_s3_exports
Information about a journal export job, including the ledger name, export ID, creation time, current status, and the parameters of the original export creation request.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|ledger_cq_id|uuid|Unique CloudQuery ID of aws_qldb_ledgers table (FK)|
|exclusive_end_time|timestamp without time zone|The exclusive end date and time for the range of journal contents that was specified in the original export request.  This member is required.|
|export_creation_time|timestamp without time zone|The date and time, in epoch time format, when the export job was created|
|export_id|text|The UUID (represented in Base62-encoded text) of the journal export job.  This member is required.|
|inclusive_start_time|timestamp without time zone|The inclusive start date and time for the range of journal contents that was specified in the original export request.  This member is required.|
|ledger_name|text|The name of the ledger.  This member is required.|
|role_arn|text|The Amazon Resource Name (ARN) of the IAM role that grants QLDB permissions for a journal export job to do the following:  * Write objects into your Amazon Simple Storage Service (Amazon S3) bucket.  * (Optional) Use your customer managed key in Key Management Service (KMS) for server-side encryption of your exported data.  This member is required.|
|bucket|text|The Amazon S3 bucket name in which a journal export job writes the journal contents|
|object_encryption_type|text|The Amazon S3 object encryption type|
|kms_key_arn|text|The Amazon Resource Name (ARN) of a symmetric key in Key Management Service (KMS)|
|prefix|text|The prefix for the Amazon S3 bucket in which a journal export job writes the journal contents|
|status|text|The current state of the journal export job.  This member is required.|
|output_format|text|The output format of the exported journal data.|
