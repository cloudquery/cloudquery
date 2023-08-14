# Table: aws_route53_domains

This table shows data for Amazon Route 53 Domains.

https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetDomainDetail.html

The composite primary key for this table is (**account_id**, **domain_name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|domain_name (PK)|`utf8`|
|tags|`json`|
|transfer_lock|`bool`|
|abuse_contact_email|`utf8`|
|abuse_contact_phone|`utf8`|
|admin_contact|`json`|
|admin_privacy|`bool`|
|auto_renew|`bool`|
|creation_date|`timestamp[us, tz=UTC]`|
|dns_sec|`utf8`|
|dnssec_keys|`json`|
|expiration_date|`timestamp[us, tz=UTC]`|
|nameservers|`json`|
|registrant_contact|`json`|
|registrant_privacy|`bool`|
|registrar_name|`utf8`|
|registrar_url|`utf8`|
|registry_domain_id|`utf8`|
|reseller|`utf8`|
|status_list|`list<item: utf8, nullable>`|
|tech_contact|`json`|
|tech_privacy|`bool`|
|updated_date|`timestamp[us, tz=UTC]`|
|who_is_server|`utf8`|
|result_metadata|`json`|