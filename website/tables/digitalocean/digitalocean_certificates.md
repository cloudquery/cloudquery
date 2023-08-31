# Table: digitalocean_certificates

This table shows data for DigitalOcean Certificates.

https://docs.digitalocean.com/reference/api/api-reference/#tag/Certificates

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|name|`utf8`|
|dns_names|`list<item: utf8, nullable>`|
|not_after|`utf8`|
|sha1_fingerprint|`utf8`|
|created_at|`utf8`|
|state|`utf8`|
|type|`utf8`|