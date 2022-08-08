
# Table: aws_kinesis_stream_enhanced_monitoring
Represents enhanced metrics types
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|stream_cq_id|uuid|Unique CloudQuery ID of aws_kinesis_streams table (FK)|
|shard_level_metrics|text[]|List of shard-level metrics|
