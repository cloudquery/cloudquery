
# Table: aws_lightsail_certificates
Describes the full details of an Amazon Lightsail SSL/TLS certificate
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|arn|text|The Amazon Resource Name (ARN) of the certificate|
|created_at|timestamp without time zone|The timestamp when the certificate was created|
|domain_name|text|The domain name of the certificate|
|eligible_to_renew|text|The renewal eligibility of the certificate|
|in_use_resource_count|bigint|The number of Lightsail resources that the certificate is attached to|
|issued_at|timestamp without time zone|The timestamp when the certificate was issued|
|issuer_ca|text|The certificate authority that issued the certificate|
|key_algorithm|text|The algorithm used to generate the key pair (the public and private key) of the certificate|
|name|text|The name of the certificate (eg, my-certificate)|
|not_after|timestamp without time zone|The timestamp when the certificate expires|
|not_before|timestamp without time zone|The timestamp when the certificate is first valid|
|renewal_summary_status|text|The renewal status of the certificate|
|renewal_summary_reason|text|The reason for the renewal status of the certificate|
|renewal_summary_updated_at|timestamp without time zone|The timestamp when the certificate was last updated|
|request_failure_reason|text|The validation failure reason, if any, of the certificate|
|revocation_reason|text|The reason the certificate was revoked|
|revoked_at|timestamp without time zone|The timestamp when the certificate was revoked|
|serial_number|text|The serial number of the certificate|
|status|text|The validation status of the certificate|
|subject_alternative_names|text[]|An array of strings that specify the alternate domains (eg, example2com) and subdomains (eg, blogexamplecom) of the certificate|
|support_code|text|The support code|
|tags|jsonb|The tag keys and optional values for the resource|
