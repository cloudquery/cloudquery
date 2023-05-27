# Table: gandi_certificates

This table shows data for Gandi Certificates.

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|sharing_id|utf8|
|id (PK)|utf8|
|cn|utf8|
|cn_unicode|utf8|
|altnames|list<item: utf8, nullable>|
|altnames_unicode|list<item: utf8, nullable>|
|contact|json|
|dates|json|
|package|json|
|software|int64|
|status|utf8|