# Table: gcp_appengine_authorized_domains

This table shows data for GCP App Engine Authorized Domains.

https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.authorizedDomains#AuthorizedDomain

The composite primary key for this table is (**project_id**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|id|`utf8`|