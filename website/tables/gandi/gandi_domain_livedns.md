# Table: gandi_domain_livedns

This table shows data for Gandi Domain LiveDNS.

The primary key for this table is **fqdn**.

## Relations

This table depends on [gandi_domains](gandi_domains).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|fqdn (PK)|utf8|
|current|utf8|
|nameservers|list<item: utf8, nullable>|
|dnssec_available|bool|
|livednssec_available|bool|