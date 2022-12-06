# Table: gandi_domains



The primary key for this table is **id**.

## Relations

The following tables depend on gandi_domains:
  - [gandi_domain_livedns](gandi_domain_livedns.md)
  - [gandi_domain_web_redirections](gandi_domain_web_redirections.md)
  - [gandi_domain_glue_records](gandi_domain_glue_records.md)
  - [gandi_domain_dnssec_keys](gandi_domain_dnssec_keys.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|autorenew|JSON|
|can_tld_lock|Bool|
|contacts|JSON|
|dates|JSON|
|fqdn|String|
|fqdn_unicode|String|
|href|String|
|nameservers|StringArray|
|services|StringArray|
|sharing_space|JSON|
|status|StringArray|
|tld|String|
|authinfo|String|
|id (PK)|String|
|sharing_id|String|
|tags|StringArray|
|trustee_roles|StringArray|