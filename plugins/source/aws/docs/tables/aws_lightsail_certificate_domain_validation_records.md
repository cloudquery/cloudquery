
# Table: aws_lightsail_certificate_domain_validation_records
Describes the domain validation records of an Amazon Lightsail SSL/TLS certificate
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|certificate_cq_id|uuid|Unique CloudQuery ID of aws_lightsail_certificates table (FK)|
|domain_name|text|The domain name of the certificate validation record|
|name|text|The name of the record|
|type|text|The DNS record type|
|value|text|The value for the DNS record|
