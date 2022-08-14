
# Table: aws_iot_ca_certificates
Describes a CA certificate.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|certificates|text[]|Certificates of the ca certificate|
|auto_registration_status|text|Whether the CA certificate configured for auto registration of device certificates|
|arn|text|The CA certificate ARN.|
|id|text|The CA certificate ID.|
|pem|text|The CA certificate data, in PEM format.|
|creation_date|timestamp without time zone|The date the CA certificate was created.|
|customer_version|integer|The customer version of the CA certificate.|
|generation_id|text|The generation ID of the CA certificate.|
|last_modified_date|timestamp without time zone|The date the CA certificate was last modified.|
|owned_by|text|The owner of the CA certificate.|
|status|text|The status of a CA certificate.|
|validity_not_after|timestamp without time zone|The certificate is not valid after this date.|
|validity_not_before|timestamp without time zone|The certificate is not valid before this date.|
