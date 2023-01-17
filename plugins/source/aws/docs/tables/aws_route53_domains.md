# Table: aws_route53_domains

https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetDomainDetail.html

The composite primary key for this table is (**account_id**, **domain_name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|domain_name (PK)|String|
|tags|JSON|
|transfer_lock|Bool|
|abuse_contact_email|String|
|abuse_contact_phone|String|
|admin_contact|JSON|
|admin_privacy|Bool|
|auto_renew|Bool|
|creation_date|Timestamp|
|dns_sec|String|
|dnssec_keys|JSON|
|expiration_date|Timestamp|
|nameservers|JSON|
|registrant_contact|JSON|
|registrant_privacy|Bool|
|registrar_name|String|
|registrar_url|String|
|registry_domain_id|String|
|reseller|String|
|status_list|StringArray|
|tech_contact|JSON|
|tech_privacy|Bool|
|updated_date|Timestamp|
|who_is_server|String|
|result_metadata|JSON|