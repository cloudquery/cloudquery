
# Table: aws_ecr_repositories
An object representing a repository.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|tags|jsonb||
|created_at|timestamp without time zone|The date and time, in JavaScript date format, when the repository was created.|
|encryption_configuration_encryption_type|text|The encryption type to use|
|encryption_configuration_kms_key|text|If you use the KMS encryption type, specify the KMS key to use for encryption. The alias, key ID, or full ARN of the KMS key can be specified|
|image_scanning_configuration_scan_on_push|boolean|The setting that determines whether images are scanned after being pushed to a repository|
|image_tag_mutability|text|The tag mutability setting for the repository.|
|registry_id|text|The Amazon Web Services account ID associated with the registry that contains the repository.|
|arn|text|The Amazon Resource Name (ARN) that identifies the repository|
|name|text|The name of the repository.|
|uri|text|The URI for the repository|
