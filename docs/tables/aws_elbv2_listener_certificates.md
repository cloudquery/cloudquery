
# Table: aws_elbv2_listener_certificates
Information about an SSL server certificate.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|listener_cq_id|uuid|Unique CloudQuery ID of aws_elbv2_listeners table (FK)|
|certificate_arn|text|The Amazon Resource Name (ARN) of the certificate.|
|is_default|boolean|Indicates whether the certificate is the default certificate|
