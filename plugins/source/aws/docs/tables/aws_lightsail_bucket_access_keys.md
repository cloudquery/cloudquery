
# Table: aws_lightsail_bucket_access_keys
Describes an access key for an Amazon Lightsail bucket
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|bucket_cq_id|uuid|Unique CloudQuery ID of aws_lightsail_buckets table (FK)|
|access_key_id|text|The ID of the access key|
|created_at|timestamp without time zone|The timestamp when the access key was created|
|last_used_date|timestamp without time zone|The date and time when the access key was most recently used|
|last_used_region|text|The AWS Region where this access key was most recently used|
|last_used_service_name|text|The name of the AWS service with which this access key was most recently used This value is N/A if the access key has not been used|
|secret_access_key|text|The secret access key used to sign requests|
|status|text|The status of the access key|
