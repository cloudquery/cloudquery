
# Table: aws_acm_certificates
Contains metadata about an ACM certificate
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|arn|text|The Amazon Resource Name (ARN) of the certificate|
|certificate_authority_arn|text|The Amazon Resource Name (ARN) of the ACM PCA private certificate authority (CA) that issued the certificate|
|created_at|timestamp without time zone|The time at which the certificate was requested.|
|domain_name|text|The fully qualified domain name for the certificate, such as www.example.com or example.com.|
|domain_validation_options|jsonb|Contains information about the initial validation of each domain name that occurs as a result of the RequestCertificate request.|
|extended_key_usages|jsonb|Contains a list of Extended Key Usage X.509 v3 extension objects.|
|failure_reason|text|The reason the certificate request failed|
|imported_at|timestamp without time zone|The date and time at which the certificate was imported|
|in_use_by|text[]|A list of ARNs for the Amazon Web Services resources that are using the certificate|
|issued_at|timestamp without time zone|The time at which the certificate was issued|
|issuer|text|The name of the certificate authority that issued and signed the certificate.|
|key_algorithm|text|The algorithm that was used to generate the public-private key pair.|
|key_usages|text[]|A list of Key Usage X.509 v3 extension objects. Each object is a string value that identifies the purpose of the public key contained in the certificate.|
|not_after|timestamp without time zone|The time after which the certificate is not valid.|
|not_before|timestamp without time zone|The time before which the certificate is not valid.|
|certificate_transparency_logging_preference|text|You can opt out of certificate transparency logging by specifying the DISABLED option|
|renewal_eligibility|text|Specifies whether the certificate is eligible for renewal|
|renewal_summary_domain_validation_options|jsonb|Contains information about the validation of each domain name in the certificate, as it pertains to ACM's managed renewal.|
|renewal_summary_status|text|The status of ACM's managed renewal (https://docs.aws.amazon.com/acm/latest/userguide/acm-renewal.html) of the certificate.  This member is required.|
|renewal_summary_updated_at|timestamp without time zone|The time at which the renewal summary was last updated.  This member is required.|
|renewal_summary_failure_reason|text|The reason that a renewal request was unsuccessful.|
|revocation_reason|text|The reason the certificate was revoked|
|revoked_at|timestamp without time zone|The time at which the certificate was revoked|
|serial|text|The serial number of the certificate.|
|signature_algorithm|text|The algorithm that was used to sign the certificate.|
|status|text|The status of the certificate.|
|subject|text|The name of the entity that is associated with the public key contained in the certificate.|
|subject_alternative_names|text[]|One or more domain names (subject alternative names) included in the certificate|
|type|text|The source of the certificate|
|tags|jsonb|The tags that have been applied to the ACM certificate.|
