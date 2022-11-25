# Table: aws_route53_domains



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
|admin_contact|JSON|
|nameservers|JSON|
|registrant_contact|JSON|
|tech_contact|JSON|
|abuse_contact_email|String|
|abuse_contact_phone|String|
|admin_privacy|Bool|
|auto_renew|Bool|
|creation_date|Timestamp|
|dns_sec|String|
|expiration_date|Timestamp|
|registrant_privacy|Bool|
|registrar_name|String|
|registrar_url|String|
|registry_domain_id|String|
|reseller|String|
|status_list|StringArray|
|tech_privacy|Bool|
|updated_date|Timestamp|
|who_is_server|String|
|result_metadata|JSON|