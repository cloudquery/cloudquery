
# Table: aws_iot_certificates
Describes a certificate.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|policies|text[]|Policies of the certificate|
|ca_certificate_id|text|The certificate ID of the CA certificate used to sign this certificate.|
|arn|text|The ARN of the certificate.|
|id|text|The ID of the certificate.|
|mode|text|The mode of the certificate.|
|pem|text|The certificate data, in PEM format.|
|creation_date|timestamp without time zone|The date and time the certificate was created.|
|customer_version|integer|The customer version of the certificate.|
|generation_id|text|The generation ID of the certificate.|
|last_modified_date|timestamp without time zone|The date and time the certificate was last modified.|
|owned_by|text|The ID of the Amazon Web Services account that owns the certificate.|
|previous_owned_by|text|The ID of the Amazon Web Services account of the previous owner of the certificate.|
|status|text|The status of the certificate.|
|transfer_data_accept_date|timestamp without time zone|The date the transfer was accepted.|
|transfer_data_reject_date|timestamp without time zone|The date the transfer was rejected.|
|transfer_data_reject_reason|text|The reason why the transfer was rejected.|
|transfer_data_transfer_date|timestamp without time zone|The date the transfer took place.|
|transfer_data_transfer_message|text|The transfer message.|
|validity_not_after|timestamp without time zone|The certificate is not valid after this date.|
|validity_not_before|timestamp without time zone|The certificate is not valid before this date.|
