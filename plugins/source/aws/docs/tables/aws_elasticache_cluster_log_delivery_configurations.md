
# Table: aws_elasticache_cluster_log_delivery_configurations
Returns the destination, format and type of the logs.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|cluster_cq_id|uuid|Unique CloudQuery ID of aws_elasticache_clusters table (FK)|
|cloudwatch_destination_log_group|text|The log group of the CloudWatch Logs destination|
|kinesis_firehose_destination_delivery_stream|text|The Kinesis Data Firehose delivery stream of the Kinesis Data Firehose destination|
|destination_type|text|Returns the destination type, either cloudwatch-logs or kinesis-firehose.|
|log_format|text|Returns the log format, either JSON or TEXT.|
|log_type|text|Refers to slow-log (https://redis.io/commands/slowlog) or engine-log.|
|message|text|Returns an error message for the log delivery configuration.|
|status|text|Returns the log delivery configuration status|
