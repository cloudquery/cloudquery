
# Table: aws_iam_server_certificates
Contains information about a server certificate without its certificate body, certificate chain, and private key.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|arn|text|The Amazon Resource Name (ARN) specifying the server certificate. For more information about ARNs and how to use them in policies, see IAM identifiers (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the IAM User Guide.|
|path|text|The path to the server certificate. For more information about paths, see IAM identifiers (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the IAM User Guide.|
|server_certificate_id|text|The stable and unique string identifying the server certificate. For more information about IDs, see IAM identifiers (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the IAM User Guide.|
|server_certificate_name|text|The name that identifies the server certificate.|
|expiration|timestamp without time zone|The date on which the certificate is set to expire. |
|upload_date|timestamp without time zone|The date when the server certificate was uploaded. |
