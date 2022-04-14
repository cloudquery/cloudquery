
# Table: aws_qldb_ledger_journal_kinesis_streams
Information about an Amazon QLDB journal stream, including the Amazon Resource Name (ARN), stream name, creation time, current status, and the parameters of the original stream creation request.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|ledger_cq_id|uuid|Unique CloudQuery ID of aws_qldb_ledgers table (FK)|
|stream_arn|text|The Amazon Resource Name (ARN) of the Kinesis Data Streams resource.  This member is required.|
|aggregation_enabled|boolean|Enables QLDB to publish multiple data records in a single Kinesis Data Streams record, increasing the number of records sent per API call|
|ledger_name|text|The name of the ledger.  This member is required.|
|role_arn|text|The Amazon Resource Name (ARN) of the IAM role that grants QLDB permissions for a journal stream to write data records to a Kinesis Data Streams resource.  This member is required.|
|status|text|The current state of the QLDB journal stream.  This member is required.|
|stream_id|text|The UUID (represented in Base62-encoded text) of the QLDB journal stream.  This member is required.|
|stream_name|text|The user-defined name of the QLDB journal stream.  This member is required.|
|arn|text|The Amazon Resource Name (ARN) of the QLDB journal stream.|
|creation_time|timestamp without time zone|The date and time, in epoch time format, when the QLDB journal stream was created|
|error_cause|text|The error message that describes the reason that a stream has a status of IMPAIRED or FAILED|
|exclusive_end_time|timestamp without time zone|The exclusive date and time that specifies when the stream ends|
|inclusive_start_time|timestamp without time zone|The inclusive start date and time from which to start streaming journal data.|
