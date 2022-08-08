
# Table: aws_rds_certificates
A CA certificate for an AWS account.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text||
|region|text|The AWS Region of the resource.|
|arn|text|The Amazon Resource Name (ARN) for the certificate.|
|certificate_identifier|text|The unique key that identifies a certificate.|
|certificate_type|text|The type of the certificate.|
|customer_override|boolean|Whether there is an override for the default certificate identifier.|
|customer_override_valid_till|timestamp without time zone|If there is an override for the default certificate identifier, when the override expires.|
|thumbprint|text|The thumbprint of the certificate.|
|valid_from|timestamp without time zone|The starting date from which the certificate is valid.|
|valid_till|timestamp without time zone|The final date that the certificate continues to be valid.|
