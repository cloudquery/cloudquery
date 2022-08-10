
# Table: aws_iot_stream_files
Represents a file to stream.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|stream_cq_id|uuid|Unique CloudQuery ID of aws_iot_streams table (FK)|
|file_id|integer|The file ID.|
|s3_location_bucket|text|The S3 bucket.|
|s3_location_key|text|The S3 key.|
|s3_location_version|text|The S3 bucket version.|
