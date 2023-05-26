# Table: heroku_domains

This table shows data for Heroku Domains.

https://devcenter.heroku.com/articles/platform-api-reference#domain

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|id (PK)|utf8|
|acm_status|utf8|
|acm_status_reason|utf8|
|app|json|
|cname|utf8|
|created_at|timestamp[us, tz=UTC]|
|hostname|utf8|
|kind|utf8|
|sni_endpoint|json|
|status|utf8|
|updated_at|timestamp[us, tz=UTC]|