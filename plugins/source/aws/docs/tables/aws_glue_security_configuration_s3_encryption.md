
# Table: aws_glue_security_configuration_s3_encryption
Specifies how Amazon Simple Storage Service (Amazon S3) data should be encrypted
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|security_configuration_cq_id|uuid|Unique CloudQuery ID of aws_glue_security_configurations table (FK)|
|kms_key_arn|text|The Amazon Resource Name (ARN) of the KMS key to be used to encrypt the data|
|s3_encryption_mode|text|The encryption mode to use for Amazon S3 data|
