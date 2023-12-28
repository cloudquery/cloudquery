# Table: cloudflare_access_applications_self_hosted_domains

This table shows data for Cloudflare Access Applications Self Hosted Domains.

The composite primary key for this table is (**application_id**, **domain**).

## Relations

This table depends on [cloudflare_access_applications](cloudflare_access_applications.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|application_id (PK)|`utf8`|
|domain (PK)|`utf8`|