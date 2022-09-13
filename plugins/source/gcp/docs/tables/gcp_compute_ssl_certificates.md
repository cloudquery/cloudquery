
# Table: gcp_compute_ssl_certificates
Represents an SSL Certificate resource.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|project_id|text|GCP Project Id of the resource|
|certificate|text|A value read into memory from a certificate file The certificate file must be in PEM format The certificate chain must be no greater than 5 certs long The chain must include at least one intermediate cert|
|creation_timestamp|text|Creation timestamp in RFC3339 text format|
|description|text|An optional description of this resource Provide this property when you create the resource|
|expire_time|text|Expire time of the certificate|
|id|text|The unique identifier for the resource This identifier is defined by the server|
|kind|text|Type of the resource Always compute#sslCertificate for SSL certificates|
|managed_domain_status|jsonb|[Output only] Detailed statuses of the domains specified for managed certificate resource|
|managed_domains|text[]|The domains for which a managed SSL certificate will be generated Each Google-managed SSL certificate supports up to the maximum number of domains per Google-managed SSL certificate (/load-balancing/docs/quotas#ssl_certificates)|
|managed_status|text|[Output only] Status of the managed certificate resource|
|name|text|Name of the resource Provided by the client when the resource is created|
|private_key|text|A value read into memory from a write-only private key file The private key file must be in PEM format For security, only insert requests include this field|
|region|text|URL of the region where the regional SSL Certificate resides This field is not applicable to global SSL Certificate|
|self_link|text|[Output only] Server-defined URL for the resource|
|self_managed_certificate|text|A local certificate file The certificate must be in PEM format The certificate chain must be no greater than 5 certs long The chain must include at least one intermediate cert|
|self_managed_private_key|text|A write-only private key in PEM format Only insert requests will include this field|
|subject_alternative_names|text[]|Domains associated with the certificate via Subject Alternative Name|
|type|text|Specifies the type of SSL certificate, either "SELF_MANAGED" or "MANAGED" If not specified, the certificate is self-managed and the fields certificate and private_key are used|
