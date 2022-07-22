
# Table: aws_lightsail_load_balancer_tls_certificates
Describes a load balancer SSL/TLS certificate
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|load_balancer_cq_id|uuid|Unique CloudQuery ID of aws_lightsail_load_balancers table (FK)|
|arn|text|The Amazon Resource Name (ARN) of the SSL/TLS certificate|
|created_at|timestamp without time zone|The time when you created your SSL/TLS certificate|
|domain_name|text|The domain name for your SSL/TLS certificate|
|domain_validation_records|jsonb|An array of LoadBalancerTlsCertificateDomainValidationRecord objects describing the records|
|failure_reason|text|The validation failure reason, if any, of the certificate|
|is_attached|boolean|When true, the SSL/TLS certificate is attached to the Lightsail load balancer|
|issued_at|timestamp without time zone|The time when the SSL/TLS certificate was issued|
|issuer|text|The issuer of the certificate|
|key_algorithm|text|The algorithm used to generate the key pair (the public and private key)|
|load_balancer_name|text|The load balancer name where your SSL/TLS certificate is attached|
|availability_zone|text|The Availability Zone|
|region_name|text|The AWS Region name|
|name|text|The name of the SSL/TLS certificate (eg, my-certificate)|
|not_after|timestamp without time zone|The timestamp when the SSL/TLS certificate expires|
|not_before|timestamp without time zone|The timestamp when the SSL/TLS certificate is first valid|
|renewal_summary_domain_validation_options|jsonb|Contains information about the validation of each domain name in the certificate, as it pertains to Lightsail's managed renewal|
|renewal_summary_renewal_status|text|The renewal status of the certificate|
|resource_type|text|The resource type (eg, LoadBalancerTlsCertificate)  * Instance - A Lightsail instance (a virtual private server)  * StaticIp - A static IP address  * KeyPair - The key pair used to connect to a Lightsail instance  * InstanceSnapshot - A Lightsail instance snapshot  * Domain - A DNS zone  * PeeredVpc - A peered VPC  * LoadBalancer - A Lightsail load balancer  * LoadBalancerTlsCertificate - An SSL/TLS certificate associated with a Lightsail load balancer  * Disk - A Lightsail block storage disk  * DiskSnapshot - A block storage disk snapshot|
|revocation_reason|text|The reason the certificate was revoked|
|revoked_at|timestamp without time zone|The timestamp when the certificate was revoked|
|serial|text|The serial number of the certificate|
|signature_algorithm|text|The algorithm that was used to sign the certificate|
|status|text|The validation status of the SSL/TLS certificate|
|subject|text|The name of the entity that is associated with the public key contained in the certificate|
|subject_alternative_names|text[]|An array of strings that specify the alternate domains (eg, example2com) and subdomains (eg, blogexamplecom) for the certificate|
|support_code|text|The support code|
|tags|jsonb|The tag keys and optional values for the resource|
