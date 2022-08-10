
# Table: aws_iot_topic_rule_actions
Describes the actions associated with a rule.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|topic_rule_cq_id|uuid|Unique CloudQuery ID of aws_iot_topic_rules table (FK)|
|cloudwatch_alarm_alarm_name|text|The CloudWatch alarm name.|
|cloudwatch_alarm_role_arn|text|The IAM role that allows access to the CloudWatch alarm.|
|cloudwatch_alarm_state_reason|text|The reason for the alarm change.|
|cloudwatch_alarm_state_value|text|The value of the alarm state|
|cloudwatch_logs_log_group_name|text|The CloudWatch log group to which the action sends data.|
|cloudwatch_logs_role_arn|text|The IAM role that allows access to the CloudWatch log.|
|cloudwatch_metric_metric_name|text|The CloudWatch metric name.|
|cloudwatch_metric_metric_namespace|text|The CloudWatch metric namespace name.|
|cloudwatch_metric_metric_unit|text|The metric unit (https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/cloudwatch_concepts.html#Unit) supported by CloudWatch.|
|cloudwatch_metric_metric_value|text|The CloudWatch metric value.|
|cloudwatch_metric_role_arn|text|The IAM role that allows access to the CloudWatch metric.|
|cloudwatch_metric_metric_timestamp|text|An optional Unix timestamp (https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/cloudwatch_concepts.html#about_timestamp).|
|dynamo_db_hash_key_field|text|The hash key name.|
|dynamo_db_hash_key_value|text|The hash key value.|
|dynamo_db_role_arn|text|The ARN of the IAM role that grants access to the DynamoDB table.|
|dynamo_db_table_name|text|The name of the DynamoDB table.|
|dynamo_db_hash_key_type|text|The hash key type|
|dynamo_db_operation|text|The type of operation to be performed|
|dynamo_db_payload_field|text|The action payload|
|dynamo_db_range_key_field|text|The range key name.|
|dynamo_db_range_key_type|text|The range key type|
|dynamo_db_range_key_value|text|The range key value.|
|dynamo_db_v2_put_item_table_name|text|The table where the message data will be written.|
|dynamo_db_v2_role_arn|text|The ARN of the IAM role that grants access to the DynamoDB table.|
|elasticsearch_endpoint|text|The endpoint of your OpenSearch domain.|
|elasticsearch_id|text|The unique identifier for the document you are storing.|
|elasticsearch_index|text|The index where you want to store your data.|
|elasticsearch_role_arn|text|The IAM role ARN that has access to OpenSearch.|
|elasticsearch_type|text|The type of document you are storing.|
|firehose_delivery_stream_name|text|The delivery stream name.|
|firehose_role_arn|text|The IAM role that grants access to the Amazon Kinesis Firehose stream.|
|firehose_batch_mode|boolean|Whether to deliver the Kinesis Data Firehose stream as a batch by using PutRecordBatch (https://docs.aws.amazon.com/firehose/latest/APIReference/API_PutRecordBatch.html). The default value is false|
|firehose_separator|text|A character separator that will be used to separate records written to the Firehose stream|
|http_url|text|The endpoint URL|
|http_auth_sigv4_role_arn|text|The ARN of the signing role.|
|http_auth_sigv4_service_name|text|The service name to use while signing with Sig V4.|
|http_auth_sigv4_signing_region|text|The signing region.|
|http_confirmation_url|text|The URL to which IoT sends a confirmation message|
|http_headers|jsonb|The HTTP headers to send with the message data.|
|iot_analytics_batch_mode|boolean|Whether to process the action as a batch|
|iot_analytics_channel_arn|text|(deprecated) The ARN of the IoT Analytics channel to which message data will be sent.|
|iot_analytics_channel_name|text|The name of the IoT Analytics channel to which message data will be sent.|
|iot_analytics_role_arn|text|The ARN of the role which has a policy that grants IoT Analytics permission to send message data via IoT Analytics (iotanalytics:BatchPutMessage).|
|iot_events_input_name|text|The name of the IoT Events input.|
|iot_events_role_arn|text|The ARN of the role that grants IoT permission to send an input to an IoT Events detector|
|iot_events_batch_mode|boolean|Whether to process the event actions as a batch|
|iot_events_message_id|text|The ID of the message|
|iot_site_wise|jsonb|Sends data from the MQTT message that triggered the rule to IoT SiteWise asset properties.|
|kafka_client_properties|jsonb|Properties of the Apache Kafka producer client.|
|kafka_destination_arn|text|The ARN of Kafka action's VPC TopicRuleDestination.|
|kafka_topic|text|The Kafka topic for messages to be sent to the Kafka broker.|
|kafka_key|text|The Kafka message key.|
|kafka_partition|text|The Kafka message partition.|
|kinesis_role_arn|text|The ARN of the IAM role that grants access to the Amazon Kinesis stream.|
|kinesis_stream_name|text|The name of the Amazon Kinesis stream.|
|kinesis_partition_key|text|The partition key.|
|lambda_function_arn|text|The ARN of the Lambda function.|
|open_search_endpoint|text|The endpoint of your OpenSearch domain.|
|open_search_id|text|The unique identifier for the document you are storing.|
|open_search_index|text|The OpenSearch index where you want to store your data.|
|open_search_role_arn|text|The IAM role ARN that has access to OpenSearch.|
|open_search_type|text|The type of document you are storing.|
|republish_role_arn|text|The ARN of the IAM role that grants access.|
|republish_topic|text|The name of the MQTT topic.|
|republish_qos|integer|The Quality of Service (QoS) level to use when republishing messages|
|s3_bucket_name|text|The Amazon S3 bucket.|
|s3_key|text|The object key|
|s3_role_arn|text|The ARN of the IAM role that grants access.|
|s3_canned_acl|text|The Amazon S3 canned ACL that controls access to the object identified by the object key|
|salesforce_token|text|The token used to authenticate access to the Salesforce IoT Cloud Input Stream. The token is available from the Salesforce IoT Cloud platform after creation of the Input Stream.|
|salesforce_url|text|The URL exposed by the Salesforce IoT Cloud Input Stream|
|sns_role_arn|text|The ARN of the IAM role that grants access.|
|sns_target_arn|text|The ARN of the SNS topic.|
|sns_message_format|text|(Optional) The message format of the message to publish|
|sqs_queue_url|text|The URL of the Amazon SQS queue.|
|sqs_role_arn|text|The ARN of the IAM role that grants access.|
|sqs_use_base64|boolean|Specifies whether to use Base64 encoding.|
|step_functions_role_arn|text|The ARN of the role that grants IoT permission to start execution of a state machine ("Action":"states:StartExecution").|
|step_functions_state_machine_name|text|The name of the Step Functions state machine whose execution will be started.|
|step_functions_execution_name_prefix|text|(Optional) A name will be given to the state machine execution consisting of this prefix followed by a UUID|
|timestream_database_name|text|The name of an Amazon Timestream database.|
|timestream_dimensions|jsonb|Metadata attributes of the time series that are written in each measure record.|
|timestream_role_arn|text|The ARN of the role that grants permission to write to the Amazon Timestream database table.|
|timestream_table_name|text|The name of the database table into which to write the measure records.|
|timestream_timestamp_unit|text|The precision of the timestamp value that results from the expression described in value|
|timestream_timestamp_value|text|An expression that returns a long epoch time value.|
