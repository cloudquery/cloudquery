# Table: gandi_domains

This table shows data for Gandi Domains.

The primary key for this table is **id**.

## Relations

The following tables depend on gandi_domains:
  - [gandi_domain_dnssec_keys](gandi_domain_dnssec_keys)
  - [gandi_domain_glue_records](gandi_domain_glue_records)
  - [gandi_domain_livedns](gandi_domain_livedns)
  - [gandi_domain_web_redirections](gandi_domain_web_redirections)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|sharing_id|`utf8`|
|autorenew|`json`|
|can_tld_lock|`bool`|
|contacts|`json`|
|dates|`json`|
|fqdn|`utf8`|
|fqdn_unicode|`utf8`|
|href|`utf8`|
|nameservers|`list<item: utf8, nullable>`|
|services|`list<item: utf8, nullable>`|
|sharing_space|`json`|
|status|`list<item: utf8, nullable>`|
|tld|`utf8`|
|authinfo|`utf8`|
|tags|`list<item: utf8, nullable>`|
|trustee_roles|`list<item: utf8, nullable>`|