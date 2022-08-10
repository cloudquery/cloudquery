
# Table: aws_glue_datacatalog_encryption_settings
Contains configuration information for maintaining Data Catalog security
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|return_connection_password_encrypted|boolean|When the ReturnConnectionPasswordEncrypted flag is set to "true", passwords remain encrypted in the responses of GetConnection and GetConnections|
|aws_kms_key_id|text|An KMS key that is used to encrypt the connection password|
|encryption_at_rest_catalog_encryption_mode|text|The encryption-at-rest mode for encrypting Data Catalog data|
|encryption_at_rest_sse_aws_kms_key_id|text|The ID of the KMS key to use for encryption at rest|
