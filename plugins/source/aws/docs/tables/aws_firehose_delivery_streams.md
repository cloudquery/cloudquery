
# Table: aws_firehose_delivery_streams
Contains information about a delivery stream
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|tags|jsonb||
|arn|text|The Amazon Resource Name (ARN) of the delivery stream|
|delivery_stream_arn|text|The Amazon Resource Name (ARN) of the delivery stream|
|delivery_stream_name|text|The name of the delivery stream|
|delivery_stream_status|text|The status of the delivery stream|
|delivery_stream_type|text|The delivery stream type|
|version_id|text|Each time the destination is updated for a delivery stream, the version ID is changed, and the current version ID is required when updating the destination This is so that the service knows it is applying the changes to the correct version of the delivery stream|
|create_timestamp|timestamp without time zone|The date and time that the delivery stream was created|
|encryption_config_failure_description_details|text|A message providing details about the error that caused the failure|
|encryption_config_failure_description_type|text|The type of error that caused the failure|
|encryption_config_key_arn|text|If KeyType is CUSTOMER_MANAGED_CMK, this field contains the ARN of the customer managed CMK|
|encryption_config_key_type|text|Indicates the type of customer master key (CMK) that is used for encryption|
|encryption_config_status|text|This is the server-side encryption (SSE) status for the delivery stream|
|failure_description_details|text|A message providing details about the error that caused the failure|
|failure_description_type|text|The type of error that caused the failure|
|last_update_timestamp|timestamp without time zone|The date and time that the delivery stream was last updated|
|source_kinesis_stream_delivery_start_timestamp|timestamp without time zone|Kinesis Data Firehose starts retrieving records from the Kinesis data stream starting with this timestamp|
|source_kinesis_stream_kinesis_stream_arn|text|The Amazon Resource Name (ARN) of the source Kinesis data stream|
|source_kinesis_stream_role_arn|text|The ARN of the role used by the source Kinesis data stream|
