
# Table: aws_iot_topic_rules
The output from the GetTopicRule operation.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|tags|jsonb|Tags of the resource|
|aws_iot_sql_version|text|The version of the SQL rules engine to use when evaluating the rule.|
|created_at|timestamp without time zone|The date and time the rule was created.|
|description|text|The description of the rule.|
|error_action_cloudwatch_alarm_name|text|The CloudWatch alarm name.|
|error_action_cloudwatch_alarm_role_arn|text|The IAM role that allows access to the CloudWatch alarm.|
|error_action_cloudwatch_alarm_state_reason|text|The reason for the alarm change.|
|error_action_cloudwatch_alarm_state_value|text|The value of the alarm state|
|error_action_cloudwatch_logs_log_group_name|text|The CloudWatch log group to which the action sends data.|
|error_action_cloudwatch_logs_role_arn|text|The IAM role that allows access to the CloudWatch log.|
|error_action_cloudwatch_metric_metric_name|text|The CloudWatch metric name.|
|error_action_cloudwatch_metric_metric_namespace|text|The CloudWatch metric namespace name.|
|error_action_cloudwatch_metric_unit|text|The metric unit (https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/cloudwatch_concepts.html#Unit) supported by CloudWatch.|
|error_action_cloudwatch_metric_value|text|The CloudWatch metric value.|
|error_action_cloudwatch_metric_role_arn|text|The IAM role that allows access to the CloudWatch metric.|
|error_action_cloudwatch_metric_timestamp|text|An optional Unix timestamp (https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/cloudwatch_concepts.html#about_timestamp).|
|error_action_dynamo_db_hash_key_field|text|The hash key name.|
|error_action_dynamo_db_hash_key_value|text|The hash key value.|
|error_action_dynamo_db_role_arn|text|The ARN of the IAM role that grants access to the DynamoDB table.|
|error_action_dynamo_db_table_name|text|The name of the DynamoDB table.|
|error_action_dynamo_db_hash_key_type|text|The hash key type|
|error_action_dynamo_db_operation|text|The type of operation to be performed|
|error_action_dynamo_db_payload_field|text|The action payload|
|error_action_dynamo_db_range_key_field|text|The range key name.|
|error_action_dynamo_db_range_key_type|text|The range key type|
|error_action_dynamo_db_range_key_value|text|The range key value.|
|error_action_dynamo_db_v2_put_item_table_name|text|The table where the message data will be written.|
|error_action_dynamo_db_v2_role_arn|text|The ARN of the IAM role that grants access to the DynamoDB table.|
|error_action_elasticsearch_endpoint|text|The endpoint of your OpenSearch domain.|
|error_action_elasticsearch_id|text|The unique identifier for the document you are storing.|
|error_action_elasticsearch_index|text|The index where you want to store your data.|
|error_action_elasticsearch_role_arn|text|The IAM role ARN that has access to OpenSearch.|
|error_action_elasticsearch_type|text|The type of document you are storing.|
|error_action_firehose_delivery_stream_name|text|The delivery stream name.|
|error_action_firehose_role_arn|text|The IAM role that grants access to the Amazon Kinesis Firehose stream.|
|error_action_firehose_batch_mode|boolean|Whether to deliver the Kinesis Data Firehose stream as a batch by using PutRecordBatch (https://docs.aws.amazon.com/firehose/latest/APIReference/API_PutRecordBatch.html). The default value is false|
|error_action_firehose_separator|text|A character separator that will be used to separate records written to the Firehose stream|
|error_action_http_url|text|The endpoint URL|
|error_action_http_auth_sigv4_role_arn|text|The ARN of the signing role.|
|error_action_http_auth_sigv4_service_name|text|The service name to use while signing with Sig V4.|
|error_action_http_auth_sigv4_signing_region|text|The signing region.|
|error_action_http_confirmation_url|text|The URL to which IoT sends a confirmation message|
|error_action_http_headers|jsonb|The HTTP headers to send with the message data.|
|error_action_iot_analytics_batch_mode|boolean|Whether to process the action as a batch|
|error_action_iot_analytics_channel_arn|text|(deprecated) The ARN of the IoT Analytics channel to which message data will be sent.|
|error_action_iot_analytics_channel_name|text|The name of the IoT Analytics channel to which message data will be sent.|
|error_action_iot_analytics_role_arn|text|The ARN of the role which has a policy that grants IoT Analytics permission to send message data via IoT Analytics (iotanalytics:BatchPutMessage).|
|error_action_iot_events_input_name|text|The name of the IoT Events input.|
|error_action_iot_events_role_arn|text|The ARN of the role that grants IoT permission to send an input to an IoT Events detector|
|error_action_iot_events_batch_mode|boolean|Whether to process the event actions as a batch|
|error_action_iot_events_message_id|text|The ID of the message|
|error_action_iot_site_wise|jsonb|Sends data from the MQTT message that triggered the rule to IoT SiteWise asset properties.|
|error_action_kafka_client_properties|jsonb|Properties of the Apache Kafka producer client.|
|error_action_kafka_destination_arn|text|The ARN of Kafka action's VPC TopicRuleDestination.|
|error_action_kafka_topic|text|The Kafka topic for messages to be sent to the Kafka broker.|
|error_action_kafka_key|text|The Kafka message key.|
|error_action_kafka_partition|text|The Kafka message partition.|
|error_action_kinesis_role_arn|text|The ARN of the IAM role that grants access to the Amazon Kinesis stream.|
|error_action_kinesis_stream_name|text|The name of the Amazon Kinesis stream.|
|error_action_kinesis_partition_key|text|The partition key.|
|error_action_lambda_function_arn|text|The ARN of the Lambda function.|
|error_action_open_search_endpoint|text|The endpoint of your OpenSearch domain.|
|error_action_open_search_id|text|The unique identifier for the document you are storing.|
|error_action_open_search_index|text|The OpenSearch index where you want to store your data.|
|error_action_open_search_role_arn|text|The IAM role ARN that has access to OpenSearch.|
|error_action_open_search_type|text|The type of document you are storing.|
|error_action_republish_role_arn|text|The ARN of the IAM role that grants access.|
|error_action_republish_topic|text|The name of the MQTT topic.|
|error_action_republish_qos|integer|The Quality of Service (QoS) level to use when republishing messages|
|error_action_s3_bucket_name|text|The Amazon S3 bucket.|
|error_action_s3_key|text|The object key|
|error_action_s3_role_arn|text|The ARN of the IAM role that grants access.|
|error_action_s3_canned_acl|text|The Amazon S3 canned ACL that controls access to the object identified by the object key|
|error_action_salesforce_token|text|The token used to authenticate access to the Salesforce IoT Cloud Input Stream. The token is available from the Salesforce IoT Cloud platform after creation of the Input Stream.|
|error_action_salesforce_url|text|The URL exposed by the Salesforce IoT Cloud Input Stream|
|error_action_sns_role_arn|text|The ARN of the IAM role that grants access.|
|error_action_sns_target_arn|text|The ARN of the SNS topic.|
|error_action_sns_message_format|text|(Optional) The message format of the message to publish|
|error_action_sqs_queue_url|text|The URL of the Amazon SQS queue.|
|error_action_sqs_role_arn|text|The ARN of the IAM role that grants access.|
|error_action_sqs_use_base64|boolean|Specifies whether to use Base64 encoding.|
|error_action_step_functions_role_arn|text|The ARN of the role that grants IoT permission to start execution of a state machine ("Action":"states:StartExecution").|
|error_action_step_functions_state_machine_name|text|The name of the Step Functions state machine whose execution will be started.|
|error_action_step_functions_execution_name_prefix|text|(Optional) A name will be given to the state machine execution consisting of this prefix followed by a UUID|
|error_action_timestream_database_name|text|The name of an Amazon Timestream database.|
|error_action_timestream_dimensions|jsonb|Metadata attributes of the time series that are written in each measure record.|
|error_action_timestream_role_arn|text|The ARN of the role that grants permission to write to the Amazon Timestream database table.|
|error_action_timestream_table_name|text|The name of the database table into which to write the measure records.|
|error_action_timestream_timestamp_unit|text|The precision of the timestamp value that results from the expression described in value|
|error_action_timestream_timestamp_value|text|An expression that returns a long epoch time value.|
|rule_disabled|boolean|Specifies whether the rule is disabled.|
|rule_name|text|The name of the rule.|
|sql|text|The SQL statement used to query the topic|
|arn|text|The rule ARN.|
