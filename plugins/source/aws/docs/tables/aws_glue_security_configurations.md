
# Table: aws_glue_security_configurations
Specifies a security configuration
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|created_time_stamp|timestamp without time zone|The time at which this security configuration was created|
|cloud_watch_encryption_mode|text|The encryption mode to use for CloudWatch data|
|cloud_watch_encryption_kms_key_arn|text|The Amazon Resource Name (ARN) of the KMS key to be used to encrypt the data|
|job_bookmarks_encryption_mode|text|The encryption mode to use for job bookmarks data|
|job_bookmarks_encryption_kms_key_arn|text|The Amazon Resource Name (ARN) of the KMS key to be used to encrypt the data|
|name|text|The name of the security configuration|
