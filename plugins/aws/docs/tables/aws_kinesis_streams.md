
# Table: aws_kinesis_streams
Represents the output for DescribeStreamSummary
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|arn|text||
|tags|jsonb||
|open_shard_count|bigint|The number of open shards in the stream|
|retention_period_hours|bigint|The current retention period, in hours|
|stream_arn|text|The Amazon Resource Name (ARN) for the stream being described|
|stream_creation_timestamp|timestamp without time zone|The approximate time that the stream was created|
|stream_name|text|The name of the stream being described|
|stream_status|text|The current status of the stream being described|
|consumer_count|bigint|The number of enhanced fan-out consumers registered with the stream|
|encryption_type|text|The encryption type used|
|key_id|text|The GUID for the customer-managed Amazon Web Services KMS key to use for encryption|
|stream_mode_details_stream_mode|text|Specifies the capacity mode to which you want to set your data stream Currently, in Kinesis Data Streams, you can choose between an on-demand capacity mode and a provisioned capacity mode for your data streams|
