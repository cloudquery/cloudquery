
# Table: aws_s3_bucket_encryption_rules
Specifies the default server-side encryption configuration.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|bucket_cq_id|uuid|Unique CloudQuery ID of aws_s3_buckets table (FK)|
|sse_algorithm|text|Server-side encryption algorithm to use for the default encryption.|
|kms_master_key_id|text|AWS Key Management Service (KMS) customer master key ID to use for the default encryption|
|bucket_key_enabled|boolean|Specifies whether Amazon S3 should use an S3 Bucket Key with server-side encryption using KMS (SSE-KMS) for new objects in the bucket|
