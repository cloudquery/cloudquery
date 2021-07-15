
# Table: aws_fsx_backups
A backup of an Amazon FSx file system.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|id|text|The ID of the backup.|
|creation_time|timestamp without time zone|The time when a particular backup was created.|
|lifecycle|text|The lifecycle status of the backup.|
|type|text|The type of the file system backup.|
|directory_information_active_directory_id|text|The ID of the AWS Managed Microsoft Active Directory instance to which the file system is joined.|
|directory_information_domain_name|text|The fully qualified domain name of the self-managed AD directory.|
|failure_details_message|text|A message describing the backup creation failure.|
|kms_key_id|text|The ID of the AWS Key Management Service (AWS KMS) key used to encrypt the backup of the Amazon FSx file system's data at rest.|
|progress_percent|integer|The current percent of progress of an asynchronous task.|
|arn|text|The Amazon Resource Name (ARN) for the backup resource.|
|tags|jsonb|Tags associated with a particular file system.|
